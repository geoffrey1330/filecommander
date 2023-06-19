# File Commander

File Commander is a command-line tool for performing various file operations.

## Installation

To install File Commander, follow these steps:

1. Make sure you have Go installed on your machine.
2. Clone the repository or download the source code.
3. Navigate to the project directory.
4. Run the following command to build the executable:

```shell
go build -o filecommander main.go search.go utils.go
```
This will generate an executable named filecommander.

## Usage
File Commander supports the following commands:

#### Create
Create a new file.

```shell
filecommander create <filename>
```
#### Read
Read the content of a file.

```shell
filecommander read <filename>
```

#### Write
Write content to a file.

```shell
filecommander write <filename> <content>
```

#### Delete
Delete a file.

```shell
filecommander delete <filename>
```
#### List
List all files in a directory.

```shell
filecommander list <directory>
```
#### Copy
Copy a file to a new destination.

```shell
filecommander copy <srcfile> <destfile>
```
#### Move
Move a file to a new destination.

```shell
filecommander move <srcfile> <destfile>
```

#### Search
Search for files in a directory based on a given search term.

```shell
filecommander search <directory> <filename>
```
OR 
```shell
filecommander search <directory> <filename extension>
```