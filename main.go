package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: filecommander <command> <args>")
		return
	}

	command := args[0]
	switch command {
	case "create":
		if len(args) < 2 {
			fmt.Println("Usage: filecommander create <filename>")
			return
		}
		createFile(args[1])
	case "read":
		if len(args) < 2 {
			fmt.Println("Usage: filecommander read <filename>")
			return
		}
		readFile(args[1])
	case "write":
		if len(args) < 3 {
			fmt.Println("Usage: filecommander write <filename> <content>")
			return
		}
		writeFile(args[1], args[2])
	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: filecommander delete <filename>")
			return
		}
		deleteFile(args[1])
	case "list":
		if len(args) < 2 {
			fmt.Println("Usage: filecommander list <directory>")
			return
		}
		listFiles(args[1])
	case "copy":
		if len(args) < 3 {
			fmt.Println("Usage: filecommander copy <srcfile> <destfile>")
			return
		}
		copyFile(args[1], args[2])
	case "move":
		if len(args) < 3 {
			fmt.Println("Usage: filecommander move <srcfile> <destfile>")
			return
		}
		moveFile(args[1], args[2])
	case "search":
		if len(args) < 3 {
			fmt.Println("Usage: filecommander search <directory> <filename>")
			return
		}
		searchFiles(args[1], args[2])
	default:
		fmt.Println("Invalid command:", command)
	}
}
