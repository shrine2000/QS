package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const fileStructurePath = "data/input.txt"

func createFoldersAndFiles(filePath string) {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			sub := strings.Split(line, "/")
			newPath := ""
			for _, s := range sub {
				if strings.Contains(s, ".") && s != "" {
					// Create files with the full path, but only if they don't already exist
					filePath := filepath.Join(newPath, s)
					_, err := os.Stat(filePath)
					if os.IsNotExist(err) {
						newFile, err := os.Create(filePath)
						if err != nil {
							panic(err)
						}
						newFile.Close()
					}
				} else {
					newPath = filepath.Join(newPath, s)
					newPath += "/"
					// Create directories, but only if they don't already exist
					_, err := os.Stat(newPath)
					if os.IsNotExist(err) {
						err := os.Mkdir(newPath, 0755)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
}

func main() {
	createFoldersAndFiles(fileStructurePath)
}
