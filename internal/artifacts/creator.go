package artifacts

// Creator defines functions for creating artifacts from template files
type Creator interface {
	Create() ([]byte, error)
}
