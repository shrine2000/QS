package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the folder path (or '.' for the current directory): ")
	folderPath, _ := reader.ReadString('\n')
	folderPath = strings.TrimSpace(folderPath)

	fmt.Print("Enter the folder name: ")
	folderName, _ := reader.ReadString('\n')
	folderName = strings.TrimSpace(folderName)

	fullPath := folderPath + "/" + folderName

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

	fmt.Print("Enter project name: ")
	projectName, _ := reader.ReadString('\n')
	readmeContent := "# " + projectName

	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error creating README.md:", err)
		return
	}
}
