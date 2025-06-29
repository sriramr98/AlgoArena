package problems

import (
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"os"
	"path"
	"slices"
	"strings"
)

var FilesToIgnore = []string{
	"schema.json",
}

func ValidateProblem(problemsPath string) {
	files, err := os.ReadDir(problemsPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The specified problem path does not exist:", problemsPath)
		} else {
			fmt.Println("Error reading the problem directory:", err)
		}
		return
	}

	schemaFile, err := os.Open(problemsPath + "/schema.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Schema fileInfo not found in the problem directory:", problemsPath)
		} else {
			fmt.Println("Error opening schema fileInfo:", err)
		}
		return
	}
	defer schemaFile.Close()

	schemaBytes, err := io.ReadAll(schemaFile)
	if err != nil {
		fmt.Println("Error reading schema fileInfo:", err)
		return
	}
	schema := string(schemaBytes)
	schemaLoader := gojsonschema.NewStringLoader(schema)

	invalidFiles := make([]string, 0)

	println("Validating problem at:", problemsPath)
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		if !strings.HasSuffix(fileInfo.Name(), ".json") {
			continue
		}
		if slices.Contains(FilesToIgnore, fileInfo.Name()) {
			continue
		}

		if !isProblemValid(fileInfo, problemsPath, schemaLoader) {
			invalidFiles = append(invalidFiles, fileInfo.Name())
			continue
		}
	}

	fmt.Printf("\n\n")
	fmt.Printf("Validation complete. %d problems were invalid out of %d.\n", len(invalidFiles), len(files))
	if len(invalidFiles) > 0 {
		fmt.Println("Invalid files:")
		for _, file := range invalidFiles {
			fmt.Println("-", file)
		}
	} else {
		fmt.Println("All files are valid.")
	}
}

func isProblemValid(fileInfo os.DirEntry, problemsPath string, schemaLoader gojsonschema.JSONLoader) bool {
	fmt.Println("------------------------------------")
	file, err := os.Open(problemsPath + "/" + fileInfo.Name())
	if err != nil {
		fmt.Println("Error opening file:", fileInfo.Name(), "-", err)
		return false
	}

	if file == nil {
		fmt.Println("Error opening file:", fileInfo.Name())
		return false
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading fileInfo:", file.Name(), "-", err)
		return false
	}

	documentLoader := gojsonschema.NewStringLoader(string(fileBytes))
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println("Error validating fileInfo:", file.Name(), "-", err)
		return false
	}

	if !result.Valid() {
		// Get last part of the fileInfo path segment
		fmt.Printf("Problem %s is invalid\n", path.Base(file.Name()))
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	fmt.Printf("Problem %s is valid\n", path.Base(file.Name()))
	fmt.Println("------------------------------------")
	return true
}
