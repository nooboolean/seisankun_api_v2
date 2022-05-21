package infrastructure

import (
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Start() error {
	app := setup()
	return endless.ListenAndServe(":"+os.Getenv("HTTP_PORT"), app)
}

func setup() *gin.Engine {
	app := gin.New()
	RouteV2(app)
	return app
}
