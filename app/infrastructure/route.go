package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/nooboolean/seisankun_api_v2/infrastructure/middleware"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers"
)

func RouteV2(app *gin.Engine) {
	app.Use(middleware.CORS())
	app.Use(gin.Logger())
	app.Use(middleware.RequestLogger())

	openAPI(app)
	authAPI(app)
}

func openAPI(app *gin.Engine) {
	openApiGroup := app.Group("/v2")
	travelController := controllers.NewTravelController(NewSqlHandler())
	memberController := controllers.NewMemberController(NewSqlHandler())

	openApiGroup.GET("/travel", travelController.Show)
	openApiGroup.POST("/travel", travelController.Create)
	openApiGroup.PUT("/travel", travelController.Update)
	openApiGroup.DELETE("/travel", travelController.Delete)

	openApiGroup.POST("/member", memberController.Create)
	openApiGroup.DELETE("/member", memberController.Delete)
}

func authAPI(app *gin.Engine) {
}
