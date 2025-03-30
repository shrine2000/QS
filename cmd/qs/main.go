package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"qs/internal/config"
	"qs/internal/project"
	"qs/internal/utils"
	"strings"
)

var (
	lang       = flag.String("lang", "", "Project language (python, go, etc.)")
	desc       = flag.String("desc", "", "Project description")
	git        = flag.Bool("git", true, "Initialize git repository")
	vscode     = flag.Bool("vscode", true, "Open in VS Code")
	venv       = flag.Bool("venv", true, "Create virtual environment for Python projects")
	showConfig = flag.Bool("config", false, "Configure QS settings")
)

func main() {
	flag.Parse()

	if *showConfig {
		if err := configureQS(); err != nil {
			utils.HandleError("Error configuring QS", err)
			return
		}
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: qs [flags] <project_name>")
		flag.PrintDefaults()
		return
	}

	cfg, err := config.NewConfig()
	if err != nil {
		utils.HandleError("Error initializing configuration", err)
		return
	}

	projectName := os.Args[len(os.Args)-1]
	creator := project.NewCreator(projectName, cfg.BasePath, cfg)

	if err := creator.Create(); err != nil {
		utils.HandleError("Error creating project", err)
		return
	}

	if *desc == "" {
		*desc, err = utils.GetUserInput("Enter a description for the project:")
		if err != nil {
			utils.HandleError("Error getting user input", err)
			return
		}
	}

	if *lang == "" {
		*lang, err = utils.GetUserInput("Enter project language (python/go):")
		if err != nil {
			utils.HandleError("Error getting user input", err)
			return
		}
	}

	switch strings.ToLower(*lang) {
	case "python":
		if err := creator.CreateProject("python"); err != nil {
			utils.HandleError("Error creating Python project structure", err)
			return
		}
	case "go":
		if err := creator.CreateProject("go"); err != nil {
			utils.HandleError("Error creating Go project structure", err)
			return
		}
	default:
		fmt.Printf("Warning: Unknown language %s, skipping project structure\n", *lang)
	}

	readmeContent := "# " + projectName + "\n\n" + *desc
	readmePath := filepath.Join(cfg.BasePath, projectName, "README.md")
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		utils.HandleError("Error creating README.md", err)
		return
	}

	if *vscode {
		projectPath := filepath.Join(cfg.BasePath, projectName)
		if err := utils.OpenInVSCode(projectPath); err != nil {
			fmt.Printf("Warning: Could not open project in VS Code: %v\n", err)
		}
	}
}

func configureQS() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	fmt.Println("QS Configuration")
	fmt.Println("----------------")

	basePath, err := utils.GetUserInput("Enter default project path (press Enter to keep current):")
	if err != nil {
		return err
	}
	if basePath != "" {
		cfg.BasePath = basePath
	}

	defaultLang, err := utils.GetUserInput("Enter default project language (press Enter to keep current):")
	if err != nil {
		return err
	}
	if defaultLang != "" {
		cfg.DefaultLang = defaultLang
	}

	gitUser, err := utils.GetUserInput("Enter Git username (press Enter to keep current):")
	if err != nil {
		return err
	}
	if gitUser != "" {
		cfg.GitUser = gitUser
	}

	gitEmail, err := utils.GetUserInput("Enter Git email (press Enter to keep current):")
	if err != nil {
		return err
	}
	if gitEmail != "" {
		cfg.GitEmail = gitEmail
	}

	return cfg.Save()
}
