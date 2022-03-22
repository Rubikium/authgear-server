// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package background

import (
	"context"
	"github.com/authgear/authgear-server/pkg/lib/audit"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/oob"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/password"
	service2 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/totp"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/anonymous"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/biometric"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/mfa"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/event"
	"github.com/authgear/authgear-server/pkg/lib/facade"
	"github.com/authgear/authgear-server/pkg/lib/feature/accountdeletion"
	"github.com/authgear/authgear-server/pkg/lib/feature/customattrs"
	"github.com/authgear/authgear-server/pkg/lib/feature/stdattrs"
	"github.com/authgear/authgear-server/pkg/lib/feature/verification"
	"github.com/authgear/authgear-server/pkg/lib/feature/welcomemessage"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	oauth2 "github.com/authgear/authgear-server/pkg/lib/oauth"
	"github.com/authgear/authgear-server/pkg/lib/oauth/pq"
	"github.com/authgear/authgear-server/pkg/lib/oauth/redis"
	"github.com/authgear/authgear-server/pkg/lib/ratelimit"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/lib/web"
	"github.com/authgear/authgear-server/pkg/util/backgroundjob"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/template"
)

// Injectors from wire.go:

func newConfigSourceController(p *deps.BackgroundProvider, c context.Context) *configsource.Controller {
	config := p.ConfigSourceConfig
	factory := p.LoggerFactory
	localFSLogger := configsource.NewLocalFSLogger(factory)
	manager := p.BaseResources
	localFS := &configsource.LocalFS{
		Logger:        localFSLogger,
		BaseResources: manager,
		Config:        config,
	}
	databaseLogger := configsource.NewDatabaseLogger(factory)
	environmentConfig := p.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	clock := _wireSystemClockValue
	databaseEnvironmentConfig := &environmentConfig.Database
	sqlBuilder := globaldb.NewSQLBuilder(databaseEnvironmentConfig)
	pool := p.DatabasePool
	handle := globaldb.NewHandle(c, pool, databaseEnvironmentConfig, factory)
	sqlExecutor := globaldb.NewSQLExecutor(c, handle)
	store := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	resolveAppIDType := configsource.NewResolveAppIDTypeDomain()
	database := &configsource.Database{
		Logger:           databaseLogger,
		BaseResources:    manager,
		TrustProxy:       trustProxy,
		Config:           config,
		Clock:            clock,
		Store:            store,
		Database:         handle,
		DatabaseConfig:   databaseEnvironmentConfig,
		ResolveAppIDType: resolveAppIDType,
	}
	controller := configsource.NewController(config, localFS, database)
	return controller
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
)

