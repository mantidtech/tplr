package helper

// TestSet is used for specifying a template function test instance
type TestSet struct {
	Name     string
	Template string
	Args     TestArgs
	Want     string
	WantErr  bool
}

// TestArgs are the arguments used when testing a template
type TestArgs map[string]interface{}
