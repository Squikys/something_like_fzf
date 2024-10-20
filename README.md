# Something Like FZF in Go

This is a Go program that recursively searches for files in a specified directory based on a user-input search term. It updates the terminal display to show the matching files and allows for navigation using arrow keys. The user can interactively search and filter the displayed results in real-time.

## Features

- Recursively walks through directories to search for files.
- Filters files by matching filenames with a given search term.
- Allows real-time input from the user to update the search term.
- Uses arrow keys for navigating through the displayed files.
- Displays the search results in a paginated format, with customizable limits on how many files to display.

## Usage

To run the program, use the following command:

```bash
go run main.go [directory] [display_limit]
```
## Dependencies

This program uses the following third-party libraries:
-mattn/go-tty: Used for capturing keyboard inputs.

```bash
go get github.com/mattn/go-tty
```
## Bugs and Optimization
I would appreciate if someone can help me optimize it and fix the bugs. Any contributions, suggestions, or improvements are welcome!
