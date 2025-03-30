package templates

import (
	"fmt"
	"os"
	"path/filepath"
)

type PythonTemplate struct {
	BaseTemplate
}

func NewPythonTemplate() *PythonTemplate {
	return &PythonTemplate{
		BaseTemplate: BaseTemplate{
			name:        "python",
			description: "Python project template with standard layout",
		},
	}
}

func (t *PythonTemplate) CreateStructure(projectPath, projectName string) error {
	// Create main package directory
	packageDir := filepath.Join(projectPath, projectName)
	if err := os.MkdirAll(packageDir, 0755); err != nil {
		return fmt.Errorf("failed to create package directory: %w", err)
	}

	// Create __init__.py in main package
	if err := os.WriteFile(filepath.Join(packageDir, "__init__.py"), []byte(""), 0644); err != nil {
		return fmt.Errorf("failed to create __init__.py: %w", err)
	}

	// Create app.py
	appContent := `def main():
    print("Hello, World!")

if __name__ == "__main__":
    main()`
	if err := os.WriteFile(filepath.Join(packageDir, "app.py"), []byte(appContent), 0644); err != nil {
		return fmt.Errorf("failed to create app.py: %w", err)
	}

	// Create tests directory
	testsDir := filepath.Join(projectPath, "tests")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		return fmt.Errorf("failed to create tests directory: %w", err)
	}

	// Create __init__.py in tests
	if err := os.WriteFile(filepath.Join(testsDir, "__init__.py"), []byte(""), 0644); err != nil {
		return fmt.Errorf("failed to create tests/__init__.py: %w", err)
	}

	// Create test_app.py
	testContent := fmt.Sprintf(`import unittest
from %s.app import main

class TestApp(unittest.TestCase):
    def test_main(self):
        pass

if __name__ == '__main__':
    unittest.main()`, projectName)
	if err := os.WriteFile(filepath.Join(testsDir, "test_app.py"), []byte(testContent), 0644); err != nil {
		return fmt.Errorf("failed to create test_app.py: %w", err)
	}

	// Create .gitignore
	gitignoreContent := `__pycache__/
*.py[cod]
*$py.class
*.so
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
*.egg-info/
.installed.cfg
*.egg
venv/
env/
ENV/
.idea/
.vscode/
*.swp
*.swo
.DS_Store
Thumbs.db`
	if err := os.WriteFile(filepath.Join(projectPath, ".gitignore"), []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Create requirements.txt
	if err := os.WriteFile(filepath.Join(projectPath, "requirements.txt"), []byte("# Add your project dependencies here"), 0644); err != nil {
		return fmt.Errorf("failed to create requirements.txt: %w", err)
	}

	// Create setup.py
	setupContent := fmt.Sprintf(`from setuptools import setup, find_packages

setup(
    name="%s",
    version="0.1.0",
    packages=find_packages(),
    install_requires=[],
    author="Your Name",
    author_email="your.email@example.com",
    description="A short description of your project",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown",
    url="https://github.com/yourusername/%s",
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.6",
)`, projectName, projectName)
	if err := os.WriteFile(filepath.Join(projectPath, "setup.py"), []byte(setupContent), 0644); err != nil {
		return fmt.Errorf("failed to create setup.py: %w", err)
	}

	return nil
}
