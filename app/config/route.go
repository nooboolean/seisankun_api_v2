package config

import (
	"github.com/gin-gonic/gin"
	"github.com/nooboolean/seisankun_api_v2/middleware"
	"github.com/nooboolean/seisankun_api_v2/presentations/handlers"
)

func RouteV2(app *gin.Engine) {
	app.Use(middleware.CORS())
	app.Use(gin.Logger())
	app.Use(middleware.RequestLogger())

	openAPI(app)
	authAPI(app)
}

func openAPI(app *gin.Engine) {
	travelHandler := handlers.NewTravelHandler()
	openApiGroup := app.Group("/v2")

	openApiGroup.GET("/travel", travelHandler.Get)
	openApiGroup.POST("/travel", travelHandler.Post)
}

func authAPI(app *gin.Engine) {
}
