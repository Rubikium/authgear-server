package userverify

import (
	"fmt"

	"github.com/skygeario/skygear-server/pkg/auth/dependency/provider/password"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	"github.com/skygeario/skygear-server/pkg/core/config"
)

func IsUserVerified(
	authInfo *authinfo.AuthInfo,
	principals []*password.Principal,
	criteria config.UserVerificationCriteria,
	verifyConfigs map[string]config.UserVerificationKeyConfiguration,
) (verified bool) {
	switch criteria {
	case config.UserVerificationCriteriaAll:
		verified = true
		for _, principal := range principals {
			for key := range verifyConfigs {
				if principal.LoginIDKey == key && !authInfo.VerifyInfo[principal.LoginID] {
					verified = false
					return
				}
			}
		}
	case config.UserVerificationCriteriaAny:
		verified = false
		for _, principal := range principals {
			for key := range verifyConfigs {
				if principal.LoginIDKey == key && authInfo.VerifyInfo[principal.LoginID] {
					verified = true
					return
				}
			}
		}
	default:
		panic(fmt.Errorf("unexpected verify criteria `%s`", criteria))
	}
	return
}

type UpdateVerifiedFlagFunc func(*authinfo.AuthInfo, []*password.Principal)

func CreateUpdateVerifiedFlagFunc(tConfig config.TenantConfiguration) UpdateVerifiedFlagFunc {
	criteria := tConfig.UserConfig.UserVerification.Criteria
	verifyConfigs := tConfig.UserConfig.UserVerification.LoginIDKeys
	return func(authInfo *authinfo.AuthInfo, principals []*password.Principal) {
		authInfo.Verified = IsUserVerified(authInfo, principals, criteria, verifyConfigs)
	}
}
