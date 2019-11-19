package handler

import (
	"context"

	"github.com/liuyuexclusive/utils/webutil"
	"github.com/liuyuexclusive/utils/ws"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	message "github.com/liuyuexclusive/future.srv.basic/proto/message"

	user "github.com/liuyuexclusive/future.srv.basic/proto/user"
)

type SendMessage struct {
	Path    string   `json:"path"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	ToList  []string `json:"to_list"`
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			token = c.Query("token")
		}
		res, err := user.NewUserService("go.micro.srv.basic", client.DefaultClient).Validate(context.TODO(), &user.ValidateRequest{Token: token})
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
		}
		c.Set("username", res.Name)
	}
}

// Send godoc
// @Summary
// @Description
// @Tags 发送消息
// @Accept  json
// @Produce  json
// @Param account body handler.SendMessage true "input"
// @Param Authorization header string true "授权码"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /send [put]
func Send(c *gin.Context) {
	var send SendMessage
	if ok := webutil.ReadBody(c, &send); ok {
		_, err := message.NewMessageService("go.micro.srv.basic", client.DefaultClient).Send(context.TODO(), &message.SendRequest{From: c.GetString("username"), ToList: send.ToList, Title: send.Title, Content: send.Content})
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		if err := ws.Send(send.Path, c.GetString("username"), send.ToList, send.Title, send.Content); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "success")
	}
}
