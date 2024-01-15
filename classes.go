package gorgeous

import (
	"fmt"
	"strings"
)

var classes = map[string]CSSClass{}
var media = map[string][]CSS{}

// Create a CSS class from a map of CSS properties and
// adds it to the global css.
// Eg:
//
//	gorgeous.Class(gorgeous.CSSClass{
//		Selector: ".my-class",
//		Props: gorgeous.CSSProps{
//			"color": "red",
//		},
//	})
//
// renders as:
//
//	`.my-class {
//		color: red;
//	}`
func Class(c *CSSClass) *CSSClass {
	classes[c.Selector] = *c

	return c
}

// Create a CSS class within a media query from a map of CSS properties and
// adds it to the global css.
// Eg:
//
//	gorgeous.Media("(max-width: 600px)", ".my-class > div", gorgeous.CSSProps{
//		"color": "red",
//	})
//
// renders as:
//
//	`@media (max-width: 600px) {
//		.my-class > div {
//			color: red;
//		}
//	}`
func Media(query string, selector string, class CSSProps) {
	RawMedia(query, renderCSSProps(selector, class))
}

// Create a CSS class within a media query from a string and adds it to the global css.
// Eg:
//
//	gorgeous.RawMedia("(max-width: 600px)", `.my-class {
//		color: red;
//	}`)
//
// renders as:
//
//	`@media (max-width: 600px) {
//		.my-class {
//			color: red;
//		}
//	}`
func RawMedia(query string, class CSS) {
	if media[query] == nil {
		media[query] = []CSS{}
	}

	media[query] = append(media[query], class)
}

// Create a CSS class from a string and adds it to the global css.
// Eg:
//
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
	if classes[name].Selector != "" {
		panic(fmt.Sprintf(`gorgeous: class '%s' is already registered`, name))
	}
	classes[name] = CSSClass{
		Selector: name,
		Raw:      class,
	}
}

// Collects all registered CSS classes into a single CSS string.
func collectClasses(document HTML) CSS {
	var collected CSS

	for _, class := range classes {
		// If a document is provided, only include classes that are used in the document.
		if !class.Include && document != "" && !strings.Contains(document.String(), extractClassName(class.Selector)) {
			continue
		}

		if class.Raw != "" {
			collected += class.Raw + "\n"
		} else {
			collected += renderCSSProps(class.Selector, class.Props) + "\n"
		}
	}

	for query, classes := range media {
		collected += "@media " + CSS(query) + " {\n"
		for _, class := range classes {
			collected += class + "\n"
		}
		collected += "}\n"
	}

	return collected
}
