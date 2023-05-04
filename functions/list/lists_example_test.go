package list

import (
	"os"
	"text/template"
)

var a = map[string]any{
	"a": []string{"a", "b", "c", "d", "e"},
	"b": []string{"f", "o", "o", "b", "a", "r"},
}

func helperApplyAndRenderTemplate(tpl string, data map[string]any) {
	t := template.New("example")
	t.Funcs(Functions())
	t, _ = t.Parse(tpl)
	_ = t.ExecuteTemplate(os.Stdout, "example", data)
}

var list = `
{{ print "{{list \"a\" \"b\" \"c\"}}" }} = {{list "a" "b" "c"}}
{{ print "{{list \"a\" 5 false}}" }} = {{list "a" 5 false}}
`

func ExampleList() {
	helperApplyAndRenderTemplate(list, a)
	// Output:
	// {{list "a" "b" "c"}} = [a b c]
	// {{list "a" 5 false}} = [a 5 false]
}

var first = `
{{ print "{{.a}}" }}       = {{.a}}
{{ print "{{first .a}}" }} = {{first .a}}
`

func ExampleFirst() {
	helperApplyAndRenderTemplate(first, a)
	// Output:
	// {{.a}}       = [a b c d e]
	// {{first .a}} = a
}

var last = `
{{ print "{{.a}}" }}      = {{.a}}
{{ print "{{last .a}}" }} = {{last .a}}
`

func ExampleLast() {
	helperApplyAndRenderTemplate(last, a)
	// Output:
	// {{.a}}      = [a b c d e]
	// {{last .a}} = e
}

var rest = `
{{ print "{{.a}}" }}      = {{.a}}
{{ print "{{rest .a}}" }} = {{rest .a}}
`

func ExampleRest() {
	helperApplyAndRenderTemplate(rest, a)
	// Output:
	// {{.a}}      = [a b c d e]
	// {{rest .a}} = [b c d e]
}

var contains = `
{{ print "{{.a}}" }}              = {{.a}}
{{ print "{{contains .a \"x\"}}" }} = {{contains .a "x"}}
{{ print "{{contains .a \"d\"}}" }} = {{contains .a "d"}}
{{ print "{{contains .a 3}}" }}   = {{contains .a 3}}
`

func ExampleContains() {
	helperApplyAndRenderTemplate(contains, a)
	// Output:
	// {{.a}}              = [a b c d e]
	// {{contains .a "x"}} = false
	// {{contains .a "d"}} = true
	// {{contains .a 3}}   = false

}

var push = `
{{ print "{{.a}}" }}          = {{.a}}
{{ print "{{push .a \"x\"}}" }} = {{push .a "x"}}
`

func ExamplePush() {
	helperApplyAndRenderTemplate(push, a)
	// Output:
	// {{.a}}          = [a b c d e]
	// {{push .a "x"}} = [a b c d e x]
}

var unshift = `
{{ print "{{.a}}" }}             = {{.a}}
{{ print "{{unshift .a \"x\"}}" }} = {{unshift .a "x"}}
`

func ExampleUnshift() {
	helperApplyAndRenderTemplate(unshift, a)
	// Output:
	// {{.a}}             = [a b c d e]
	// {{unshift .a "x"}} = [x a b c d e]
}

var filter = `
{{ print "{{.a}}" }}            = {{.a}}
{{ print "{{filter .a \"x\"}}" }} = {{filter .a "x"}}
{{ print "{{filter .a 3}}" }}   = {{filter .a 3}}
{{ print "{{filter .a \"c\"}}" }} = {{filter .a "c"}}

{{ print "{{.b}}" }}            = {{.b}}
{{ print "{{filter .b \"o\"}}" }} = {{filter .b "o"}}
`

func ExampleFilter() {
	helperApplyAndRenderTemplate(filter, a)
	// Output:
	// {{.a}}            = [a b c d e]
	// {{filter .a "x"}} = [a b c d e]
	// {{filter .a 3}}   = [a b c d e]
	// {{filter .a "c"}} = [a b d e]
	//
	// {{.b}}            = [f o o b a r]
	// {{filter .b "o"}} = [f b a r]
}

var pop = `
{{ print "{{.a}}" }}     = {{.a}}
{{ print "{{pop .a}}" }} = {{pop .a}}
`

func ExamplePop() {
	helperApplyAndRenderTemplate(pop, a)
	// Output:
	// {{.a}}     = [a b c d e]
	// {{pop .a}} = [a b c d]
}
