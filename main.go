package main

import (
	"github.com/liuyuexclusive/future.web.ws/handler"

	"github.com/liuyuexclusive/utils/webutil"
	"github.com/liuyuexclusive/utils/ws"

	"github.com/gin-gonic/gin"
	_ "github.com/liuyuexclusive/future.web.ws/docs"
	"github.com/sirupsen/logrus"
)

type start struct{}

func (s *start) Start(engine *gin.Engine) {
	engine.Use(handler.Validate())
	ws.Serve(engine, "/ws")
	engine.PUT("/send", handler.Send)
}

// @title Future对外开放API
// @version 1.0
// @description XXXXXXXXXXX
// @host
// @BasePath
func main() {
	if err := webutil.Startup("go.micro.ws.ws", new(start), func(options *webutil.Options) {
		options.Port = ":9001"
		options.IsLogToES = false
	}); err != nil {
		logrus.Fatal(err)
	}
}
