package main
 
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
 
func main() {
  router := gin.Default()

	v1 := router.Group("api/v2")
  v1.GET("/hello", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!!")
  })
 
  router.Run(":3000")
}