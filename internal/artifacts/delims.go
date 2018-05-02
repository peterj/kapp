package artifacts

// Delims holds the template delimiters for the project file (e.g. {{ }} or [[ ]])
type Delims struct {
	Left  string
	Right string
}

// NewDelims creates new delimiters instance
func NewDelims(left, right string) *Delims {
	return &Delims{
		Left:  left,
		Right: right,
	}
}

// NewDefaultDelims creates new default delimiters ("{{ }}")
func NewDefaultDelims() *Delims {
	return &Delims{
		Left:  "{{",
		Right: "}}",
	}
}
