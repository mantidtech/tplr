package tplr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

// ReadDataFile reads the given file into a map of interfaces
func ReadDataFile(filename string) (map[string]interface{}, error) {
	vars := make(map[string]interface{})

	dr, err := GetFileReader(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %w", err)
	}

	d, err := ioutil.ReadAll(dr)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file %s: %w", filename, err)
	}

	err = json.Unmarshal(d, &vars)
	if err != nil {
		return nil, fmt.Errorf("failed to parse data file %s: %w", filename, err)
	}

	return vars, nil
}

// ReadStringsAsFile returns the given set of strings as an io.Reader
func ReadStringsAsFile(s ...string) (res io.Reader, err error) {
	var b bytes.Buffer
	b.WriteString(strings.Join(s, " "))
	res = &b

	if b.Len() == 0 {
		err = fmt.Errorf("empty input")
	}

	return res, err
}
