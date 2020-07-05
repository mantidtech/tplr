package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/mantidtech/tplr"
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
		fmt.Printf("%s\n", tplr.Version())
		os.Exit(0)
	} else if *help {
		showHelp()
		os.Exit(0)
	}

	var tpl io.Reader
	if *templateFile != "" {
		tpl, err = tplr.GetFileReader(*templateFile)
	} else {
		tpl, err = tplr.ReadStringsAsFile(s.Args()...)
	}
	if err != nil {
		errorAndExit("Failed to read template file: %v\n", err)
	}

	t, err := tplr.Load(tpl)
	if err != nil {
		errorAndExit("Failed to load template: %v\n", err)
	}

	out, err := tplr.GetFileWriter(*outputFile, *force)
	if err != nil {
		errorAndExit("Failed to open output file: %v\n", err)
	}

	vars, err := tplr.ReadDataFile(*dataFile)
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
	fmt.Printf("Usage: %s [-f] [-o <output file>] [-d <data file>] [-t <template file>] [inline template]\n", app)
	fmt.Printf("Usage: %s [-h|-v]\n", app)
	fmt.Print("\n")
	fmt.Printf("Where:\n")
	fmt.Printf("  -o <output file>   is a file to write to (default: stdout)\n")
	fmt.Printf("  -d <data file>     is a json file containing the templated variables (default: stdin)\n")
	fmt.Printf("  -t <template file> is a file using the go templating notation.\n")
	fmt.Printf("     If this is not specified, the template is taken from the remaining program args\n")
	fmt.Print("\n")
	fmt.Printf("Options:\n")
	fmt.Printf("  -f If the destination file already exits, overwrite it.  (default is to do nothing)\n")
	fmt.Print("\n")
	fmt.Printf("Information:\n")
	fmt.Printf("  -h Prints this messge\n")
	fmt.Printf("  -v Prints the program version number and exits\n")
}

func errorAndExit(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