func newAccountDeletionRunner(p *deps.BackgroundProvider, c context.Context, ctrl *configsource.Controller) *backgroundjob.Runner {
	factory := p.LoggerFactory
	pool := p.DatabasePool
	environmentConfig := p.EnvironmentConfig
	databaseEnvironmentConfig := &environmentConfig.Database
	handle := globaldb.NewHandle(c, pool, databaseEnvironmentConfig, factory)
	sqlBuilder := globaldb.NewSQLBuilder(databaseEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(c, handle)
	clockClock := _wireSystemClockValue
	store := &accountdeletion.Store{
		Handle:      handle,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	userServiceFactory := &UserServiceFactory{
		BackgroundProvider: p,
	}
	runnableLogger := accountdeletion.NewRunnableLogger(factory)
	runnable := &accountdeletion.Runnable{
		Store:              store,
		AppContextResolver: ctrl,
		UserServiceFactory: userServiceFactory,
		Logger:             runnableLogger,
	}
	runner := accountdeletion.NewRunner(factory, runnable)
	return runner
}

func newUserService(ctx context.Context, p *deps.BackgroundProvider, appID string, appContext *config.AppContext) *UserService {
	pool := p.DatabasePool
	databaseConfig := NewDatabaseConfig()
	configConfig := appContext.Config
	secretConfig := configConfig.SecretConfig
	databaseCredentials := deps.ProvideDatabaseCredentials(secretConfig)
	factory := p.LoggerFactory
	handle := appdb.NewHandle(ctx, pool, databaseConfig, databaseCredentials, factory)
	appConfig := configConfig.AppConfig
	configAppID := appConfig.ID
	sqlBuilderApp := appdb.NewSQLBuilderApp(databaseCredentials, configAppID)
	sqlExecutor := appdb.NewSQLExecutor(ctx, handle)
	clockClock := _wireSystemClockValue
	store := &user.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	manager := p.BaseResources
	defaultLanguageTag := deps.ProvideDefaultLanguageTag(configConfig)
	supportedLanguageTags := deps.ProvideSupportedLanguageTags(configConfig)
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	httpConfig := appConfig.HTTP
	localizationConfig := appConfig.Localization
	environmentConfig := p.EnvironmentConfig
	staticAssetURLPrefix := environmentConfig.StaticAssetURLPrefix
	staticAssetResolver := &web.StaticAssetResolver{
		Context:            ctx,
		Config:             httpConfig,
		Localization:       localizationConfig,
		StaticAssetsPrefix: staticAssetURLPrefix,
		Resources:          manager,
	}
	translationService := &translation.Service{
		Context:        ctx,
		TemplateEngine: engine,
		StaticAssets:   staticAssetResolver,
	}
	logger := ratelimit.NewLogger(factory)
	redisPool := p.RedisPool
	hub := p.RedisHub
	redisConfig := NewRedisConfig()
	redisCredentials := deps.ProvideRedisCredentials(secretConfig)
	appredisHandle := appredis.NewHandle(redisPool, hub, redisConfig, redisCredentials, factory)
	storageRedis := &ratelimit.StorageRedis{
		AppID: configAppID,
		Redis: appredisHandle,
	}
	limiter := &ratelimit.Limiter{
		Logger:  logger,
		Storage: storageRedis,
		Clock:   clockClock,
	}
	welcomeMessageConfig := appConfig.WelcomeMessage
	noopTaskQueue := NewNoopTaskQueue()
	remoteIP := ProvideRemoteIP()
	userAgentString := ProvideUserAgentString()
	eventLogger := event.NewLogger(factory)
	sqlBuilder := appdb.NewSQLBuilder(databaseCredentials)
	storeImpl := &event.StoreImpl{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	rawQueries := &user.RawQueries{
		Store: store,
	}
	authenticationConfig := appConfig.Authentication
	identityConfig := appConfig.Identity
	featureConfig := configConfig.FeatureConfig
	identityFeatureConfig := featureConfig.Identity
	serviceStore := &service.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	loginidStore := &loginid.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	loginIDConfig := identityConfig.LoginID
	typeCheckerFactory := &loginid.TypeCheckerFactory{
		Config:    loginIDConfig,
		Resources: manager,
	}
	checker := &loginid.Checker{
		Config:             loginIDConfig,
		TypeCheckerFactory: typeCheckerFactory,
	}
	normalizerFactory := &loginid.NormalizerFactory{
		Config: loginIDConfig,
	}
	provider := &loginid.Provider{
		Store:             loginidStore,
		Config:            loginIDConfig,
		Checker:           checker,
		NormalizerFactory: normalizerFactory,
		Clock:             clockClock,
	}
	oauthStore := &oauth.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	oauthProvider := &oauth.Provider{
		Store: oauthStore,
		Clock: clockClock,
	}
	anonymousStore := &anonymous.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	anonymousProvider := &anonymous.Provider{
		Store: anonymousStore,
		Clock: clockClock,
	}
	biometricStore := &biometric.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	biometricProvider := &biometric.Provider{
		Store: biometricStore,
		Clock: clockClock,
	}
	serviceService := &service.Service{
		Authentication:        authenticationConfig,
		Identity:              identityConfig,
		IdentityFeatureConfig: identityFeatureConfig,
		Store:                 serviceStore,
		LoginID:               provider,
		OAuth:                 oauthProvider,
		Anonymous:             anonymousProvider,
		Biometric:             biometricProvider,
	}
	store2 := &service2.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	passwordStore := &password.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	authenticatorConfig := appConfig.Authenticator
	authenticatorPasswordConfig := authenticatorConfig.Password
	passwordLogger := password.NewLogger(factory)
	historyStore := &password.HistoryStore{
		Clock:       clockClock,
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	passwordChecker := password.ProvideChecker(authenticatorPasswordConfig, historyStore)
	housekeeperLogger := password.NewHousekeeperLogger(factory)
	housekeeper := &password.Housekeeper{
		Store:  historyStore,
		Logger: housekeeperLogger,
		Config: authenticatorPasswordConfig,
	}
	passwordProvider := &password.Provider{
		Store:           passwordStore,
		Config:          authenticatorPasswordConfig,
		Clock:           clockClock,
		Logger:          passwordLogger,
		PasswordHistory: historyStore,
		PasswordChecker: passwordChecker,
		Housekeeper:     housekeeper,
	}
	totpStore := &totp.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	authenticatorTOTPConfig := authenticatorConfig.TOTP
	totpProvider := &totp.Provider{
		Store:  totpStore,
		Config: authenticatorTOTPConfig,
		Clock:  clockClock,
	}
	authenticatorOOBConfig := authenticatorConfig.OOB
	oobStore := &oob.Store{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	storeRedis := &oob.StoreRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	oobLogger := oob.NewLogger(factory)
	oobProvider := &oob.Provider{
		Config:    authenticatorOOBConfig,
		Store:     oobStore,
		CodeStore: storeRedis,
		Clock:     clockClock,
		Logger:    oobLogger,
	}
	service3 := &service2.Service{
		Store:       store2,
		Password:    passwordProvider,
		TOTP:        totpProvider,
		OOBOTP:      oobProvider,
		RateLimiter: limiter,
	}
	verificationLogger := verification.NewLogger(factory)
	verificationConfig := appConfig.Verification
	userProfileConfig := appConfig.UserProfile
	verificationStoreRedis := &verification.StoreRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	storePQ := &verification.StorePQ{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	verificationService := &verification.Service{
		RemoteIP:          remoteIP,
		Logger:            verificationLogger,
		Config:            verificationConfig,
		UserProfileConfig: userProfileConfig,
		Clock:             clockClock,
		CodeStore:         verificationStoreRedis,
		ClaimStore:        storePQ,
		RateLimiter:       limiter,
	}
	httpProto := ProvideHTTPProto()
	httpHost := ProvideHTTPHost()
	imagesCDNHost := environmentConfig.ImagesCDNHost
	pictureTransformer := &stdattrs.PictureTransformer{
		HTTPProto:     httpProto,
		HTTPHost:      httpHost,
		ImagesCDNHost: imagesCDNHost,
	}
	serviceNoEvent := &stdattrs.ServiceNoEvent{
		UserProfileConfig: userProfileConfig,
		Identities:        serviceService,
		UserQueries:       rawQueries,
		UserStore:         store,
		ClaimStore:        storePQ,
		Transformer:       pictureTransformer,
	}
	customattrsServiceNoEvent := &customattrs.ServiceNoEvent{
		Config:      userProfileConfig,
		UserQueries: rawQueries,
		UserStore:   store,
	}
	queries := &user.Queries{
		RawQueries:         rawQueries,
		Store:              store,
		Identities:         serviceService,
		Authenticators:     service3,
		Verification:       verificationService,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
	}
	resolverImpl := &event.ResolverImpl{
		Users: queries,
	}
	hookLogger := hook.NewLogger(factory)
	hookConfig := appConfig.Hook
	webhookKeyMaterials := deps.ProvideWebhookKeyMaterials(secretConfig)
	syncHTTPClient := hook.NewSyncHTTPClient(hookConfig)
	asyncHTTPClient := hook.NewAsyncHTTPClient()
	deliverer := &hook.Deliverer{
		Config:             hookConfig,
		Secret:             webhookKeyMaterials,
		Clock:              clockClock,
		SyncHTTP:           syncHTTPClient,
		AsyncHTTP:          asyncHTTPClient,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
	}
	sink := &hook.Sink{
		Logger:    hookLogger,
		Deliverer: deliverer,
	}
	auditLogger := audit.NewLogger(factory)
	auditDatabaseCredentials := deps.ProvideAuditDatabaseCredentials(secretConfig)
	writeHandle := auditdb.NewWriteHandle(ctx, pool, databaseConfig, auditDatabaseCredentials, factory)
	auditdbSQLBuilderApp := auditdb.NewSQLBuilderApp(auditDatabaseCredentials, configAppID)
	writeSQLExecutor := auditdb.NewWriteSQLExecutor(ctx, writeHandle)
	writeStore := &audit.WriteStore{
		SQLBuilder:  auditdbSQLBuilderApp,
		SQLExecutor: writeSQLExecutor,
	}
	auditSink := &audit.Sink{
		Logger:   auditLogger,
		Database: writeHandle,
		Store:    writeStore,
	}
	eventService := event.NewService(ctx, remoteIP, userAgentString, eventLogger, handle, clockClock, localizationConfig, storeImpl, resolverImpl, sink, auditSink)
	welcomemessageProvider := &welcomemessage.Provider{
		Translation:          translationService,
		RateLimiter:          limiter,
		WelcomeMessageConfig: welcomeMessageConfig,
		TaskQueue:            noopTaskQueue,
		Events:               eventService,
	}
	rawCommands := &user.RawCommands{
		Store:                  store,
		Clock:                  clockClock,
		WelcomeMessageProvider: welcomemessageProvider,
	}
	commands := &user.Commands{
		RawCommands:        rawCommands,
		RawQueries:         rawQueries,
		Events:             eventService,
		Verification:       verificationService,
		UserProfileConfig:  userProfileConfig,
		StandardAttributes: serviceNoEvent,
		CustomAttributes:   customattrsServiceNoEvent,
	}
	userProvider := &user.Provider{
		Commands: commands,
		Queries:  queries,
	}
	storeDeviceTokenRedis := &mfa.StoreDeviceTokenRedis{
		Redis: appredisHandle,
		AppID: configAppID,
		Clock: clockClock,
	}
	storeRecoveryCodePQ := &mfa.StoreRecoveryCodePQ{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	mfaService := &mfa.Service{
		DeviceTokens:  storeDeviceTokenRedis,
		RecoveryCodes: storeRecoveryCodePQ,
		Clock:         clockClock,
		Config:        authenticationConfig,
		RateLimiter:   limiter,
	}
	stdattrsService := &stdattrs.Service{
		UserProfileConfig: userProfileConfig,
		ServiceNoEvent:    serviceNoEvent,
		Identities:        serviceService,
		UserQueries:       rawQueries,
		UserStore:         store,
		Events:            eventService,
	}
	authorizationStore := &pq.AuthorizationStore{
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
	}
	storeRedisLogger := idpsession.NewStoreRedisLogger(factory)
	idpsessionStoreRedis := &idpsession.StoreRedis{
		Redis:  appredisHandle,
		AppID:  configAppID,
		Clock:  clockClock,
		Logger: storeRedisLogger,
	}
	sessionConfig := appConfig.Session
	request := NewDummyHTTPRequest()
	trustProxy := environmentConfig.TrustProxy
	cookieManager := deps.NewCookieManager(request, trustProxy, httpConfig)
	cookieDef := session.NewSessionCookieDef(sessionConfig)
	idpsessionManager := &idpsession.Manager{
		Store:     idpsessionStoreRedis,
		Clock:     clockClock,
		Config:    sessionConfig,
		Cookies:   cookieManager,
		CookieDef: cookieDef,
	}
	redisLogger := redis.NewLogger(factory)
	redisStore := &redis.Store{
		Context:     ctx,
		Redis:       appredisHandle,
		AppID:       configAppID,
		Logger:      redisLogger,
		SQLBuilder:  sqlBuilderApp,
		SQLExecutor: sqlExecutor,
		Clock:       clockClock,
	}
	oAuthConfig := appConfig.OAuth
	sessionManager := &oauth2.SessionManager{
		Store:  redisStore,
		Clock:  clockClock,
		Config: oAuthConfig,
	}
	accountDeletionConfig := appConfig.AccountDeletion
	coordinator := &facade.Coordinator{
		Events:                eventService,
		Identities:            serviceService,
		Authenticators:        service3,
		Verification:          verificationService,
		MFA:                   mfaService,
		UserCommands:          commands,
		UserQueries:           queries,
		StdAttrsService:       stdattrsService,
		PasswordHistory:       historyStore,
		OAuth:                 authorizationStore,
		IDPSessions:           idpsessionManager,
		OAuthSessions:         sessionManager,
		IdentityConfig:        identityConfig,
		AccountDeletionConfig: accountDeletionConfig,
		Clock:                 clockClock,
	}
	userFacade := &facade.UserFacade{
		UserProvider: userProvider,
		Coordinator:  coordinator,
	}
	userService := &UserService{
		AppDBHandle: handle,
		UserFacade:  userFacade,
	}
	return userService
}
