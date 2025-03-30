package templates

// Template defines the interface for project templates
type Template interface {
	// Name returns the name of the template
	Name() string
	// Description returns a description of the template
	Description() string
	// CreateStructure creates the project structure at the given path
	CreateStructure(projectPath, projectName string) error
}

// BaseTemplate provides common functionality for templates
type BaseTemplate struct {
	name        string
	description string
}

// Name returns the name of the template
func (t *BaseTemplate) Name() string {
	return t.name
}

// Description returns a description of the template
func (t *BaseTemplate) Description() string {
	return t.description
}
