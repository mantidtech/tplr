package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// const pkg = "github.com/mantidtech/tplr"
const root = "functions"

func main() {

	dir, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, d := range dir {
		if !d.IsDir() {
			continue
		}

		//if !IsFunctionDir(fmt.Sprintf("%s/%s", root, d.Name())) {
		//	continue
		//}

		fnDir := fmt.Sprintf("%s/%s", root, d.Name())

		fmt.Printf("= %s\n", d.Name())
		err := readFunctions(fnDir)
		if err != nil {
			panic(err)
		}

	}

	//d := processDirMap(dir)
	//for dirName, dirSet := range d {
	//	path := filepath.Join(template.Base, dirSet.Path.Directory())
	//
	//}
}

/*
func IsFunctionDir(dirName string) bool {
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		return false
	}


}
*/

func readFunctions(dirName string) error {
	fs := token.NewFileSet()

	f, err := parser.ParseDir(fs, dirName, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, astPkg := range f {
		for _, astFile := range astPkg.Files {
			ast.Inspect(astFile, collector())
			//fmt.Printf("-> %s %s\n", pName, astFile.Name.Name)
			//spew.Dump(astFile)
		}
	}

	return nil
}

/*
func nodeAsString(fSet *token.FileSet, n ast.Node) (string, error) {
	var b bytes.Buffer
	err := format.Node(&b, fSet, n)
	if err != nil {
		return "", fmt.Errorf("failed to convert node to string: %w", err)
	}
	return b.String(), nil
}
*/

func collector() func(n ast.Node) bool {
	return func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.FuncDecl:
			if strings.HasPrefix(v.Name.String(), "Test") || strings.HasPrefix(v.Name.String(), "Example") {
				return false
			}
			var cmt []string
			if v.Doc != nil && v.Doc.List != nil {
				for _, c := range v.Doc.List {
					cmt = append(cmt, c.Text)
				}
			}
			fmt.Printf("%s %s\n", v.Name.String(), strings.Join(cmt, " "))
		case *ast.Ident:
			// fmt.Printf("%s\n", v.String())
		}
		return true
	}
}
