package gorgeous

import "fmt"

var classes = map[string]CSS{}

func Class(name string, class CSSProps) {
	RegisterClassString(name, renderCSSProps(name, class))
}
func RegisterClassString(name string, class CSS) {
	if classes[name] != "" {
		panic(fmt.Sprintf(`gorgeous: class '%s' is already registered`, name))
	}
	classes[name] = class
}

func collectClasses() CSS {
	var collected CSS

	for _, class := range classes {
		collected += class + "\n"
	}

	return collected
}
