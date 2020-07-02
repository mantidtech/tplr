![Go](https://github.com/mantidtech/tplr/workflows/Go/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/mantidtech/tplr/badge.svg?branch=master)](https://coveralls.io/github/mantidtech/tplr?branch=master)
![License](https://img.shields.io/github/license/mantidtech/tplr)

# `tplr` (templater)

A tool to create files rendered from go templates and json

```
Usage: tplr [-o <output file>] [-d <data file>] [-t <template file>] [inline template]
Usage: tplr [-h|-v]

Where:
  -o <output file>   is a file to write to (default: stdout)
  -d <data file>     is a json file containing the templated variables (default: stdin)
  -f <template file> is a file using the go templating notation.
     If this is not specified, the template is taken from the remaining program args

Information:
  -h Prints this messge
  -v Prints the program version number and exits
```

The templating markup language used is defined by the go [text/template](https://pkg.go.dev/text/template) package

---
## Install

Using go
```
go get -u github.com/mantidtech/tplr/cmd/tplr
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

In addition to the [standard template functions](https://pkg.go.dev/text/template?tab=doc#hdr-Functions), `tplr` introduces:

### Composition Actions

#### `{{ include "name" pipeline }}`

 Works like `template`, but allows usage within pipelines
 eg
``` 
{{- include "foo" . | indent 6 -}}
```

### Logical Operators

#### `{{ isZero pipeline }}`

Returns `true` if the pipeline is empty (ie the `zero` value of its type) OR 
if it's a pointer and the dereferenced value is zero, OR 
if the type of the pipeline has a `length` (eg array, slice, map, string), and the length is zero

#### `{{ whenEmpty pipeline1 pipeline0 }}`

If `pipeline0` is empty (ie the `zero` value of its type), returns `pipeline1` otherwise `pipeline0`

slightly backwards to look at, makes sense when chained though:
```
{{ .SomeValue | whenEmpty "something else" }}
```

---
### Formatting

#### `{{ bracketWith S pipeline }}`

Surrounds the pipeline with bracket pairs taken from `S`

`S` must be an even-length string.  The first half is used as the opening bracket, and the second as closing.

eg 
```
{{ bracketWith "<>" "html" }}
```
produces
```
<html>
```

#### `{{ bracket pipeline }}`

Surrounds the `pipeline` with `(` & `)`

Equivalent to `{{ bracketWith "()" pipeline }}`.


#### `{{ prefix X N S }}`

Prefixes all lines in the string `S` with `N` times the prefix string `X`

#### `{{ suffix X N S }}`

Appends all lines in the string `S` with `N` times the suffix string `X`

#### `{{ indent N S }}`

Indents all lines in the string `S` with `N` tabs.

Equivalent to `{{ prefix "\t" N S }}`.

eg
```
{{- indent 2 "Hello\n\tWorld" -}}
```
produces
```
        Hello
            World
```

#### `{{ joinWith S ARG_1..ARG_N }}`

Concatenates the args into string joined by the string `S`

#### `{{ join ARG_1..ARG_N }}`

Concatenates the args into a single string

Equivalent to `{{ joinWith "" ARG }}`.

#### `{{ splitOn S ARG }}`

Splits the arg into a list of strings on occurrences of the string `S`

#### `{{ padLeft N ARG }}`

Returns the arg with a width of at least `N` characters, adding spaces at the front to pad out the string

#### `{{ padRight N ARG }}`

Returns the arg with a width of at least `N` characters, adding spaces at the end to pad out the string

#### `{{ uppercaseFirst ARG }}`

Returns the arg with the first character capitalised. There is no effect to other characters.

#### `{{ toLower ARG }}`

Returns the arg with the all characters converted to lowercase.

#### `{{ toUpper ARG }}`

Returns the arg with the all characters converted to uppercase.

#### `{{ rep N ARG }}`

Returns the arg printed `N` number of times.

#### `{{ space [N] }}`

Returns a space character.  
If `N` is supplied, returns the given number of spaces.
Equivalent to `{{ rep N " " }}`.

#### `{{ tab [N] }}`

Returns a tab character.  
If `N` is supplied, return the given number of tabs.
Equivalent to `{{ rep N "\t" }}`.

#### `{{ nl [N] }}`

Returns a newline character.  
If `N` is supplied, return the given number of newlines.
Equivalent to `{{ rep N "\n" }}`.

#### `{{ typeName ARG }}`

Returns the name of the Go type for the underlying variable of the argument

---
### Encoding

#### `{{ toJSON ARG }}`

Converts the given arg to a JSON string

#### `{{ formatJSON S ARG }}`

Pretty-prints the pipeline.  Returns an error if the arg isn't valid JSON.
Each element in the JSON object or array begins on a new line, indented by one or more copies of `S` according to the nesting depth of the object.

#### `{{ toYAML ARG }}`

Converts the given arg to a YAML string

#### `{{ toBase64 ARG }}`

Converts the given arg to encoded base64

#### `{{ fromBase64 ARG }}`

Decode the given base64 `ARG`

---
### Lists

These methods work when the value of the pipeline is an array or slice (and return an error otherwise)

#### `{{ list ITEM_1.ITEM_N }}`

Creates a new list with the given set of items

#### `{{ first LIST }}`

Returns the first item in the list

#### `{{ last LIST }}`

Returns the last item of a list

#### `{{ rest LIST }}` / `{{ pop LIST }}`

Returns a list of everything except for the first item (aliased as pop)

#### `{{ push LIST ITEM }}`

Adds an item to the end of a list, returning the new list

#### `{{ unshift LIST ITEM }}`

Adds an item to the beginning of a list, returning the new list

#### `{{ slice I J LIST }}`

Returns a slice of the list, ie all items between indexes I (inclusive) and J (exclusive)

#### `{{ contains LIST ITEM }}`

Returns `true` if the item are present in the list

#### `{{ filter LIST ITEM }}`

Returns a list with all instances of all item removed from it


---
### Functions as a Library

You can also import the template functions to use with your own templating code.

    go get github.com/mantidtech/tplr/functions
    
and use:

    import github.com/mantidtech/tplr/functions
    

A convenience function, `functions.All(t *template.Template)`, is provided to return all the functions as a `template.FuncMap`

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
