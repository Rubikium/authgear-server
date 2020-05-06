// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package sso

import (
	"github.com/gorilla/mux"
	"github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/audit"
	auth2 "github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	redis3 "github.com/skygeario/skygear-server/pkg/auth/dependency/auth/redis"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/bearertoken"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/oob"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/recoverycode"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authenticator/totp"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/hook"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/identity/anonymous"
	loginid2 "github.com/skygeario/skygear-server/pkg/auth/dependency/identity/loginid"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/identity/oauth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/interaction"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/interaction/adaptors"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/interaction/flows"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/interaction/redis"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/loginid"
	oauth2 "github.com/skygeario/skygear-server/pkg/auth/dependency/oauth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oauth/handler"
	pq3 "github.com/skygeario/skygear-server/pkg/auth/dependency/oauth/pq"
	redis2 "github.com/skygeario/skygear-server/pkg/auth/dependency/oauth/redis"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oidc"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/passwordhistory/pq"
	oauth3 "github.com/skygeario/skygear-server/pkg/auth/dependency/principal/oauth"
	password2 "github.com/skygeario/skygear-server/pkg/auth/dependency/principal/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/session"
	redis4 "github.com/skygeario/skygear-server/pkg/auth/dependency/session/redis"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/sso"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/urlprefix"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/core/async"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	pq2 "github.com/skygeario/skygear-server/pkg/core/auth/authinfo/pq"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/db"
	handler2 "github.com/skygeario/skygear-server/pkg/core/handler"
	"github.com/skygeario/skygear-server/pkg/core/logging"
	"github.com/skygeario/skygear-server/pkg/core/time"
	"github.com/skygeario/skygear-server/pkg/core/validation"
	"net/http"
)

// Injectors from wire.go:

func newAuthHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	provider := urlprefix.NewProvider(r)
	authHandlerHTMLProvider := sso.ProvideAuthHandlerHTMLProvider(provider)
	ssoProvider := sso.ProvideSSOProvider(context, tenantConfiguration)
	timeProvider := time.NewProvider()
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, provider, timeProvider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	store := redis.ProvideStore(context, tenantConfiguration, timeProvider)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	loginidProvider := loginid2.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration, reservedNameChecker)
	oauthProvider := oauth.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider)
	anonymousProvider := anonymous.ProvideProvider(sqlBuilder, sqlExecutor)
	identityAdaptor := &adaptors.IdentityAdaptor{
		LoginID:   loginidProvider,
		OAuth:     oauthProvider,
		Anonymous: anonymousProvider,
	}
	passwordhistoryStore := pq.ProvidePasswordHistoryStore(timeProvider, sqlBuilder, sqlExecutor)
	passwordChecker := audit.ProvidePasswordChecker(tenantConfiguration, passwordhistoryStore)
	passwordProvider := password.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, factory, passwordhistoryStore, passwordChecker, tenantConfiguration)
	totpProvider := totp.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	executor := auth.ProvideTaskExecutor(m)
	queue := async.ProvideTaskQueue(context, txContext, requestID, tenantConfiguration, executor)
	oobProvider := oob.ProvideProvider(tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider, engine, provider, queue)
	bearertokenProvider := bearertoken.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	recoverycodeProvider := recoverycode.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	authenticatorAdaptor := &adaptors.AuthenticatorAdaptor{
		Password:     passwordProvider,
		TOTP:         totpProvider,
		OOBOTP:       oobProvider,
		BearerToken:  bearertokenProvider,
		RecoveryCode: recoverycodeProvider,
	}
	authinfoStore := pq2.ProvideStore(sqlBuilderFactory, sqlExecutor)
	userprofileStore := userprofile.ProvideStore(timeProvider, sqlBuilder, sqlExecutor)
	hookProvider := hook.ProvideHookProvider(context, sqlBuilder, sqlExecutor, requestID, tenantConfiguration, txContext, timeProvider, authinfoStore, userprofileStore, loginidProvider, factory)
	userProvider := interaction.ProvideUserProvider(authinfoStore, userprofileStore, timeProvider, hookProvider, provider, queue, tenantConfiguration)
	interactionProvider := interaction.ProvideProvider(store, timeProvider, factory, identityAdaptor, authenticatorAdaptor, userProvider, oobProvider, tenantConfiguration, hookProvider)
	authorizationStore := &pq3.AuthorizationStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	grantStore := redis2.ProvideGrantStore(context, factory, tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider)
	eventStore := redis3.ProvideEventStore(context, tenantConfiguration)
	accessEventProvider := auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionStore := redis4.ProvideStore(context, tenantConfiguration, timeProvider, factory)
	authAccessEventProvider := &auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionProvider := session.ProvideSessionProvider(r, sessionStore, authAccessEventProvider, tenantConfiguration)
	idTokenIssuer := oidc.ProvideIDTokenIssuer(tenantConfiguration, provider, authinfoStore, userprofileStore, timeProvider)
	tokenGenerator := _wireTokenGeneratorValue
	tokenHandler := handler.ProvideTokenHandler(r, tenantConfiguration, factory, authorizationStore, grantStore, grantStore, grantStore, accessEventProvider, sessionProvider, idTokenIssuer, tokenGenerator, timeProvider)
	insecureCookieConfig := auth.ProvideSessionInsecureCookieConfig(m)
	cookieConfiguration := session.ProvideSessionCookieConfiguration(r, insecureCookieConfig, tenantConfiguration)
	userController := flows.ProvideUserController(authinfoStore, userprofileStore, tokenHandler, cookieConfiguration, sessionProvider, hookProvider, timeProvider, tenantConfiguration)
	authAPIFlow := &flows.AuthAPIFlow{
		Interactions:   interactionProvider,
		UserController: userController,
	}
	httpHandler := provideAuthHandler(txContext, tenantConfiguration, authHandlerHTMLProvider, ssoProvider, oAuthProvider, authAPIFlow)
	return httpHandler
}

