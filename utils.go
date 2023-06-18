package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created:", filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}

func writeFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content written to file:", filename)
}

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted:", filename)
}

func listFiles(directory string) {
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", path)
			return nil
		}

		if !info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})
}

func copyFile(src, dest string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied from", src, "to", dest)
}

func moveFile(src, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		fmt.Println("Error moving file:", err)
		return
	}

	fmt.Println("File moved from", src, "to", dest)
}
