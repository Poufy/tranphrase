package controllers

import (
	"fmt"
	"os"

	"tranphrase/services"

	"github.com/gin-gonic/gin"
)

func UploadTranslateAndDownloadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "File upload failed: %v", err)
		return
	}

	filePath := fmt.Sprintf("temp/%s", file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.String(500, "Failed to save uploaded file: %v", err)
		return
	}

	defer os.Remove(filePath)

	translatedFilePath, err := services.ProcessExcelFile(filePath)
	if err != nil {
		c.String(500, "Failed to process the Excel file: %v", err)
		return
	}

	defer os.Remove(translatedFilePath)

	c.FileAttachment(translatedFilePath, "translated.xlsx")
}
