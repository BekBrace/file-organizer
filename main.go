// Name: File Organizer
// Version: 1.0
// Objective: Organizing files and categorizing by extension and moving each file in a separate folder.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// setting up the main function
func main() {
	// 1- Accept the directory path from the user
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [directory path]")
		return
	}
	// Get the directory path from the command line argument
	dirPath := os.Args[1]
	fmt.Println("Organizing files in:", dirPath)

	// creating a function to read/get the files from the directory
	files, err := readFiles(dirPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	// Loop through each file and organize them by category

	for _, file := range files {
		if !file.IsDir() {
			category := getCategory(file.Name())
			categoryPath := filepath.Join(dirPath, category)
			err := createDir(dirPath, category)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}
			// Move the file/s to the category folder
			srcPath := filepath.Join(dirPath, file.Name())
			err = moveFile(srcPath, categoryPath)
			if err != nil {
				fmt.Println("Error moving the file for X reason:", err)
			} else {
				fmt.Printf("Moved: %s -> %s\n", file.Name(), category)
			}
		}
	}
}

// function to read files from the directory
func readFiles(dirPath string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// function to organize files into different folders based on their extensions
func getCategory(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return "Images"
	case ".mov", ".mkv", ".mp4":
		return "Videos"
	case ".txt", ".doc", ".pdf":
		return "Documents"
	case ".js", ".jsx", ".py", ".hs":
		return "Programming"
	default:
		return "Others"
	}
}

// function to create a directory for a category
func createDir(dirPath, category string) error {
	categoryPath := filepath.Join(dirPath, category)
	return os.MkdirAll(categoryPath, os.ModePerm)
}

// function to move from source path to destination path
func moveFile(filePath, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath))
	return os.Rename(filePath, destPath)
}
