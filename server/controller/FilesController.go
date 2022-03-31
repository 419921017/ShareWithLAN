package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func FilesController(ctx *gin.Context) {
	file, err := ctx.FormFile("raw")
	if err != nil {
		return
	}
	executable, err := os.Executable()
	if err != nil {
		return
	}
	dir := filepath.Dir(executable)
	fileName := uuid.New().String()
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		return
	}
	fullpath := filepath.Join(uploads, fileName+filepath.Ext(file.Filename))
	fileErr := ctx.SaveUploadedFile(file, filepath.Join(dir, fullpath))
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	ctx.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
