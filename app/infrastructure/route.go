package infrastructure

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nooboolean/seisankun_api_v2/infrastructure/middleware"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

func RouteV2(app *gin.Engine) {
	app.Use(middleware.CORS())
	app.Use(gin.Logger())
	app.Use(middleware.RequestLogger())

	openAPI(app)
	authAPI(app)
}

func openAPI(app *gin.Engine) {
	// openApiGroup := app.Group("/v2")
}

func authAPI(app *gin.Engine) {
	authApiGroup := app.Group("/v2", gin.BasicAuth(gin.Accounts{
		os.Getenv("SEISANKUN_API_BASIC_USER"): os.Getenv("SEISANKUN_API_BASIC_PASSWORD"),
	}))
	travelController := controllers.NewTravelController(NewSqlHandler(), repositories.NewTransaction(NewTransactionHandler()))
	memberController := controllers.NewMemberController(NewSqlHandler(), repositories.NewTransaction(NewTransactionHandler()))
	paymentController := controllers.NewPaymentController(NewSqlHandler(), repositories.NewTransaction(NewTransactionHandler()))
	borrowingController := controllers.NewBorrowingController(NewSqlHandler())
	calculationController := controllers.NewCalculationController(NewSqlHandler())

	authApiGroup.GET("/travel", travelController.Show)
	authApiGroup.POST("/travel", travelController.Create)
	authApiGroup.PUT("/travel", travelController.Update)
	authApiGroup.DELETE("/travel", travelController.Delete)

	authApiGroup.POST("/member", memberController.Create)
	authApiGroup.DELETE("/member", memberController.Delete)

	authApiGroup.GET("/payment", paymentController.Show)
	authApiGroup.GET("/payment/history", paymentController.Index)
	authApiGroup.POST("/payment", paymentController.Create)
	authApiGroup.PUT("/payment", paymentController.Update)
	authApiGroup.DELETE("/payment", paymentController.Delete)

	authApiGroup.GET("/borrowing/statuses", borrowingController.Show)
	authApiGroup.GET("/borrowing/history", borrowingController.Index)

	authApiGroup.GET("/calculation/results", calculationController.Index)
}
