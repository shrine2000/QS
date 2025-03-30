package templates

import (
	"fmt"
	"os"
	"path/filepath"
)

type GoTemplate struct {
	BaseTemplate
}

func NewGoTemplate() *GoTemplate {
	return &GoTemplate{
		BaseTemplate: BaseTemplate{
			name:        "go",
			description: "Go project template with standard layout",
		},
	}
}

func (t *GoTemplate) CreateStructure(projectPath, projectName string) error {
	// Create cmd directory
	cmdDir := filepath.Join(projectPath, "cmd", projectName)
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		return fmt.Errorf("failed to create cmd directory: %w", err)
	}

	// Create main.go
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`
	if err := os.WriteFile(filepath.Join(cmdDir, "main.go"), []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	// Create internal directory
	internalDir := filepath.Join(projectPath, "internal", "pkg")
	if err := os.MkdirAll(internalDir, 0755); err != nil {
		return fmt.Errorf("failed to create internal directory: %w", err)
	}

	// Create pkg directory
	pkgDir := filepath.Join(projectPath, "pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return fmt.Errorf("failed to create pkg directory: %w", err)
	}

	// Create test directory
	testDir := filepath.Join(projectPath, "test")
	if err := os.MkdirAll(testDir, 0755); err != nil {
		return fmt.Errorf("failed to create test directory: %w", err)
	}

	// Create .gitignore
	gitignoreContent := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories to remove
vendor/

# Go workspace file
go.work

# IDE specific files
.idea/
.vscode/
*.swp
*.swo

# OS specific files
.DS_Store
Thumbs.db`
	if err := os.WriteFile(filepath.Join(projectPath, ".gitignore"), []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Create go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.21`, projectName)
	if err := os.WriteFile(filepath.Join(projectPath, "go.mod"), []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	return nil
}