var (
	_wireTokenGeneratorValue = handler.TokenGenerator(oauth2.GenerateToken)
)

func newAuthResultHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	validator := auth.ProvideValidator(m)
	provider := sso.ProvideSSOProvider(context, tenantConfiguration)
	timeProvider := time.NewProvider()
	store := redis.ProvideStore(context, tenantConfiguration, timeProvider)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	loginidProvider := loginid2.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration, reservedNameChecker)
	oauthProvider := oauth.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider)
	anonymousProvider := anonymous.ProvideProvider(sqlBuilder, sqlExecutor)
	identityAdaptor := &adaptors.IdentityAdaptor{
		LoginID:   loginidProvider,
		OAuth:     oauthProvider,
		Anonymous: anonymousProvider,
	}
	passwordhistoryStore := pq.ProvidePasswordHistoryStore(timeProvider, sqlBuilder, sqlExecutor)
	passwordChecker := audit.ProvidePasswordChecker(tenantConfiguration, passwordhistoryStore)
	passwordProvider := password.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, factory, passwordhistoryStore, passwordChecker, tenantConfiguration)
	totpProvider := totp.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	urlprefixProvider := urlprefix.NewProvider(r)
	executor := auth.ProvideTaskExecutor(m)
	queue := async.ProvideTaskQueue(context, txContext, requestID, tenantConfiguration, executor)
	oobProvider := oob.ProvideProvider(tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider, engine, urlprefixProvider, queue)
	bearertokenProvider := bearertoken.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	recoverycodeProvider := recoverycode.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	authenticatorAdaptor := &adaptors.AuthenticatorAdaptor{
		Password:     passwordProvider,
		TOTP:         totpProvider,
		OOBOTP:       oobProvider,
		BearerToken:  bearertokenProvider,
		RecoveryCode: recoverycodeProvider,
	}
	authinfoStore := pq2.ProvideStore(sqlBuilderFactory, sqlExecutor)
	userprofileStore := userprofile.ProvideStore(timeProvider, sqlBuilder, sqlExecutor)
	hookProvider := hook.ProvideHookProvider(context, sqlBuilder, sqlExecutor, requestID, tenantConfiguration, txContext, timeProvider, authinfoStore, userprofileStore, loginidProvider, factory)
	userProvider := interaction.ProvideUserProvider(authinfoStore, userprofileStore, timeProvider, hookProvider, urlprefixProvider, queue, tenantConfiguration)
	interactionProvider := interaction.ProvideProvider(store, timeProvider, factory, identityAdaptor, authenticatorAdaptor, userProvider, oobProvider, tenantConfiguration, hookProvider)
	authorizationStore := &pq3.AuthorizationStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	grantStore := redis2.ProvideGrantStore(context, factory, tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider)
	eventStore := redis3.ProvideEventStore(context, tenantConfiguration)
	accessEventProvider := auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionStore := redis4.ProvideStore(context, tenantConfiguration, timeProvider, factory)
	authAccessEventProvider := &auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionProvider := session.ProvideSessionProvider(r, sessionStore, authAccessEventProvider, tenantConfiguration)
	idTokenIssuer := oidc.ProvideIDTokenIssuer(tenantConfiguration, urlprefixProvider, authinfoStore, userprofileStore, timeProvider)
	tokenGenerator := _wireTokenGeneratorValue
	tokenHandler := handler.ProvideTokenHandler(r, tenantConfiguration, factory, authorizationStore, grantStore, grantStore, grantStore, accessEventProvider, sessionProvider, idTokenIssuer, tokenGenerator, timeProvider)
	insecureCookieConfig := auth.ProvideSessionInsecureCookieConfig(m)
	cookieConfiguration := session.ProvideSessionCookieConfiguration(r, insecureCookieConfig, tenantConfiguration)
	userController := flows.ProvideUserController(authinfoStore, userprofileStore, tokenHandler, cookieConfiguration, sessionProvider, hookProvider, timeProvider, tenantConfiguration)
	authAPIFlow := &flows.AuthAPIFlow{
		Interactions:   interactionProvider,
		UserController: userController,
	}
	httpHandler := provideAuthResultHandler(txContext, requireAuthz, validator, provider, authAPIFlow)
	return httpHandler
}

func newLinkHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	validator := auth.ProvideValidator(m)
	provider := sso.ProvideSSOProvider(context, tenantConfiguration)
	urlprefixProvider := urlprefix.NewProvider(r)
	timeProvider := time.NewProvider()
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, timeProvider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	store := redis.ProvideStore(context, tenantConfiguration, timeProvider)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	loginidProvider := loginid2.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration, reservedNameChecker)
	oauthProvider := oauth.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider)
	anonymousProvider := anonymous.ProvideProvider(sqlBuilder, sqlExecutor)
	identityAdaptor := &adaptors.IdentityAdaptor{
		LoginID:   loginidProvider,
		OAuth:     oauthProvider,
		Anonymous: anonymousProvider,
	}
	passwordhistoryStore := pq.ProvidePasswordHistoryStore(timeProvider, sqlBuilder, sqlExecutor)
	passwordChecker := audit.ProvidePasswordChecker(tenantConfiguration, passwordhistoryStore)
	passwordProvider := password.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, factory, passwordhistoryStore, passwordChecker, tenantConfiguration)
	totpProvider := totp.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	executor := auth.ProvideTaskExecutor(m)
	queue := async.ProvideTaskQueue(context, txContext, requestID, tenantConfiguration, executor)
	oobProvider := oob.ProvideProvider(tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider, engine, urlprefixProvider, queue)
	bearertokenProvider := bearertoken.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	recoverycodeProvider := recoverycode.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	authenticatorAdaptor := &adaptors.AuthenticatorAdaptor{
		Password:     passwordProvider,
		TOTP:         totpProvider,
		OOBOTP:       oobProvider,
		BearerToken:  bearertokenProvider,
		RecoveryCode: recoverycodeProvider,
	}
	authinfoStore := pq2.ProvideStore(sqlBuilderFactory, sqlExecutor)
	userprofileStore := userprofile.ProvideStore(timeProvider, sqlBuilder, sqlExecutor)
	hookProvider := hook.ProvideHookProvider(context, sqlBuilder, sqlExecutor, requestID, tenantConfiguration, txContext, timeProvider, authinfoStore, userprofileStore, loginidProvider, factory)
	userProvider := interaction.ProvideUserProvider(authinfoStore, userprofileStore, timeProvider, hookProvider, urlprefixProvider, queue, tenantConfiguration)
	interactionProvider := interaction.ProvideProvider(store, timeProvider, factory, identityAdaptor, authenticatorAdaptor, userProvider, oobProvider, tenantConfiguration, hookProvider)
	authorizationStore := &pq3.AuthorizationStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	grantStore := redis2.ProvideGrantStore(context, factory, tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider)
	eventStore := redis3.ProvideEventStore(context, tenantConfiguration)
	accessEventProvider := auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionStore := redis4.ProvideStore(context, tenantConfiguration, timeProvider, factory)
	authAccessEventProvider := &auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionProvider := session.ProvideSessionProvider(r, sessionStore, authAccessEventProvider, tenantConfiguration)
	idTokenIssuer := oidc.ProvideIDTokenIssuer(tenantConfiguration, urlprefixProvider, authinfoStore, userprofileStore, timeProvider)
	tokenGenerator := _wireTokenGeneratorValue
	tokenHandler := handler.ProvideTokenHandler(r, tenantConfiguration, factory, authorizationStore, grantStore, grantStore, grantStore, accessEventProvider, sessionProvider, idTokenIssuer, tokenGenerator, timeProvider)
	insecureCookieConfig := auth.ProvideSessionInsecureCookieConfig(m)
	cookieConfiguration := session.ProvideSessionCookieConfiguration(r, insecureCookieConfig, tenantConfiguration)
	userController := flows.ProvideUserController(authinfoStore, userprofileStore, tokenHandler, cookieConfiguration, sessionProvider, hookProvider, timeProvider, tenantConfiguration)
	authAPIFlow := &flows.AuthAPIFlow{
		Interactions:   interactionProvider,
		UserController: userController,
	}
	httpHandler := provideLinkHandler(txContext, requireAuthz, validator, provider, oAuthProvider, authAPIFlow)
	return httpHandler
}

func newLoginHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	validator := auth.ProvideValidator(m)
	provider := sso.ProvideSSOProvider(context, tenantConfiguration)
	timeProvider := time.NewProvider()
	store := redis.ProvideStore(context, tenantConfiguration, timeProvider)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	loginidProvider := loginid2.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration, reservedNameChecker)
	oauthProvider := oauth.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider)
	anonymousProvider := anonymous.ProvideProvider(sqlBuilder, sqlExecutor)
	identityAdaptor := &adaptors.IdentityAdaptor{
		LoginID:   loginidProvider,
		OAuth:     oauthProvider,
		Anonymous: anonymousProvider,
	}
	passwordhistoryStore := pq.ProvidePasswordHistoryStore(timeProvider, sqlBuilder, sqlExecutor)
	passwordChecker := audit.ProvidePasswordChecker(tenantConfiguration, passwordhistoryStore)
	passwordProvider := password.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, factory, passwordhistoryStore, passwordChecker, tenantConfiguration)
	totpProvider := totp.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	urlprefixProvider := urlprefix.NewProvider(r)
	executor := auth.ProvideTaskExecutor(m)
	queue := async.ProvideTaskQueue(context, txContext, requestID, tenantConfiguration, executor)
	oobProvider := oob.ProvideProvider(tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider, engine, urlprefixProvider, queue)
	bearertokenProvider := bearertoken.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	recoverycodeProvider := recoverycode.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration)
	authenticatorAdaptor := &adaptors.AuthenticatorAdaptor{
		Password:     passwordProvider,
		TOTP:         totpProvider,
		OOBOTP:       oobProvider,
		BearerToken:  bearertokenProvider,
		RecoveryCode: recoverycodeProvider,
	}
	authinfoStore := pq2.ProvideStore(sqlBuilderFactory, sqlExecutor)
	userprofileStore := userprofile.ProvideStore(timeProvider, sqlBuilder, sqlExecutor)
	hookProvider := hook.ProvideHookProvider(context, sqlBuilder, sqlExecutor, requestID, tenantConfiguration, txContext, timeProvider, authinfoStore, userprofileStore, loginidProvider, factory)
	userProvider := interaction.ProvideUserProvider(authinfoStore, userprofileStore, timeProvider, hookProvider, urlprefixProvider, queue, tenantConfiguration)
	interactionProvider := interaction.ProvideProvider(store, timeProvider, factory, identityAdaptor, authenticatorAdaptor, userProvider, oobProvider, tenantConfiguration, hookProvider)
	authorizationStore := &pq3.AuthorizationStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	grantStore := redis2.ProvideGrantStore(context, factory, tenantConfiguration, sqlBuilder, sqlExecutor, timeProvider)
	eventStore := redis3.ProvideEventStore(context, tenantConfiguration)
	accessEventProvider := auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionStore := redis4.ProvideStore(context, tenantConfiguration, timeProvider, factory)
	authAccessEventProvider := &auth2.AccessEventProvider{
		Store: eventStore,
	}
	sessionProvider := session.ProvideSessionProvider(r, sessionStore, authAccessEventProvider, tenantConfiguration)
	idTokenIssuer := oidc.ProvideIDTokenIssuer(tenantConfiguration, urlprefixProvider, authinfoStore, userprofileStore, timeProvider)
	tokenGenerator := _wireTokenGeneratorValue
	tokenHandler := handler.ProvideTokenHandler(r, tenantConfiguration, factory, authorizationStore, grantStore, grantStore, grantStore, accessEventProvider, sessionProvider, idTokenIssuer, tokenGenerator, timeProvider)
	insecureCookieConfig := auth.ProvideSessionInsecureCookieConfig(m)
	cookieConfiguration := session.ProvideSessionCookieConfiguration(r, insecureCookieConfig, tenantConfiguration)
	userController := flows.ProvideUserController(authinfoStore, userprofileStore, tokenHandler, cookieConfiguration, sessionProvider, hookProvider, timeProvider, tenantConfiguration)
	authAPIFlow := &flows.AuthAPIFlow{
		Interactions:   interactionProvider,
		UserController: userController,
	}
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, timeProvider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	httpHandler := provideLoginHandler(txContext, requireAuthz, validator, provider, authAPIFlow, oAuthProvider)
	return httpHandler
}

func newAuthRedirectHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	provider := sso.ProvideSSOProvider(context, tenantConfiguration)
	urlprefixProvider := urlprefix.NewProvider(r)
	timeProvider := time.NewProvider()
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, timeProvider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	httpHandler := provideAuthRedirectHandler(provider, oAuthProvider)
	return httpHandler
}

func newLoginAuthURLHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	validator := auth.ProvideValidator(m)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	provider := time.NewProvider()
	store := pq.ProvidePasswordHistoryStore(provider, sqlBuilder, sqlExecutor)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	passwordProvider := password2.ProvidePasswordProvider(sqlBuilder, sqlExecutor, provider, store, factory, tenantConfiguration, reservedNameChecker)
	ssoProvider := sso.ProvideSSOProvider(context, tenantConfiguration)
	urlprefixProvider := urlprefix.NewProvider(r)
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, provider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	ssoSsoAction := providerLoginSSOAction()
	httpHandler := provideAuthURLHandler(txContext, requireAuthz, validator, passwordProvider, ssoProvider, tenantConfiguration, oAuthProvider, ssoSsoAction)
	return httpHandler
}

func newLinkAuthURLHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	validator := auth.ProvideValidator(m)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	provider := time.NewProvider()
	store := pq.ProvidePasswordHistoryStore(provider, sqlBuilder, sqlExecutor)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	passwordProvider := password2.ProvidePasswordProvider(sqlBuilder, sqlExecutor, provider, store, factory, tenantConfiguration, reservedNameChecker)
	ssoProvider := sso.ProvideSSOProvider(context, tenantConfiguration)
	urlprefixProvider := urlprefix.NewProvider(r)
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, provider, loginIDNormalizerFactory, redirectURLFunc)
	oAuthProvider := provideOAuthProviderFromRequestVars(r, oAuthProviderFactory)
	ssoSsoAction := providerLinkSSOAction()
	httpHandler := provideAuthURLHandler(txContext, requireAuthz, validator, passwordProvider, ssoProvider, tenantConfiguration, oAuthProvider, ssoSsoAction)
	return httpHandler
}

func newUnlinkHandler(r *http.Request, m auth.DependencyMap) http.Handler {
	context := auth.ProvideContext(r)
	tenantConfiguration := auth.ProvideTenantConfig(context, m)
	txContext := db.ProvideTxContext(context, tenantConfiguration)
	requestID := auth.ProvideLoggingRequestID(r)
	factory := logging.ProvideLoggerFactory(context, requestID, tenantConfiguration)
	requireAuthz := handler2.NewRequireAuthzFactory(factory)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(context, tenantConfiguration)
	provider := oauth3.ProvideOAuthProvider(sqlBuilder, sqlExecutor)
	store := pq2.ProvideStore(sqlBuilderFactory, sqlExecutor)
	timeProvider := time.NewProvider()
	userprofileStore := userprofile.ProvideStore(timeProvider, sqlBuilder, sqlExecutor)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	loginidProvider := loginid2.ProvideProvider(sqlBuilder, sqlExecutor, timeProvider, tenantConfiguration, reservedNameChecker)
	hookProvider := hook.ProvideHookProvider(context, sqlBuilder, sqlExecutor, requestID, tenantConfiguration, txContext, timeProvider, store, userprofileStore, loginidProvider, factory)
	urlprefixProvider := urlprefix.NewProvider(r)
	loginIDNormalizerFactory := loginid.ProvideLoginIDNormalizerFactory(tenantConfiguration)
	redirectURLFunc := ProvideRedirectURIForAPIFunc()
	oAuthProviderFactory := sso.ProvideOAuthProviderFactory(tenantConfiguration, urlprefixProvider, timeProvider, loginIDNormalizerFactory, redirectURLFunc)
	httpHandler := providerUnlinkHandler(txContext, requireAuthz, provider, store, userprofileStore, hookProvider, oAuthProviderFactory)
	return httpHandler
}

