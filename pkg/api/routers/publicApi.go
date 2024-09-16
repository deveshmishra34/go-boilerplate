package routers

import (
	catsHandlers "github.com/deveshmishra34/groot/pkg/api/handlers/cats"
	"github.com/deveshmishra34/groot/pkg/api/handlers/errors"
	healthHandlers "github.com/deveshmishra34/groot/pkg/api/handlers/healthz"
	usersHandlers "github.com/deveshmishra34/groot/pkg/api/handlers/users"
	"github.com/deveshmishra34/groot/pkg/api/middlewares"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/config"
	"github.com/deveshmishra34/groot/pkg/utils/constants"
)

var publicApiRouter *Router

func InitPublicAPIRouter() {
	logger.Debug("Initializing public api router ...")
	publicApiRouter = &Router{}
	publicApiRouter.Name = "public API"
	publicApiRouter.Init()

	// order is important here
	// first register development middlewares
	if config.DevModeFlag {
		logger.Debug("Registering public api development middlewares ...")
		registerPublicApiDevModeMiddleware()
	}

	// next register middlwares
	logger.Debug("Registering public api middlewares ...")
	registerPublicAPIMiddlewares()

	// next register all health check routes
	logger.Debug("Registering public api health routes ...")
	registerPublicApiHealthCheckHandlers()

	// next register security related middleware
	logger.Debug("Registering public api security middlewares ...")
	registerPublicApiSecurityMiddlewares()

	// next register all routes
	logger.Debug("Registering public api public routes ...")
	registerPublicAPIRoutes()

	// finally register default fallback error handlers
	// 404 is handled here as the last route
	logger.Debug("Registering public api error handlers ...")
	registerPublicApiErrorHandlers()

	logger.Debug("Public api registration complete.")
}

func PublicAPIRouter() *Router {
	return publicApiRouter
}

func registerPublicAPIMiddlewares() {
	publicApiRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	publicApiRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.TimeoutMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	if config.Feature(constants.FEATURE_GZIP).IsEnabled() {
		publicApiRouter.RegisterMiddleware(middlewares.GzipMiddleware())
	}
}

func registerPublicApiDevModeMiddleware() {
	publicApiRouter.RegisterMiddleware(middlewares.BodyDumpMiddleware())
}

func registerPublicApiSecurityMiddlewares() {
	publicApiRouter.RegisterMiddleware(middlewares.XSSCheckMiddleware())

	if config.Feature(constants.FEATURE_CORS).IsEnabled() {
		publicApiRouter.RegisterMiddleware(middlewares.CORSMiddleware())
	}

}

func registerPublicApiErrorHandlers() {
	logger.Debug("registerPublicApiErrorHandlers")
	publicApiRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	publicApiRouter.Echo.RouteNotFound("/*", errors.NotFound)
}

func registerPublicApiHealthCheckHandlers() {
	health := publicApiRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerPublicAPIRoutes() {
	// add all public routes here

	// cats routes
	cats := publicApiRouter.Echo.Group("/cats")
	cats.GET("", catsHandlers.Index)
	cats.GET("/:id", catsHandlers.Get)
	cats.POST("", catsHandlers.Post)
	cats.PUT("/:id", catsHandlers.Put)
	cats.DELETE("/:id", catsHandlers.Delete)

	// users routes
	users := publicApiRouter.Echo.Group("/users")
	users.GET("", usersHandlers.GetAll)
	users.GET("/:id", usersHandlers.Get)
	users.PUT("/:id", usersHandlers.Put)
	users.DELETE("/:id", usersHandlers.Delete)

	users.POST("/register", usersHandlers.Post)
	users.POST("/login", usersHandlers.LoginUserWithPassword)
	users.POST("/get-access-token", usersHandlers.LoginUserWithPassword)
	users.POST("/request-otp", usersHandlers.SendOtp)

	// add more routes here ...
}
