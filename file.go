package tplr

import (
	"fmt"
	"io"
	"os"
)

// GetFileReader returns a Reader for the given filename, or '-' for stdin
func GetFileReader(filename string) (io.Reader, error) {
	if filename == "-" || filename == "" {
		return os.Stdin, nil
	}

	if !FileExists(filename) {
		return nil, fmt.Errorf("'%s' doesn't exist or isn't a regular file", filename)
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open '%s': %w", filename, err)
	}

	return f, nil
}

// GetFileWriter returns a Writer for the given filename, or '-' for stdout
func GetFileWriter(filename string, force bool) (io.Writer, error) {
	if filename == "-" || filename == "" {
		return os.Stdout, nil
	}

	if FileExists(filename) && !force {
		return nil, fmt.Errorf("'%s' already exists - wont overwrite without force option", filename)
	}

	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create '%s': %w", filename, err)
	}

	return f, nil
}

// FileExists checks that a file exists and is a regular file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsRegular()
}
