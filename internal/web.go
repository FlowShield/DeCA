package app

import (
	"github.com/cloudslit/newca/internal/middleware"
	"github.com/gin-gonic/gin"

	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/internal/router"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())
	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// Router register
	r.Register(app)

	return app
}
