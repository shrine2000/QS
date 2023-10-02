package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func main() {
	var folderPath string
	var folderName string
	var projectName string

	flag.StringVar(&folderPath, "path", "", "Folder path (leave empty for the current directory)")
	flag.StringVar(&folderName, "name", "", "Folder name")
	flag.StringVar(&projectName, "project", "", "Project name")
	flag.Parse()

	if folderName == "" || projectName == "" {
		fmt.Println("Usage: qs -name <folder_name> -project <project_name>")
		return
	}

	if folderPath == "" {
		currentUser, err := user.Current()
		if err != nil {
			handleError("Error getting user information", err)
			return
		}
		folderPath = filepath.Join(currentUser.HomeDir, "Documents")
	}

	fullPath := filepath.Join(folderPath, folderName)

	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		handleError("Error creating folder", err)
		return
	}
	fmt.Println("Folder created successfully.")

	err = os.Chdir(fullPath)
	if err != nil {
		handleError("Error changing directory", err)
		return
	}

	cmd := exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		handleError("Error initializing Git", err)
		return
	}
	fmt.Println("Git repository initialized.")

	readmeContent := "# " + projectName

	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		handleError("Error creating README.md", err)
		return
	}
	fmt.Println("README.md created successfully.")

	cmd = exec.Command("code", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		handleError(fmt.Sprintf("Error opening %s in VS Code", projectName), err)
		return
	}
	fmt.Printf("Opening %s in VS Code\n", projectName)
}

func handleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
}
