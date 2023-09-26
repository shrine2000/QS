package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var folderPath string
	var folderName string
	var projectName string

	flag.StringVar(&folderPath, "path", ".", "Folder path (or '.' for the current directory)")
	flag.StringVar(&folderName, "name", "", "Folder name")
	flag.StringVar(&projectName, "project", "", "Project name")
	flag.Parse()

	if folderName == "" || projectName == "" {
		fmt.Println("Usage: qs -name <folder_name> -project <project_name>")
		return
	}

	fullPath := filepath.Join(folderPath, folderName)

	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		fmt.Println("Error making folder:", err)
		return
	}
	fmt.Println("Folder created successfully.")

	err = os.Chdir(fullPath)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	cmd := exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error initializing Git:", err)
		return
	}
	fmt.Println("Git repository initialized.")

	readmeContent := "# " + projectName

	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error creating README.md:", err)
		return
	}
	fmt.Println("README.md created successfully.")
}
