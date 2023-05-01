package services

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/xuri/excelize/v2"
)

func translateText(text, source, target string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: source,
			To:   target,
		},
	)
	if err != nil {
		return "", err
	}

	return translated, nil
}

func ProcessExcelFile(filePath string) (string, error) {
	// Read the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return "", err
	}

	// Get the first sheet name
	sheet := f.GetSheetName(0)

	// Loop through the rows and translate the first column
	rows, _ := f.GetRows(sheet)
	for i, row := range rows {
		if len(row) > 0 {
			translatedText, err := translateText(row[0], "en", "ar") // Replace "es" with the target language code
			if err != nil {
				return "", err
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", i+1), translatedText)
		}
	}

	// Save the translated Excel file
	translatedFilePath := "temp/translated.xlsx"
	writeExcelFile(f, translatedFilePath)

	return translatedFilePath, nil
}

func writeExcelFile(file *excelize.File, path string) error {
	err := file.SaveAs(path)
	if err != nil {
		return err
	}
	return nil
}
