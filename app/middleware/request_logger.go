package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		fmt.Print("Request: ")
		fmt.Println(c.Request)
		fmt.Print("Parameters; ")

		// TODO: 本番環境でコンソールロギングは必要無いので変えられるようにしておく
		requestLogging(c)

		c.Next()
		fmt.Print("Response: ")
		fmt.Println(blw.body.String())
	}
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func requestLogging(c *gin.Context) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])
	fmt.Println(b)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))
}
