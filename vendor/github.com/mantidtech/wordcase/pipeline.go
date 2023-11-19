package wordcase

// Pipeline defines operations to apply to a string to tokenise it
type Pipeline func(string) Tokens

// NewPipeline creates a new empty pipeline
func NewPipeline() Pipeline {
	return func(s string) Tokens {
		return Tokens{s}
	}
}

// TokenizeUsing uses the given functions to create tokens.
//
//	SeparatorTest holds the logic on deciding if we're going to split on a given rune
//	IsRuneSeparator determines if we consider a give rune a separator
//	set del to true to delete the separating rune, or false to make it part of the next token
func (f Pipeline) TokenizeUsing(test SeparatorTest, sep IsRuneSeparator, del bool) Pipeline {
	return func(s string) Tokens {
		t := f(s)
		t = t.Tokenize(test, sep, del)
		return t
	}
}

// WithFormatter adds a token formatter.
// The formatter function supplied will be applied to each that the given selector matches
func (f Pipeline) WithFormatter(formatter Formatter, selector TokenSelector) Pipeline {
	return func(s string) Tokens {
		return f(s).Format(formatter, selector)
	}
}

// WithAllFormatter adds a token formatter that applies to every token
func (f Pipeline) WithAllFormatter(formatter Formatter) Pipeline {
	return func(s string) Tokens {
		return f(s).FormatAll(formatter)
	}
}

// JoinWith generates a function that combines tokens together with the given glue
func (f Pipeline) JoinWith(sep string) Combiner {
	return func(s string) string {
		return f(s).Join(sep)
	}
}
