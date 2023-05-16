package gorgeous

import "fmt"

var classes = map[string]CSS{}

// Create a CSS class from a map of CSS properties and
// adds it to the global css.
// Eg:
//	gorgeous.Class("my-class", gorgeous.CSSProps{
//		"color": "red",
//	})
//
// renders as:
//
//	`.my-class {
//		color: red;
//	}`
func Class(name string, class CSSProps) {
	RawClass(name, renderCSSProps(name, class))
}

// Create a CSS class from a string and adds it to the global css.
// Eg:
//	gorgeous.RawClass("my-class", `.my-class {
//		color: red;
//	}`)
//
// renders as:
//
//	`.my-class {
//		color: red;
//	}`
func RawClass(name string, class CSS) {
	if classes[name] != "" {
		panic(fmt.Sprintf(`gorgeous: class '%s' is already registered`, name))
	}
	classes[name] = class
}

// Collects all registered CSS classes into a single CSS string.
func collectClasses() CSS {
	var collected CSS

	for _, class := range classes {
		collected += class + "\n"
	}

	return collected
}
