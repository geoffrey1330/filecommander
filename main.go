package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: filecommander <command> <args>\n")
		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		fmt.Fprintf(os.Stderr, "  create <filename>   Create a new file\n")
		fmt.Fprintf(os.Stderr, "  read <filename>     Read the contents of a file\n")
		fmt.Fprintf(os.Stderr, "  write <filename> <content>  Write content to a file\n")
		fmt.Fprintf(os.Stderr, "  delete <filename>   Delete a file\n")
		fmt.Fprintf(os.Stderr, "  list <directory>    List files in a directory\n")
		fmt.Fprintf(os.Stderr, "  copy <srcfile> <destfile>   Copy a file to a new location\n")
		fmt.Fprintf(os.Stderr, "  move <srcfile> <destfile>   Move a file to a new location\n")
		fmt.Fprintf(os.Stderr, "  search <directory> <filename>   Search for a file in a directory\n")
		fmt.Fprintf(os.Stderr, "  open <filename>     Open a file in the default editor\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	// Define flags
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

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
	case "open":
		if len(args) < 2 {
			fmt.Println("Usage: filecommander open <filename>")
			return
		}
		openFile(args[1])
	default:
		fmt.Println("Invalid command:", command)
	}
}
