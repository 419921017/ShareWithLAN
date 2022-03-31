package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadsController(ctx *gin.Context) {
	if path := ctx.Param("path"); path != "" {
		target := filepath.Join(getUploadsDir(), path)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attacthment; filename="+path)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(target)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

func getUploadsDir() (uploads string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	uploads = filepath.Join(dir, "uploads")
	return
}
