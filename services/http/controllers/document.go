package controllers

import (
	"github.com/nugrohosam/gofilestatic/helpers"
	requests "github.com/nugrohosam/gofilestatic/services/http/requests/v1"
	"github.com/nugrohosam/gofilestatic/usecases"

	"github.com/gin-gonic/gin"
)

// DocumentHandlerGetFile is use
func DocumentHandlerGetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, _ := usecases.GetFile(file, "document")

		c.File(filePath)
	}
}

// DocumentHandlerUpload is use
func DocumentHandlerUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var documentUpload requests.DocumentUpload
		err := c.ShouldBind(&documentUpload)
		if err != nil {
			panic(err.Error())
		}

		imageBuff, err := helpers.ReadFileRequest(documentUpload.File)
		filePath := usecases.StoreFile(imageBuff, documentUpload.FilePath, documentUpload.File.Filename, "document")

		c.JSON(200, gin.H{
			"status":    true,
			"file_name": filePath,
		})
	}
}

// DocumentHandlerPdfInImage is use
func DocumentHandlerPdfInImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, _ := usecases.GetFilePdfInImage(file, "document")

		c.File(filePath)
	}
}

// DocumentHandlerDocInPdf is use
func DocumentHandlerDocInPdf() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, _ := usecases.GetFilePdfInImage(file, "document")

		c.File(filePath)
	}
}
