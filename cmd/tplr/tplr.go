package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gitlab.com/mantidtech/tplr"
)

func main() {
	s := flag.NewFlagSet("tplr args", flag.ExitOnError)

	templateFile := s.String("t", "", "Read the template from the file with the given name")
	dataFile := s.String("d", "-", "File to read data from")
	outputFile := s.String("o", "-", "Write the processed template to the named file")
	force := s.Bool("f", false, "Overwrite the destination file if it already exits (otherwise do nothing)")
	help := s.Bool("h", false, "Shows this help message")
	showVersion := s.Bool("v", false, "Display version information")

	err := s.Parse(os.Args[1:])
	if err != nil {
		errorAndExit("Error while parsing flags: %v\n", err)
	}

	if *showVersion {
		fmt.Printf("version: %s\n", tplr.Version())
		os.Exit(0)
	} else if *help {
		showHelp()
		os.Exit(0)
	}

	var tpl io.Reader
	if *templateFile != "" {
		tpl, err = tplr.GetFileReader(*templateFile)
		if err != nil {
			errorAndExit("Failed to open template file: %v\n", err)
		}
	} else {
		var b bytes.Buffer
		b.WriteString(strings.Join(s.Args(), " "))
		tpl = &b

		if b.Len() == 0 {
			errorAndExit("No template supplied\n")
		}
	}

	t, err := tplr.Load(tpl)
	if err != nil {
		errorAndExit("Failed to load template: %v\n", err)
	}

	out, err := tplr.GetFileWriter(*outputFile, *force)
	if err != nil {
		errorAndExit("Failed to open data file: %v\n", err)
	}

	vars, err := getData(*dataFile)
	if err != nil {
		errorAndExit("%v\n", err)
	}

	err = tplr.GenerateFromTemplate(out, t, vars)
	if err != nil {
		errorAndExit("Failed to generate output: %v\n", err)
	}
}

func showHelp() {
	_, app := path.Split(os.Args[0])
	fmt.Printf("%s version %s\n\n", app, tplr.Version())
	fmt.Printf("Usage: %s [-o <output file>] [-d <data file>] [-t <template file>] [inline template]\n", app)
	fmt.Printf("Usage: %s [-h|-v]\n", app)
	fmt.Print("\n")
	fmt.Printf("Where:\n")
	fmt.Printf("  -o <output file>   is a file to write to (default: stdout)\n")
	fmt.Printf("  -d <data file>     is a json file containing the templated variables (default: stdin)\n")
	fmt.Printf("  -f <template file> is a file using the go templating notation.\n")
	fmt.Printf("     If this is not specified, the template is taken from the remaining program args\n")
	fmt.Print("\n")
	fmt.Printf("Information:\n")
	fmt.Printf("  -h Prints this messge\n")
	fmt.Printf("  -v Prints the program version number and exits\n")
}

func getData(filename string) (map[string]interface{}, error) {
	vars := make(map[string]interface{})

	dr, err := tplr.GetFileReader(filename)
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

func errorAndExit(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
