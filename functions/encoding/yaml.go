package encoding

import "gopkg.in/yaml.v2"

// ToYAML returns the given value as a yaml string
func ToYAML(val interface{}) (string, error) {
	b, err := yaml.Marshal(val)
	return string(b), err
}
