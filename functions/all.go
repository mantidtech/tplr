package functions

import "text/template"

// All returns all the templating functions
func All(t *template.Template) template.FuncMap {
	return CombineFunctionLists(
		EncodingFunctions(),
		ListFunctions(),
		LogicFunctions(),
		MiscellaneousFunctions(),
		StringFunctions(),
		TemplateFunctions(t),
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
