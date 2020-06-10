package main

import (
	"net/http"

	"github.com/liuyuexclusive/future.web.ws/handler"

	"github.com/liuyuexclusive/utils/web"
	"github.com/liuyuexclusive/utils/ws"

	"github.com/gin-gonic/gin"
	_ "github.com/liuyuexclusive/future.web.ws/docs"
	"github.com/sirupsen/logrus"
)

type start struct{}

func (s *start) Start(engine *gin.Engine) {
	ws.Serve(engine, "/ws")
	basic := engine.Group("/ws")
	authorized := basic.Group("/")

	authorized.Use(handler.Validate())
	{
		authorized.PUT("/send", handler.Send)

	}

	basic.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test ok")
	})
}

// @title Future对外开放API
// @version 1.0
// @description XXXXXXXXXXX
// @host
// @BasePath
func main() {
	if err := web.Startup("go.micro.api.ws", new(start), func(options *web.Options) {
		options.Port = ":9001"
		options.IsLogToES = false
	}); err != nil {
		logrus.Fatal(err)
	}
}
