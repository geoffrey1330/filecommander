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
	if err := verifyCodeCommand(); err != nil {
		return err
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// On Windows, use "code.cmd"
		cmd = exec.Command("code.cmd", filename)
	} else {
		cmd = exec.Command("code", filename)
	}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func verifyCodeCommand() error {
	_, err := exec.LookPath("code")
	if err != nil {
		fmt.Println("code command not found. Attempting to install...")
		err = installCodeCommand()
		if err != nil {
			return fmt.Errorf("failed to verify and install code command: %s", err)
		}
		fmt.Println("code command installed successfully.")
	}
	return nil
}

func installCodeCommand() error {
	var command string
	var args []string
	if runtime.GOOS == "windows" {
		command = "powershell"
		args = []string{"-Command", "& { Invoke-WebRequest -Uri 'https://go.microsoft.com/fwlink/?LinkID=760868' -OutFile 'vscode.zip' }"}
	} else {
		command = "bash"
		args = []string{"-c", "curl -o vscode.zip -L https://go.microsoft.com/fwlink/?LinkID=760868 && unzip vscode.zip && sudo mv 'VSCode.app' '/usr/local/bin/code'"}
	}

	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install code command: %s", err)
	}

	return nil
}
