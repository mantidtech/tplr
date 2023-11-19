# Wordcase

The wordcase package provides methods for converting words to different "casing" forms, eg:
* SnakeCase
* KebabCase
* DotCase
* ScreamingSnakeCase
* CamelCase
* PascalCase
* Words
* TitleCase
 
When those methods aren't enough, it also offers a way to build a function that behave exactly* how you want.

eg, to create a title casing method you could do something like:
```
    titleCase := wordcase.
        NewPipeline().
        TokenizeUsing(SimpleCategorizer, unicode.IsSpace, true). // tokenise on whitespace, don't consider the separator as part of the token
        WithAllFormatter(strings.ToLower).                       // convert all tokens to lowercase
        WithFormatter(strings.ToUpper, wordcase.KeyWords).       // convert keywords to uppercase (eg HTML, XML, ID -- can supply your own list)
        WithAllFormatter(wordcase.UppercaseFirst).               // convert the first character of each token to uppercase
        JoinWith(" ")                                            // join the tokens back together using a space as a separator

    // use it a la:
    fmt.Println(titleCase("oNe two Html THREE")) // will print: One Two HTML Three
```


This package is pretty much overkill, but it does handle a bunch of edge cases pretty nicely 
(that's hard to do with straight regex substitution).

\* exactly-ish 

---
## Functions

The package contains some example pipelines for common conversions.  
If the methods don't behave exactly how you'd like, you can copy the pipelines 
from the standalone.go file and modify them to suit.

### SnakeCase

eg

`SnakeCase("One example id")` -> `"one.example.id"`

### KebabCase

eg

`KebabCase("One example id")` -> `"one-example-id"`

### DotCase

eg

`DotCase("One example id")' -> `"one.example.id"`

### ScreamingSnakeCase

eg

`ScreamingSnakeCase("One example id")' -> `"ONE_EXAMPLE_ID"`

### CamelCase

eg

`CamelCase("One example id")` -> `"oneExampleID"`

### PascalCase

eg

`PascalCase("One example id")` -> `"OneExampleID"`

### Words

Converts an "otherly-cased" string into words separated by spaces (actual case of each word left untouched)

eg

`Words("oneExampleID")` -> `"one Example ID"`

### TitleCase

Converts each word in a string to uppercase.

Note: there's a `Title` function in [strings](https://pkg.go.dev/strings?tab=doc#Title) which may be what you want if you don't want 
non-alphanumeric characters stripped, or your string is already broken into space-separated words

eg

`TitleCase("ONE_EXAMPLE_ID")` -> `"One Example ID"`


---
## Pipelines

Perform transformations using a pipeline of processors.

A `Pipeline` type has the signature `func(string) Tokens`.

For a function to work as a pipeline stage, it must have a `Pipeline` compatible receiver and return a `Pipeline` 

Pipeline are built using the following kinds of stages:
* Tokenization
* Formatting
* Combination


### `Tokenization` 
Splits a string into tokens. 

Performed with the method:
```
TokenizeUsing(test SeparatorTest, sep IsRuneSeparator, del bool)
```

Where:
* a SeparatorTest holds the logic on deciding if we're going to split on a given rune
* a IsRuneSeparator determines if we consider a give rune a separator
* set del to true to delete the separating rune, or false to make it part of the next token


Multiple tokenization steps can be specified, creating a multi-pass parser.  
This can simplify gnarly situations with complicated rules on what makes a token. 

#### SeparatorTest

Has the signature: `func(word string, idx int, test IsRuneSeparator) bool`

A SeparatorTest determines if the given index in the given string is a token separator.
By having the whole parse text as reference, smarter tests can be performed using lookahead/lookbehind along with the IsRuneSeparator test.

There's two SeparatorTests provides currently:

##### LookAroundCategorizer
 
LookAroundCategorizer considers a rune to be a separator if it passes IsRuneSeparator(c), 
but also so long as bot of the preceding and succeeding runes are NOT separators

##### SimpleCategorizer

SimpleCategorizer considers only the current rune when deciding on separators.


### Formatting

Formatters are applied to tokens, usually to change the case in some way, but any string operation can be applied. 

Performed with the method:
```
WithFormatter(formatter Formatter, selector TokenSelector)
```

There is also a convenience method:
```
WithAllFormatter(formatter Formatter)
```
which is simple shorthand for:
```
WithFormatter(formatter, ToAll)
```


Where:
* Formatter is a function that takes a string, does stuff, returns a string
* TokenSelector is a function that specifies which tokens to apply formatting

#### Formatters

A formatter has the function signature `func(string) string`, so many of the operations in the `strings` library can be applied (eg `ToLower`, `ToUpper`).

One additional method is currently supplied: `UppercaseFirst` which converts the first character of a string to uppercase, and leaves the rest untouched

#### TokenSelectors

A `TokenSelector` determines which tokens to apply formatting.  
They have the signature `func(t Tokens) []int`, where `t` is the current list of tokens, and `[]int` is a slice of indexes on which the formatter should operate.

Provided Selectors include:

* `ToFirst` - returns an index to the first token
* `ToLast` - returns an index to the last token
* `ToRest` - returns an index to all tokens except the first
* `ToAll` - returns an index to all tokens

There are 3 logic selectors that operate on other selectors:
* `Not(sFn TokenSelector)` - inverts the selection list of the provided selector
* `And(a, b TokenSelector)` - returns a selector that matches where both the selectors `a` and `b` match
* `Or(a, b TokenSelector)` - returns a selector that matches where either of the selectors `a` or `b` match

There is a `TokenSelector` generator method `KeyWordFn(keywords []string)` which creates a selector using a list of words provided.

Two lists are provided currently:
* `GoLintKeywords` - a list of words that [golint](https://github.com/golang/lint/blob/master/lint.go#L740) uses as a set of common initialisms.
* `UsefulKeyWords` - a list that includes all the words from `GoLintKeywords` plus a few more initialisms such as `grpc`, `yaml` and `toml`


### Combination

A combination stage is for joining the tokens into new words again.

Performed with the stage
```
JoinWith(sep string)
```

Currently, only one `Combiner` is supported and must be the last stage of a pipeline.

This method simply joins the tokens together using the string `sep` as glue
