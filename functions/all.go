package functions

import (
	"text/template"

	"github.com/mantidtech/tplr/functions/console"
	"github.com/mantidtech/tplr/functions/encoding"
	"github.com/mantidtech/tplr/functions/list"
	"github.com/mantidtech/tplr/functions/logic"
	"github.com/mantidtech/tplr/functions/strings"
	"github.com/mantidtech/tplr/functions/templates"
)

// All returns all the templating functions
func All(t *template.Template) template.FuncMap {
	return CombineFunctionLists(
		console.Functions(),
		encoding.Functions(),
		list.Functions(),
		logic.Functions(),
		strings.Functions(),
		templates.Functions(t),
	)
}

// CombineFunctionLists together from zero more supplied lists
func CombineFunctionLists(fnList ...template.FuncMap) template.FuncMap {
	res := make(template.FuncMap)
	for _, fnl := range fnList {
		for k, fn := range fnl {
			res[k] = fn
		}
	}
	return res
}
