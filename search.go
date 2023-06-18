package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// File represents information about a file
type File struct {
	FileName   string
	FilePath   string
	Size       int64
	ModifiedAt time.Time
}

func searchFiles(directory, searchTerm string) {
	results, err := SearchFilesByName(directory, searchTerm)
	if err != nil {
		fmt.Println("Error searching files:", err)
		return
	}

	fmt.Printf("Search results for '%s' in directory '%s':\n", searchTerm, directory)
	for _, result := range results {
		fmt.Printf("File Name: %s\n", result.FileName)
		fmt.Printf("File Path: %s\n", result.FilePath)
		fmt.Println("------------")
	}
}

func SearchFilesByName(directory, searchTerm string) ([]File, error) {
	var results []File

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		if strings.Contains(fileName, searchTerm) || strings.HasSuffix(fileName, "."+searchTerm) {
			file := File{
				FileName:   fileName,
				FilePath:   path,
				Size:       info.Size(),
				ModifiedAt: info.ModTime(),
			}
			results = append(results, file)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
