package task

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userverify"
	"github.com/skygeario/skygear-server/pkg/auth/model"
	"github.com/skygeario/skygear-server/pkg/auth/task/spec"
	"github.com/skygeario/skygear-server/pkg/core/async"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/errors"
)

func AttachVerifyCodeSendTask(
	executor *async.Executor,
	authDependency auth.DependencyMap,
) {
	// TODO(wire): fix task
	executor.Register(spec.VerifyCodeSendTaskName, &VerifyCodeSendTask{})
}

type VerifyCodeSendTask struct {
	CodeSenderFactory        userverify.CodeSenderFactory `dependency:"UserVerifyCodeSenderFactory"`
	AuthInfoStore            authinfo.Store               `dependency:"AuthInfoStore"`
	UserProfileStore         userprofile.Store            `dependency:"UserProfileStore"`
	UserVerificationProvider userverify.Provider          `dependency:"UserVerificationProvider"`
	PasswordAuthProvider     password.Provider            `dependency:"PasswordAuthProvider"`
	IdentityProvider         principal.IdentityProvider   `dependency:"IdentityProvider"`
	TxContext                db.TxContext                 `dependency:"TxContext"`
	Logger                   *logrus.Entry                `dependency:"HandlerLogger"`
}

func (v *VerifyCodeSendTask) Run(ctx context.Context, param interface{}) (err error) {
	return db.WithTx(v.TxContext, func() error { return v.run(param) })
}

func (v *VerifyCodeSendTask) run(param interface{}) (err error) {
	taskParam := param.(spec.VerifyCodeSendTaskParam)
	loginID := taskParam.LoginID
	userID := taskParam.UserID

	v.Logger.WithFields(logrus.Fields{"user_id": taskParam.UserID}).Debug("Sending verification code")

	authInfo := authinfo.AuthInfo{}
	err = v.AuthInfoStore.GetAuth(userID, &authInfo)
	if err != nil {
		return
	}

	userProfile, err := v.UserProfileStore.GetUserProfile(userID)
	if err != nil {
		return
	}

	// We don't check realms. i.e. Verifying a email means every email login IDs
	// of that email is verified, regardless the realm.
	principals, err := v.PasswordAuthProvider.GetPrincipalsByLoginID("", loginID)
	if err != nil {
		return
	}

	var userPrincipal *password.Principal
	for _, principal := range principals {
		if principal.UserID == authInfo.ID {
			userPrincipal = principal
			break
		}
	}
	if userPrincipal == nil {
		err = errors.WithDetails(errors.New("login ID not found"), errors.Details{"user_id": userID})
		return
	}

	verifyCode, err := v.UserVerificationProvider.CreateVerifyCode(userPrincipal)
	if err != nil {
		return
	}

	codeSender := v.CodeSenderFactory.NewCodeSender(taskParam.URLPrefix, userPrincipal.LoginIDKey)
	user := model.NewUser(authInfo, userProfile)
	if err = codeSender.Send(*verifyCode, user); err != nil {
		err = errors.WithDetails(err, errors.Details{"user_id": userID})
		return
	}

	return nil
}
