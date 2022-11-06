// Package clipboard read/write on clipboard
package clipboard

// Read read string from clipboard
func Read() (string, error) {
	return read()
}

// Write write string to clipboard
func Write(text string) error {
	return write(text)
}
