package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"qs/internal/config"
	"qs/pkg/templates"
)

type Creator struct {
	projectName string
	basePath    string
	templates   map[string]templates.Template
	config      *config.Config
}

func NewCreator(projectName, basePath string, cfg *config.Config) *Creator {
	c := &Creator{
		projectName: projectName,
		basePath:    basePath,
		templates:   make(map[string]templates.Template),
		config:      cfg,
	}

	// Register available templates
	c.templates["python"] = templates.NewPythonTemplate()
	c.templates["go"] = templates.NewGoTemplate()

	return c
}

func (c *Creator) Create() error {
	if err := c.createBaseDirectory(); err != nil {
		return err
	}

	if err := c.initializeGit(); err != nil {
		return err
	}

	return nil
}

func (c *Creator) CreateProject(lang string) error {
	template, exists := c.templates[lang]
	if !exists {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	projectPath := filepath.Join(c.basePath, c.projectName)
	if err := template.CreateStructure(projectPath, c.projectName); err != nil {
		return fmt.Errorf("failed to create project structure: %w", err)
	}

	if lang == "python" {
		if err := c.createVirtualEnvironment(); err != nil {
			fmt.Printf("Warning: Could not create virtual environment: %v\n", err)
		}
	}

	return nil
}

func (c *Creator) createBaseDirectory() error {
	fullPath := filepath.Join(c.basePath, c.projectName)
	return os.MkdirAll(fullPath, 0755)
}

func (c *Creator) initializeGit() error {
	cmd := exec.Command("git", "init")
	cmd.Dir = filepath.Join(c.basePath, c.projectName)
	if err := cmd.Run(); err != nil {
		return err
	}

	if c.config.GitUser != "" {
		cmd = exec.Command("git", "config", "user.name", c.config.GitUser)
		cmd.Dir = filepath.Join(c.basePath, c.projectName)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if c.config.GitEmail != "" {
		cmd = exec.Command("git", "config", "user.email", c.config.GitEmail)
		cmd.Dir = filepath.Join(c.basePath, c.projectName)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Creator) createVirtualEnvironment() error {
	cmd := exec.Command("python3", "-m", "venv", "venv")
	cmd.Dir = filepath.Join(c.basePath, c.projectName)
	return cmd.Run()
}
