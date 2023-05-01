package middleware

import (
	"net/http"
	"tranphrase/services"

	"github.com/gin-gonic/gin"
)

func ExcelFileExtensionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
			c.Abort()
			return
		}

		// Check file extension
		if !services.ValidExcelExtension(file.Filename) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
			c.Abort()
			return
		}

		c.Next()
	}
}
