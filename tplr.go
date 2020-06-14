package tplr

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"

	"github.com/mantidtech/tplr/functions"
)

const templateName = "tplr"

// Load a template from the supplied Reader and create a new Template object
func Load(r io.Reader) (*template.Template, error) {
	tSet := template.New(templateName)
	tSet.Funcs(functions.All(tSet))

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read template: %w", err)
	}

	t, err := tSet.Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	return t, nil
}

// GenerateFromTemplate generates text from the template and data supplied and writes it to the given Writer
func GenerateFromTemplate(w io.Writer, t *template.Template, vars map[string]interface{}) error {
	var err error
	var f bytes.Buffer
	err = t.ExecuteTemplate(&f, templateName, vars)
	if err != nil {
		return fmt.Errorf("failed to apply template: %w", err)
	}

	_, err = w.Write(f.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return err
}
