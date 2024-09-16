package routers

import (
	"github.com/deveshmishra34/groot/pkg/api/handlers/errors"
	healthHandlers "github.com/deveshmishra34/groot/pkg/api/handlers/healthz"
	usersHandlers "github.com/deveshmishra34/groot/pkg/api/handlers/users"
	"github.com/deveshmishra34/groot/pkg/api/middlewares"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/config"
	"github.com/deveshmishra34/groot/pkg/utils/constants"
)

var protectedApiRouter *Router

func InitProtectedAPIRouter() {
	logger.Debug("Initializing protected api router ...")
	protectedApiRouter = &Router{}
	protectedApiRouter.Name = "protected API"
	protectedApiRouter.Init()

	// order is important here
	// first register development middlewares
	if config.DevModeFlag {
		logger.Debug("Registering protected api development middlewares ...")
		registerProtectedApiDevModeMiddleware()
	}

	// next register middlwares
	logger.Debug("Registering protected api middlewares ...")
	registerProtectedAPIMiddlewares()

	// next register all health check routes
	logger.Debug("Registering protected api health routes ...")
	registerProtectedApiHealthCheckHandlers()

	// next register security related middleware
	logger.Debug("Registering protected api security middlewares ...")
	registerProtectedApiSecurityMiddlewares()

	// next register all routes
	logger.Debug("Registering protected api protected routes ...")
	registerProtectedAPIRoutes()

	// finally register default fallback error handlers
	// 404 is handled here as the last route
	logger.Debug("Registering protected api error handlers ...")
	registerProtectedApiErrorHandlers()

	logger.Debug("Protected api registration complete.")
}

func ProtectedAPIRouter() *Router {
	return protectedApiRouter
}

func registerProtectedAPIMiddlewares() {
	protectedApiRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	protectedApiRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.TimeoutMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	if config.Feature(constants.FEATURE_GZIP).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.GzipMiddleware())
	}
}

func registerProtectedApiDevModeMiddleware() {
	protectedApiRouter.RegisterMiddleware(middlewares.BodyDumpMiddleware())
}

func registerProtectedApiSecurityMiddlewares() {
	protectedApiRouter.RegisterMiddleware(middlewares.XSSCheckMiddleware())

	if config.Feature(constants.FEATURE_CORS).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.CORSMiddleware())
	}

	if config.Feature(constants.FEATURE_ORY_KRATOS).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.AuthenticationWithOryKetoMiddleware())
	}

	// if config.Feature(constants.FEATURE_ORY_KETO).IsEnabled() {
	// 	// keto middleware <- this will check if the user has the right permissions like system admin
	// 	protectedApiRouter.RegisterMiddleware(middlewares.AuthenticationWithOryKetoMiddleware())
	// }

	protectedApiRouter.RegisterMiddleware(middlewares.AuthenticationWithJWTMiddleware())
}

func registerProtectedApiErrorHandlers() {
	protectedApiRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	protectedApiRouter.Echo.RouteNotFound("/*", errors.NotFound)
}

func registerProtectedApiHealthCheckHandlers() {
	health := protectedApiRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerProtectedAPIRoutes() {

	// users routes
	users := protectedApiRouter.Echo.Group("/users")
	users.GET("/info", usersHandlers.Info)
	users.POST("/logout", usersHandlers.LogoutHandler)
	users.POST("/exchange-token", usersHandlers.ExchangeToken)

	// add more routes here ...
}
