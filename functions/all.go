// Package functions collects all the functions defined for the various template methods together into a single package
package functions

import (
	"text/template"

	"github.com/mantidtech/tplr/functions/console"
	"github.com/mantidtech/tplr/functions/datetime"
	"github.com/mantidtech/tplr/functions/encoding"
	"github.com/mantidtech/tplr/functions/helper"
	"github.com/mantidtech/tplr/functions/list"
	"github.com/mantidtech/tplr/functions/logic"
	"github.com/mantidtech/tplr/functions/math"
	"github.com/mantidtech/tplr/functions/strings"
	"github.com/mantidtech/tplr/functions/templates"
)

// All returns all the templating functions
func All(t *template.Template) template.FuncMap {
	return CombineFunctionLists(
		strings.Functions(),
		list.Functions(),
		logic.Functions(),
		math.Functions(),
		datetime.Functions(),
		encoding.Functions(),
		console.Functions(),
		templates.Functions(t),
	)
}

// CombineFunctionLists together from zero more supplied lists
func CombineFunctionLists(fnList ...template.FuncMap) template.FuncMap {
	return helper.Combine(fnList...)
}
