![Go](https://github.com/mantidtech/tplr/workflows/Go/badge.svg)
![License](https://img.shields.io/github/license/mantidtech/tplr)

# `tplr` (`t`em`pl`ate`r`)

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

#### `{{ indent N pipeline }}`

Indents all lines in the `pipeline` with `N` tabs

eg
```
{{- indent 2 "Hello\n\tWorld" -}}
```
produces
```
        Hello
            World
```

#### `{{ joinWith S pipeline }}`

Concatenates the `pipeline` into string joined by the srting `S`

#### `{{ join pipeline }}`

Concatenates the `pipeline` into a single string

Equivalent to `{{ joinWith "" pipeline }}`.

#### `{{ splitOn S pipeline }}`

Splits the pipeline into a list of strings on occurrences of the string `S`

#### `{{ spIndent N pipeline }}`

Indents all lines in the `pipeline` with `N` spaces

#### `{{ padLeft N pipeline }}`

Returns the `pipeline` with a width of at least `N` characters, adding spaces at the front to pad out the string

#### `{{ padRight N pipeline }}`

Returns the `pipeline` with a width of at least `N` characters, adding spaces at the end to pad out the string

#### `{{ uppercaseFirst pipeline }}`

Returns the `pipeline` with the first character capitalised. There is no effect to other characters.

#### `{{ toLower pipeline }}`

Returns the `pipeline` with the all characters converted to lowercase.

#### `{{ toUpper pipeline }}`

Returns the `pipeline` with the all characters converted to uppercase.

#### `{{ rep N pipeline }}`

Returns the `pipeline` `N` number of times.

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

#### `{{ typeName pipeline }}`

Returns the name of the Go type for the underlying variable of the pipeline

#### `{{ toJSON pipeline }}`

Converts the given pipeline to a JSON string

#### `{{ formatJSON S pipeline }}`

Pretty-prints the pipeline.  Returns an error if the pipeline isn't valid JSON.
Each element in the JSON object or array begins on a new line, indented by one or more copies of `S` according to the nesting depth of the object.

#### `{{ toYAML pipeline }}`

Converts the given pipeline to a YAML string


---
### Lists

These methods work when the value of the pipeline is an array or slice (and return an error otherwise)

#### `{{ first pipeline }}`

returns the first item in the list

#### `{{ rest  pipeline }}`

returns a list of everything but the first item

#### `{{ last  pipeline }}`

returns the last item of a list

---
### Functions as a Library

You can also import the template functions to use with your own templating code.

    go get github.com/mantidtech/tplr/functions
    
and use:

    import github.com/mantidtech/tplr/functions
    

A convenience function, `functions.All(t *template.Template)`, is provided to return all the functions as a `template.FuncMap`

---
## To Do

* Read data in other formats (yaml, toml, kv pairs, etc)

* Process multiple templates and/or datasets at once

* Greatly expand functions, including for
    * Time
    * Wordcase library (which I still need to release)
    * Data format translations, eg toJSON, toYAML, base64{enc,dec}
    * Anything else people might find useful

* Less basic error messages/presentation

---
(c) 2020 - Julian Peterson `<code@mantid.org>` - [MIT Licensed](LICENSE.md)
