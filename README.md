# QuickStart (QS)

A command-line tool to quickly create new project structures with best practices.

## Features

- Create new projects with predefined structures
- Support for multiple programming languages (Python, Go)
- Git repository initialization
- VS Code integration
- Virtual environment setup for Python projects
- Customizable project templates
- Configuration management

## Installation

### Prerequisites
- Go 1.21 or later
- Git
- VS Code (optional)

### Local Installation

1. Clone the repository:
```bash
git clone https://github.com/shrine2000/QS
cd QS
```

2. Build and install:
```bash
make install
```

3. Add ~/.local/bin to your PATH (if not already added):
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
# or for zsh:
# echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
```

4. Reload your shell configuration:
```bash
source ~/.bashrc
# or for zsh:
# source ~/.zshrc
```

### Uninstallation
```bash
make uninstall
```

### Clean Build Files
```bash
make clean
```

## Usage

Basic usage:
```bash
qs my-project
```

With flags:
```bash
qs -lang python -desc "My awesome project" my-project
```

Available flags:
- `-lang`: Project language (python, go)
- `-desc`: Project description
- `-git`: Initialize git repository (default: true)
- `-vscode`: Open in VS Code (default: true)
- `-venv`: Create virtual environment for Python projects (default: true)
- `-config`: Configure QS settings

## Configuration

The configuration file is stored at `~/.qs/config.json`. Run `qs -config` to set up your preferences:
- Default project path (default: ~/Documents)
- Default project language
- Git username and email
- VS Code path

Example config.json:
```json
{
  "base_path": "/home/user/Documents",
  "default_lang": "python",
  "git_user": "John Doe",
  "git_email": "john@example.com"
}
```

## Project Templates

### Python Project Structure
```
project_name/
├── project_name/        # Main source code
│   ├── __init__.py    
│   ├── app.py         # Main application code
├── tests/             # Unit tests
│   ├── __init__.py    
│   ├── test_app.py    # Test cases
├── .gitignore         # Git ignore rules
├── requirements.txt   # Python dependencies
├── setup.py          # Package configuration
└── README.md         # Project documentation
```

### Go Project Structure
```
project_name/
├── cmd/               # Main applications
│   └── project_name/
│       └── main.go    # Entry point
├── internal/          # Private application code
│   └── pkg/          # Internal packages
├── pkg/              # Public library code
├── test/             # Additional test files
├── .gitignore        # Git ignore rules
├── go.mod            # Go module definition
└── README.md         # Project documentation
```

## Creating New Templates

QS uses a simple interface for templates. To create a new template:

1. Create a new file in `pkg/templates/` (e.g., `rust.go`)
2. Implement the `Template` interface:
```go
type Template interface {
    Name() string
    Description() string
    CreateStructure(projectPath, projectName string) error
}
```
3. Register your template in `internal/project/creator.go`

Example template:
```go
type RustTemplate struct {
    BaseTemplate
}

func NewRustTemplate() *RustTemplate {
    return &RustTemplate{
        BaseTemplate: BaseTemplate{
            name:        "rust",
            description: "Rust project template",
        },
    }
}

func (t *RustTemplate) CreateStructure(projectPath, projectName string) error {
    // Create project structure
    return nil
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
