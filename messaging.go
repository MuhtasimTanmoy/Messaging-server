package main

import (
	"io"
	"net/http"
	"os"

	"github.com/MuhtasimTanmoy/messaging_server/internal/app/controller"
	"github.com/MuhtasimTanmoy/messaging_server/internal/package/broadcast"
	"github.com/gin-gonic/gin"
)

func main() {

	if os.Getenv("AppMode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("var/logs/gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r := gin.Default()
	r.Static("/static", "./web/static/")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/chat", controller.Chat)

	socket := &broadcast.Websocket{}
	socket.Init()
	r.GET("/ws", func(c *gin.Context) {
		socket.HandleConnections(c.Writer, c.Request, c.DefaultQuery("channel", ""))
	})
	go socket.HandleMessages()

	r.Run()
}
