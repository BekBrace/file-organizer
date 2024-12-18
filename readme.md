# GO programming [Think like a programmer]

# Creating an application that organizes your files into different folders based on each file extension.

# Remember! You can improve upon this app.

## 1. Project setup :)

1. **Create the Project Directory:**
   ```bash
   mkdir file-organizer
   cd file-organizer
   ```

2. **Initialize the Go Module:**
   ```bash
   go mod init file-organizer
   ```

3. **Create the Main File:**
   ```bash
   new-item main.go
   ```

---

## 2. Coding the app :)

### Step 1: Setting Up the Main Function

Open `main.go` and start with the basic structure:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath" // Import this for working with file paths
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [directory path]")
		return
	}
	dirPath := os.Args[1]
	fmt.Println("Organizing files in:", dirPath)
}
```

**Explanation:**
- We check for a directory path argument and print it.

### Step 2: Reading the Directory

Add a function to read files in the specified directory:

```go
func getFiles(dirPath string) ([]os.FileInfo, error) {
	files, err := os.ReadDir(dirPath) // Reads the directory
	if err != nil {
		return nil, err // If there's an error, return it
	}
	return files, nil // Return the list of files
}
```

**Update the `main` function:**

```go
files, err := getFiles(dirPath)
if err != nil {
	fmt.Println("Error reading directory:", err)
	return
}

for _, file := range files {
	fmt.Println(file.Name()) // Print each file's name
}
```

### Step 3: Categorizing Files

Add a function to determine file categories:

```go
func getCategory(fileName string) string {
	ext := filepath.Ext(fileName) // Get file extension
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
```

Update the loop in `main` to use this function:

```go
for _, file := range files {
	if !file.IsDir() { // Skip directories
		category := getCategory(file.Name())
		fmt.Printf("File: %s -> Category: %s\n", file.Name(), category)
	}
}
```

### Step 4: Creating Category Folders

Add a function to create directories:

```go
func createDir(dirPath, category string) error {
	categoryPath := filepath.Join(dirPath, category) // Create category-specific path
	return os.MkdirAll(categoryPath, os.ModePerm) // Create the folder if it doesn't exist
}
```

Integrate this in the loop:

```go
categoryPath := filepath.Join(dirPath, category)
err := createDir(dirPath, category)
if err != nil {
	fmt.Println("Error creating folder:", err)
	continue // Skip to the next file if there’s an error
}
```

### Step 5: Moving Files

Add a function to move files:

```go
func moveFile(filePath, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath)) // Get destination path
	return os.Rename(filePath, destPath) // Move the file
}
```

Complete the loop:

```go
srcPath := filepath.Join(dirPath, file.Name()) // Get source file path
err := moveFile(srcPath, categoryPath) // Move the file to its category folder
if err != nil {
	fmt.Println("Error moving file:", err)
}

```

## 6. Running the Program

1. Build and run the program:
   ```bash
   go run main.go /path/to/directory
   ```

2. Observe the organized files in their respective folders.

---

## 7. Expanding the Project

1. Add support for more file types.
2. Implement error logging.
3. Create a configuration file for custom categories.
