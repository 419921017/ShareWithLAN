package server

import (
	"ShareWithLAN/server/controller"
	"ShareWithLAN/server/ws"
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

var port = "27149"

func Run() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.GET("/", func(context *gin.Context) {
	//	_, err := context.Writer.Write([]byte("123"))
	//	if err != nil {
	//		return
	//	}
	//})
	hub := initWs()
	router.GET("/ws", func(ctx *gin.Context) {
		ws.HttpController(ctx, hub)
	})
	router.POST("/upload/:path", controller.UploadsController)
	router.POST("/api/v1/files", controller.FilesController)
	router.GET("/api/v1/qrcodes", controller.QrcodeController)
	router.GET("/api/v1/addresses", controller.AddressesController)
	router.POST("/api/v1/texts", controller.TextController)

	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.StaticFS("/static", http.FS(staticFiles))

	router.NoRoute(func(context *gin.Context) {
		path := context.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer func(reader fs.File) {
				err := reader.Close()
				if err != nil {
					log.Fatal(err)

				}
			}(reader)

			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			context.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			context.Status(http.StatusNotFound)
		}
	})
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func initWs() (hub *ws.Hub) {
	hub = ws.NewHub()
	go hub.Run()
	return
}
