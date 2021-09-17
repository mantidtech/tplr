j![Go](https://github.com/mantidtech/tplr/workflows/Go/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/mantidtech/tplr/badge.svg?branch=master)](https://coveralls.io/github/mantidtech/tplr?branch=master)
![License](https://img.shields.io/github/license/mantidtech/tplr)

# `tplr` (templater)

A tool to create files rendered from go templates and json

```
Usage: tplr [-f] [-o <output file>] [-d <data file>] [-t <template file>] [inline template]
Usage: tplr [-h|-v]

Where:
  -o <output file>   is a file to write to (default: stdout)
  -d <data file>     is a json file containing the templated variables (default: stdin)
  -t <template file> is a file using the go templating notation.
     If this is not specified, the template is taken from the remaining program args

Options:
  -f If the destination file already exits, overwrite it.  (default is to do nothing)

Information:
  -h Prints this messge
  -v Prints the program version number and exits

```

The templating markup language used is defined by the go [text/template](https://pkg.go.dev/text/template) package

---
## Install

Using go
```
go install github.com/mantidtech/tplr/cmd/tplr@latest
```

---
## Examples
```bash
    echo '{"to":"World"}' | tplr 'Hello {{.to}}!'
    # displays:
    Hello World!
``` 
```bash
    cat << EOF > sample.tpl
    <html>
        <head>
            <title>{{ .title }}</title>
        <head>
        <body>
            <ul>
            {{- range .items -}}
                {{- include "listItem" . | indent 3 -}}
            {{- end }}
            </ul>
        </body>
    </html>
    
    {{- define "listItem" }}
    <li>
        {{.}}
    </li>
    {{- end -}}
    EOF

    cat << EOF > sample.json
    {
        "title": "My List Of Doom",
        "items": [
            "Acquire lair",
            "Acquire henchmen",
            "Acquire doomsday weapon"
        ]
    }
    EOF

    tplr -d sample.json -t sample.tpl -o sample.html

    cat sample.html
    # displays:
    <html>
        <head>
            <title>My List Of Doom</title>
        <head>
        <body>
            <ul>
                <li>
                    Acquire lair
                </li>
                <li>
                    Acquire henchmen
                </li>
                <li>
                    Acquire doomsday weapon
                </li>
            </ul>
        </body>
    </html>
``` 
NOTE: `tplr` uses the [text/template](https://pkg.go.dev/text/template) and not the [html/template](https://pkg.go.dev/text/template) package.  
This means there's no special translation of html elements, and therefore output isn't protected against code injection 
-- so trust your data source if you intend to use `tplr` in any sort of public facing html generation process.


---
## Template Functions

In addition to the [standard template functions](https://pkg.go.dev/text/template?tab=doc#hdr-Functions), 
`tplr` introduces several collections of functions

* `strings` - Operations on strings (and pipelines convertable to strings)
* `list` - Operations to operate on lists (arrays) of values
* `logic` - Logical operations
* `encoding` and `decoding` - For marshalling and unmarshalling data structures
* `templates` - Meta-functions for template  processing 
* `console` - Operations specific to processing templates to a terminal


---
### String Operations

* #### `{{ toColumns NUMBER TEXT }}`

ToColumns formats the given `TEXT` to not take more than `NUMBER` characters per line,
splitting on the space before the word that would take the line over

Note: Embedded newlines have no special treatment, so text containing them could look wonky.
Either strip them first, or break the input into multiple strings and process individually

eg
```
{{ toColumns 5 "a b c d e f g" }}
```
produces
```
a b c
d e f
g
```

* #### `{{ bracketWith BRACKETPAIR PIPELINE }}`

Surrounds the pipeline with bracket pairs taken from the string `BRACKETPAIR`

`BRACKETPAIR` must be an even-length string.  The first half is used as the opening bracket and the second as closing.

eg
```
{{ bracketWith "<>" "html" }}
```
produces
```
<html>
```

* #### `{{ bracket PIPELINE }}`

Surrounds the `PIPELINE` with `(` & `)`.

Equivalent to `{{ bracketWith "()" PIPELINE }}`.

* #### `{{ prefix PREFIX COUNT STRING }}`

Prefixes all lines in the string `STRING` with `COUNT` times the prefix string `PREFIX`.

eg:
```gotemplate
{{ prefix "> " 2 "Foo" }}
```
produces:
```
> > Foo
```

* #### `{{ suffix SUFFIX COUNT STRING }}`

Appends all lines in the string `STRING` with `COUNT` times the suffix string `SUFFIX`.

* #### `{{ indent COUNT STRING }}`

Indents all lines in the string `STRING` with `COUNT` tabs.

Equivalent to `{{ prefix "\t" COUNT STRING }}`.

eg
```
{{- indent 2 "Hello\n\tWorld" -}}
```
produces
```
        Hello
            World
```

* #### `{{ splitOn GLUE CONTENT }}`

Splits `CONTENT` into a list of strings on occurrences of the string `GLUE`.

* #### `{{ padLeft COUNT PIPELINE }}`

Returns `PIPELINE` with a width of at least `COUNT` characters, adding spaces at the front to pad out the string.

* #### `{{ padRight COUNT PIPELINE }}`

Returns `PIPELINE` with a width of at least `COUNT` characters, adding spaces at the end to pad out the string.

* #### `{{ uppercaseFirst ARG }}`

Returns the arg with the first character capitalised. There is no effect to other characters.

* #### `{{ toLower STRING }}`

Returns `STRING` with the all characters converted to lowercase.

* #### `{{ toUpper STRING }}`

Returns `STRING` with the all characters converted to uppercase.

* #### `{{ rep COUNT ARG }}`

Returns the `ARG` printed `COUNT` number of times.

* #### `{{ space [COUNT] }}` / `{{ sp [COUNT] }}`

Returns a space character.  
If `COUNT` is supplied, returns the given number of spaces.
Equivalent to `{{ rep COUNT " " }}`.

* #### `{{ tab [COUNT] }}`

Returns a tab character.  
If `COUNT` is supplied, return the given number of tabs.
Equivalent to `{{ rep COUNT "\t" }}`.

* #### `{{ nl [COUNT] }}`

Returns a newline character.  
If `COUNT` is supplied, return the given number of newlines.
Equivalent to `{{ rep COUNT "\n" }}`.

* #### `{{ typeName PIPELINE }}`

Returns the name of the Go type for the underlying variable of the argument `PIPELINE`.

* #### `{{ camelCase STRING }}`

Convert the given `STRING` to camelCase.

eg:
```gotemplate
{{- camelCase "foo bar" -}}
```
produces:
```
fooBar
```

* #### `{{ dotCase STRING }}`

Convert the given `STRING` to dotCase.

eg:
```gotemplate
{{- dotCase "foo bar" -}}
```
produces:
```
foo.bar
```

* #### `{{ kebabCase STRING }}`

Convert the given `STRING` to kebabCase.

eg:
```gotemplate
{{- kebabCase "foo bar" -}}
```
produces:
```
foo-bar
```

* #### `{{ pascalCase STRING }}`

Convert the given `STRING` to pascalCase.

eg:
```gotemplate
{{- pascalCase "foo bar" -}}
```
produces:
```
FooBar
```

* #### `{{ screamingSnakeCase STRING }}`

Convert the given `STRING` to screamingSnakeCase.

eg:
```gotemplate
{{- screamingSnakeCase "foo bar" -}}
```
produces:
```
FOO_BAR
```

* #### `{{ snakeCase STRING }}`

Convert the given `STRING` to snakeCase.

eg:
```gotemplate
{{- snakeCase "foo bar" -}}
```
produces:
```
foo_bar
```

* #### `{{ titleCase STRING }}`

Convert the given `STRING` to titleCase.

eg:
```gotemplate
{{- titleCase "foo html" -}}
```
produces:
```
Foo Html
```

* #### `{{ titleCaseWithAbbr STRING }}`

Convert the given `STRING` to titleCaseWithAbbr.

eg:
```gotemplate
{{- titleCase "foo html" -}}
```
produces:
```
Foo HTML
```

* #### `{{ toWords STRING }}`

Convert the given `STRING` to toWords.

eg:
```gotemplate
{{- toWords "foo_bar_baz" -}}
```
produces:
```
foo bar baz
```

* #### `{{ q STRING }}`

Quotes the given `STRING` with single quotes.

eg:
```gotemplate
{{- q "foo" -}}
```
produces:
```
'foo'
```

* #### `{{ qq STRING }}`

Quotes the given `STRING` with double quotes.

eg:
```gotemplate
{{- qq "foo" -}}
```
produces:
```
"foo"
```

* #### `{{ qb STRING }}`

Quotes the given `STRING` with back-quotes.

eg:
```gotemplate
{{- qb "foo" -}}
```
produces:
```
`foo`
```

* #### `{{ trim STRING }}`

Trims whitespace from the start and end of `STRING`.

eg:
```gotemplate
X{{- trim "   foo   " -}}X
```
produces:
```
XfooX
```

* #### `{{ typeKind PIPELINE }}`

Returns the 'kind' of the `PIPELINE`.

eg:
```gotemplate
{{ typeKind 6 }}
{{ typeKind "foo" }}
{{ list "A" "B" "C" | typeKind }}
```
produces:
```
int
string
slice
```

* #### `{{ unindent COUNT STRING }}`

Removes up to `COUNT` spaces from the start of every line in `STRING`.

* #### `{{ now [FORMAT] }}`

Returns the current time as a string in the given `FORMAT` or `time.RFC3339` if none is specified.

Time formats are specified using standard go formatting as defined at https://pkg.go.dev/time#pkg-constants

eg:
```gotemplate
{{ now }}
{{ now "Mon Jan _2" }}
{{ now "15:04:05 MST" }}
```
produces:
```
2021-09-16T21:41:58Z10:00
Thu Sep 16
21:41:58 AEST
```

---
### List Functions

These methods work when the value of the pipeline is an array or slice (and return an error otherwise)

* #### `{{ list ITEM_1..ITEM_N }}`

Creates a new list with the given set of items.

* #### `{{ first LIST }}` / `{{ shift LIST }}`

Returns the first item in the list.

* #### `{{ last LIST }}` /  `{{ pop LIST }}`

Returns the last item of a list.

* #### `{{ rest LIST }}`

Returns a list of everything except for the first item.

* #### `{{ push LIST ITEM }}`

Adds an item to the end of a list, returning the new list.

* #### `{{ unshift LIST ITEM }}`

Adds an item to the start of a list, returning the list.

* #### `{{ slice I J LIST }}`

Returns a slice of the `LIST`, ie all items between indexes `I` (inclusive) and `J` (exclusive).

* #### `{{ contains LIST ITEM }}`

Returns `true` if the item are present in the list.

* #### `{{ filter LIST ITEM }}`

Returns a list with all instances of all item removed from it.

* #### `{{ joinWith GLUE ARG_1..ARG_N }}`

Concatenates the args into string joined by the string `GLUE`.

* #### `{{ join ARG_1..ARG_N }}`

Concatenates the args into a single string.

Equivalent to `{{ joinWith "" ARG }}`.


---
### Logical Functions

* #### `{{ isZero PIPELINE }}`

Returns `true` if the pipeline is empty (ie the `zero` value of its type) OR
if it's a pointer and the dereferenced value is zero, OR
if the type of the pipeline has a `length` (eg array, slice, map, string), and the length is zero

* #### `{{ and ARG_1...ARG_n }}`

Returns `ARG_n` if all the arguments evaluate to non-zero or `""` otherwise.

* #### `{{ or ARG_1...ARG_n }}`

Returns the first argument that evaluates to non-zero (from left to right) or `""` if none do.

* #### `{{ whenEmpty VALUE COND }}`

Returns `VALUE` if `COND` is a zero value, otherwise it returns `COND`.

eg:
```gotemplate
{{ "thing" | whenEmpty "something else" }}
{{ "" | whenEmpty "something else" }}
```
Produces:
```
thing
something else
```

* #### `{{ when VALUE COND }}`

Returns `VALUE` if `COND` is not zero, otherwise it returns an empty string.

eg:
```gotemplate
X{{ "thing" | when "something else" }}X
X{{ "" | when "something else" }}X
```
Produces:
```
Xsomething elseX
XX
```


---
### Encoding and Decoding

* #### `{{ toBase64 ARG }}`

Converts the given arg to base64 encoding

* #### `{{ fromBase64 ARG }}`

Decode the given base64 `ARG`

* #### `{{ toJSON ARG }}`

Converts the given arg to a JSON string

* #### `{{ formatJSON INDENT ARG }}`

Pretty-prints the pipeline.  Returns an error if the arg isn't valid JSON.
Each element in the JSON object or array begins on a new line, 
indented by one or more copies of `INDENT` according to the nesting depth of the object.

* #### `{{ toYAML PIPELINE }}`

Converts the given `PIPELINE` to a YAML string


---
### Operations with Templates

* #### `{{ include "NAME" PIPELINE }}`

Works like `template`, but allows usage within pipelines.
eg:
``` 
{{- include "foo" . | indent 6 -}}
```


---
### Console (Terminal) Functions

* #### `{{ terminalWidth }}`

Returns the width of the terminal if available or zero otherwise.


---
### Functions as a Library

You can also import the template functions to use with your own templating code.

    go get github.com/mantidtech/tplr/functions
    
and use:

    import github.com/mantidtech/tplr/functions
    

A convenience function, `functions.All(t *template.Template)`, 
is provided to return all the functions as a `template.FuncMap`

docs are at https://pkg.go.dev/github.com/mantidtech/tplr/functions

---
## To Do

* Process multiple templates and/or datasets at once

* Greatly expand functions, including for
    * Time
    * Anything else people might find useful

* Less basic error messages/presentation

---
(c) 2020 - Julian Peterson `<code@mantid.org>` - [MIT Licensed](LICENSE.md)
