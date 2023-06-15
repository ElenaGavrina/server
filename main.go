package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		filename, err := uuid.Parse(filepath.Base(file.Filename))
		if err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"fileName": filename,
		})
	})
	
	r.Run(":8080")
}
