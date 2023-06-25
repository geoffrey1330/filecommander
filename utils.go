package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func createFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		// File exists, prompt user to overwrite
		answer := ""
		for answer != "y" && answer != "n" {
			fmt.Print("File already exists. Do you want to overwrite it? (y/n): ")
			_, err := fmt.Scanln(&answer)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if answer == "n" {
			return
		}
	}

	// Create the file
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

func openFile(filename string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// On Windows, use "code.cmd"
		cmd = exec.Command("code.cmd", filename)
		handleCmdOs(cmd)
		err := cmd.Run()

		if err != nil {
			cmd = exec.Command("notepad", filename)
			handleCmdOs(cmd)
			runErr := cmd.Run()

			if runErr != nil {
				fmt.Println(runErr)
			}
		}

	} else {
		// implement such that the default editor is opened
		cmd = exec.Command("editor", filename)
		handleCmdOs(cmd)
		runErr := cmd.Run()

		if runErr != nil {
			fmt.Println(runErr)
		}
	}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func handleCmdOs(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
