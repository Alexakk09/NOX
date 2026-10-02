package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ValidatePDF(filename string) error {
	ext := filepath.Ext(filename)

	if ext != ".pdf" && ext != ".PDF" {
		return fmt.Errorf("only PDF files are allowed")
	}

	return nil
}

func SaveAndExtractPDF(filename string, file io.Reader) (string, error) {
	if err := ValidatePDF(filename); err != nil {
		return "", err
	}

	uploadDir := "uploads"

	err := os.MkdirAll(uploadDir, 0755)
	if err != nil {
		return "", err
	}

	safeName := filepath.Base(filename)
	filePath := filepath.Join(uploadDir, safeName)

	output, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(output, file)
	if err != nil {
		output.Close()
		return "", err
	}

	err = output.Close()
	if err != nil {
		return "", err
	}

	text, err := ExtractPDFText(filePath)
	if err != nil {
		return "", err
	}

	err = os.Remove(filePath)
	if err != nil {
		return "", err
	}

	return text, nil
}