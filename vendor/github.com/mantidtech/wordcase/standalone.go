package wordcase

import (
	"strings"
)

// SnakeCase concatenates tokens into a string separated by underscores
var SnakeCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	JoinWith("_")

// KebabCase concatenates tokens into a string separated by hyphens
var KebabCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	JoinWith("-")

// DotCase concatenates tokens into a string separated by dots (periods)
var DotCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	JoinWith(".")

// ScreamingSnakeCase concatenates tokens into a string separated by an underscore and with every letter converted to uppercase
var ScreamingSnakeCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToUpper).
	JoinWith("_")

// CamelCase creates a string from tokens by making the first rune of each token uppercase (except the first) and concatenating them together
var CamelCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	WithFormatter(UppercaseFirst, ToRest).
	WithFormatter(strings.ToUpper, And(ToRest, LintWords)).
	JoinWith("")

// PascalCase creates a string from tokens by making the first rune of each token uppercase and concatenating them together
var PascalCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	WithAllFormatter(UppercaseFirst).
	WithFormatter(strings.ToUpper, LintWords).
	JoinWith("")

// Words concatenates tokens into a space separated string
var Words = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithFormatter(strings.ToUpper, LintWords).
	JoinWith(" ")

// TitleCase creates a string from tokens by making the first rune of each token uppercase and joining them with spaces
var TitleCase = NewPipeline().
	TokenizeUsing(LookAroundCategorizer, NotLetterOrDigit, true).
	TokenizeUsing(LookAroundCategorizer, NotLowerOrDigit, false).
	WithAllFormatter(strings.ToLower).
	WithFormatter(strings.ToUpper, LintWords).
	WithAllFormatter(UppercaseFirst).
	JoinWith(" ")
