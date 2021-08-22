package tplr

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"

	"github.com/mantidtech/tplr/functions"
)

// Tplr manages loading and rendering templates
type Tplr struct {
	name     string
	Template *template.Template
}

// New creates a new tplr instance
func New(name string) *Tplr {
	return &Tplr{
		name: name,
	}
}

// Load a template from the supplied Reader and create a new Template object
func (t *Tplr) Load(r io.Reader) error {
	tSet := template.New(t.name)
	tSet.Funcs(functions.All(tSet))

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read template: %w", err)
	}

	tpl, err := tSet.Parse(string(b))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	t.Template = tpl
	return nil
}

// Generate text from the template and data supplied and writes it to the given Writer
func (t *Tplr) Generate(w io.Writer, vars map[string]interface{}) error {
	var err error
	var f bytes.Buffer
	err = t.Template.ExecuteTemplate(&f, t.name, vars)
	if err != nil {
		return fmt.Errorf("failed to apply template: %w", err)
	}

	_, err = w.Write(f.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return err
}