// wire.go:

func provideOAuthProviderFromRequestVars(r *http.Request, spf *sso.OAuthProviderFactory) sso.OAuthProvider {
	vars := mux.Vars(r)
	return spf.NewOAuthProvider(vars["provider"])
}

func ProvideRedirectURIForAPIFunc() sso.RedirectURLFunc {
	return RedirectURIForAPI
}

func provideAuthHandler(
	tx db.TxContext,
	cfg *config.TenantConfiguration,
	hp sso.AuthHandlerHTMLProvider,
	sp sso.Provider,
	op sso.OAuthProvider,
	f OAuthHandlerInteractionFlow,
) http.Handler {
	h := &AuthHandler{
		TxContext:               tx,
		TenantConfiguration:     cfg,
		AuthHandlerHTMLProvider: hp,
		SSOProvider:             sp,
		OAuthProvider:           op,
		Interactions:            f,
	}
	return h
}

func provideAuthResultHandler(
	tx db.TxContext,
	requireAuthz handler2.RequireAuthz,
	v *validation.Validator,
	sp sso.Provider,
	f OAuthResultInteractionFlow,
) http.Handler {
	h := &AuthResultHandler{
		TxContext:    tx,
		Validator:    v,
		SSOProvider:  sp,
		Interactions: f,
	}
	return requireAuthz(h, h)
}

func provideLinkHandler(
	tx db.TxContext,
	requireAuthz handler2.RequireAuthz,
	v *validation.Validator,
	sp sso.Provider,
	op sso.OAuthProvider,
	f OAuthLinkInteractionFlow,
) http.Handler {
	h := &LinkHandler{
		TxContext:     tx,
		Validator:     v,
		SSOProvider:   sp,
		OAuthProvider: op,
		Interactions:  f,
	}
	return requireAuthz(h, h)
}

func provideLoginHandler(
	tx db.TxContext,
	requireAuthz handler2.RequireAuthz,
	v *validation.Validator,
	sp sso.Provider,
	f OAuthLoginInteractionFlow,
	op sso.OAuthProvider,
) http.Handler {
	h := &LoginHandler{
		TxContext:     tx,
		Validator:     v,
		SSOProvider:   sp,
		OAuthProvider: op,
		Interactions:  f,
	}
	return requireAuthz(h, h)
}

func provideAuthRedirectHandler(
	sp sso.Provider,
	op sso.OAuthProvider,
) http.Handler {
	h := &AuthRedirectHandler{
		SSOProvider:   sp,
		OAuthProvider: op,
	}
	return h
}

func provideAuthURLHandler(
	tx db.TxContext,
	requireAuthz handler2.RequireAuthz,
	v *validation.Validator,
	pp password2.Provider,
	sp sso.Provider,
	cfg *config.TenantConfiguration,
	op sso.OAuthProvider,
	action ssoAction,
) http.Handler {
	h := &AuthURLHandler{
		TxContext:                  tx,
		Validator:                  v,
		PasswordAuthProvider:       pp,
		SSOProvider:                sp,
		OAuthConflictConfiguration: cfg.AppConfig.AuthAPI.OnIdentityConflict.OAuth,
		OAuthProvider:              op,
		Action:                     action,
	}
	return requireAuthz(h, h)
}

func providerLoginSSOAction() ssoAction {
	return ssoActionLogin
}

func providerLinkSSOAction() ssoAction {
	return ssoActionLink
}

func providerUnlinkHandler(
	tx db.TxContext,
	requireAuthz handler2.RequireAuthz,
	oap oauth3.Provider,
	ais authinfo.Store,
	ups userprofile.Store,
	hp hook.Provider,
	spf *sso.OAuthProviderFactory,
) http.Handler {
	h := &UnlinkHandler{
		TxContext:         tx,
		OAuthAuthProvider: oap,
		AuthInfoStore:     ais,
		UserProfileStore:  ups,
		HookProvider:      hp,
		ProviderFactory:   spf,
	}
	return requireAuthz(h, h)
}
