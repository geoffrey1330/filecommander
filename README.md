# File Commander

File Commander is a command-line tool for performing various file operations such as `DELETE`, `COPY`, `MOVE`,`READ`, `SEARCH`, `LIST` with the terminal.

## Installation

To install File Commander, follow this [Guide](https://github.com/geoffrey1330/filecommander/blob/main/Installation.md)

## Getting Repo

follow these steps:

1. Make sure you have Go installed on your machine.

```shell
# Check Go version
go version
```

or visit https://go.dev/ to install Go on your local machine.

2. Clone the repository or download the source code.

```shell
git clone https://github.com/geoffrey1330/filecommander.git
```

3. Navigate to the project directory.

```shell
cd filecommander
```

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

#### Open

Open a file in the default editor.

```shell
filecommander open <filename>
```

# File commander Contribution Guide

Welcome to the contribution guide for File commander! We appreciate your interest in contributing to our open-source Go project. This guide will help you understand the process of contributing and how to get started.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Contributing Guidelines](#contributing-guidelines)
   - [Code Style](#code-style)
   - [Pull Requests](#pull-requests)
   - [Branching](#branching)
   - [Commit Messages](#commit-messages)
   - [Testing](#testing)
   - [Documentation](#documentation)
3. [Contact](#contact)

## Getting Started

To contribute to File commander, follow these steps:

1. Fork the File commander repository to your GitHub account.
2. Clone the forked repository to your local development environment.
3. Set up the project and its dependencies by following the instructions in the project's README file.
4. Create a new branch for your changes.

Congratulations! You are now ready to start making contributions.

## Contributing Guidelines

### Code Style

- Follow the Go community's [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) to maintain a consistent code style throughout the project.
- Use descriptive variable and function names to improve code readability.

### Pull Requests

- Before submitting a pull request, ensure your code compiles and passes all tests.
- Include a clear and concise description of the changes made in your pull request.
- Reference any related issues by mentioning their numbers in the pull request description.
- Be responsive to any feedback or suggestions given during the code review process.

### Branching

- Create a new branch for each feature, bug fix, or enhancement you are working on.
- Use a descriptive name for your branch that reflects the purpose of the changes.

### Commit Messages

- Write clear and concise commit messages that describe the changes made in the commit.
- Start the commit message with a verb in the imperative form (e.g., "Fix," "Add," "Update") to indicate the action performed.

### Testing

- Write tests to cover your changes and ensure they pass before submitting a pull request.
- Follow the project's existing testing conventions and patterns.
- If applicable, update or create new documentation for any changes to existing features or the addition of new features.

### Documentation

- Contribute to the project's documentation to help other developers understand the code and its usage.
- Keep the documentation up-to-date with the latest changes made to the project.

## Contact

If you have any questions or need assistance with the project, you can reach out to us via:

- Email: [israelgeoffrey13@gmail.com](mailto:israelgeoffrey13@gmail.com)

Thank you for your interest and contributions to File commander! We appreciate your support in making this open-source project even better.
