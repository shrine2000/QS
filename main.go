package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	var folderPath string
	var folderName string
	var projectName string
	var description string

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

	description, err = getUserInput("Enter a description for the project:")
	if err != nil {
		handleError("Error getting user input", err)
		return
	}

	readmeContent := "# " + projectName + "\n\n" + description + "\n\n## License\nThis project is licensed under the [MIT License](./LICENSE)."

	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		handleError("Error creating README.md", err)
		return
	}
	fmt.Println("README.md created successfully.")

	licenseContent := generateMITLicense()
	err = os.WriteFile("LICENSE", []byte(licenseContent), 0644)
	if err != nil {
		handleError("Error creating LICENSE file", err)
		return
	}
	fmt.Println("LICENSE file created successfully.")

	ides := findIDEs()
	if len(ides) > 0 {
		fmt.Println("Installed IDEs:")
		i := 1
		for ideExec, ideName := range ides {
			fmt.Printf("%d. %s (%s)\n", i, ideName, ideExec)
			i++
		}

		fmt.Printf("Enter the number of the IDE to open: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		choice := strings.TrimSpace(input)
		if choice == "" {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		print(ides)
		ideExecutable, exists := ides[choice]
		if !exists {
			fmt.Println("Invalid choice.")
			return
		}

		fmt.Printf("Opening %s...\n", ides[fmt.Sprintf("ide%d", choice)])
		openIDE(fmt.Sprintf("ide%d", choice), fullPath)

		fmt.Printf("Opening %s...\n", ideExecutable)
	} else {
		fmt.Println("No known IDEs found.")
	}

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

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func generateMITLicense() string {
	return `MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`
}

func findIDEs() map[string]string {
	ideExecutables := map[string]string{
		"code":    "Visual Studio Code",
		"pycharm": "PyCharm",
		"subl":    "Sublime Text",
		"atom":    "Atom",
	}

	installedIDEs := make(map[string]string)

	for ideExec := range ideExecutables {
		path, err := exec.LookPath(ideExec)
		if err == nil {
			installedIDEs[ideExec] = path
		}
	}

	return installedIDEs
}

func openIDE(ideExecutable, projectPath string) {
	cmd := exec.Command(ideExecutable, projectPath)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error opening IDE: %v\n", err)
	}
}
