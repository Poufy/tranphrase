package services

import (
	"path/filepath"
	"strings"
)

func ValidExcelExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".xlsx" || ext == ".xls"
}
