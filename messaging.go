package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
	"github.com/MuhtasimTanmoy/messaging_server/internal/app/controller"
	"github.com/MuhtasimTanmoy/messaging_server/internal/package/broadcast"
	"github.com/MuhtasimTanmoy/messaging_server/internal/package/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"flag"
)

func main() {

    fmt.Println("Initiate Messaging Server")

    ///////////////////////////////////////////////////////////////////

    // Load Config file
    var configFile string

    // Flag
    flag.StringVar(&configFile, "config", "config.dist.yml", "config")
    flag.Parse()

    viper.SetConfigFile(configFile)
    err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

    logger.ConfigListing()

    ///////////////////////////////////////////////////////////////////

	if os.Getenv("AppMode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("var/logs/gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}
	r := gin.Default()


	r.Static("/static", "./web/static/")
	r.LoadHTMLGlob("templates/*")


	// Route resolve
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})
	r.GET("/chat", controller.Chat)


	// Socket Initiation
	socket := &broadcast.Websocket{}
	socket.Init()


	// Websokcket
	r.GET("/ws", func(c *gin.Context) {
	     fmt.Println("WS called")
	     fmt.Println(c.DefaultQuery("channel", ""))
		socket.HandleConnections(c.Writer, c.Request, c.DefaultQuery("channel", ""))
	})
	go socket.HandleMessages()
	r.Run()
}
