package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func OpenInVSCode(path string) error {
	cmd := exec.Command("code", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func HandleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
}
