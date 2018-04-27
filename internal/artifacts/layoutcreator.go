package artifacts

// LayoutCreator should be implemented by project layouts to create project files
type LayoutCreator interface {
	Create(outputFolder string) error
}
