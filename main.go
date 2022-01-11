package main

import (
	"github.com/gin-gonic/gin"
	"go-dummyapp-moysklad/internal/server"
)

const appBaseUrl = "/echo/api/moysklad/vendor/1.0/apps/:appId/:accountId"

func main() {
	//cfg := config.NewConfig("internal/config/config.yml")
	appServer := server.NewServer()
	//log.Println(*appServer.Apps[0].(*dummy.Dummy))
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	router.PUT(appBaseUrl, appServer.ActivateHandler)
	router.DELETE(appBaseUrl, appServer.DeleteHandler)
	router.GET(appBaseUrl, appServer.StatusHandler)

	router.LoadHTMLFiles("internal/dummy-sloudel.sorochinsky/iframe.html")

	router.GET("/echo/iframe/:appUid", appServer.IframeHandler)
	router.POST("/echo/:appUid/update-settings", appServer.UpdateSettingsHandler)

	router.Run(":8002")
}
