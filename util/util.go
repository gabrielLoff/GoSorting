package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFolderIfNotExists(folderPath string) error {
	// Check if the folder already exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Folder doesn't exist, create it
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println("Folder created successfully.")
	} else {
		fmt.Println("Folder already exists.")
	}

	return nil
}

func CleanImages(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if isImageFile(file.Name()) {
			filePath := filepath.Join(folderPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", filePath, err)
			} else {
				fmt.Printf("Deleted file: %s\n", filePath)
			}
		}
	}

	return nil
}

func isImageFile(filename string) bool {
	// You can customize this function based on the image file extensions you want to consider
	extensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	lowercaseFilename := strings.ToLower(filename)

	for _, ext := range extensions {
		if strings.HasSuffix(lowercaseFilename, ext) {
			return true
		}
	}

	return false
}
