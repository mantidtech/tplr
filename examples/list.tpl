{{.}}

{{ list "a" 1 false}}
{{ range list 1 2 3 }}{{bracket .}}{{ end }}


{{ first .a }}
{{ rest .a }}
{{ last .a }}
{{ pop .a }}

{{ contains .a "x" }}
{{ contains .a "d" }}
{{ contains .a 3 }}

{{ push .a "x" }}
{{ unshift .a "x" }}
{{ filter .a "c" }}
{{ filter .b "o" }}

{{ slice 1 3 .a }}
