package routes

import (
	"tranphrase/controllers"
	"tranphrase/middleware"

	"github.com/gin-gonic/gin"
)

func TranslateRoutes(r *gin.RouterGroup) {
	r.POST("/translate", middleware.ExcelFileExtensionCheck(), controllers.UploadTranslateAndDownloadExcel)
}
