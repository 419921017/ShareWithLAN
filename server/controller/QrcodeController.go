package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func QrcodeController(ctx *gin.Context) {
	if content := ctx.Query("content"); content != "" {
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Data(http.StatusOK, "image/png", png)
	} else {
		ctx.Status(http.StatusBadRequest)
	}
}
