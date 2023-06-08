package gorgeous

import (
	"fmt"
	"regexp"
)

// Create a CSS class from any number of CSS classes. The properties of the CSS classes
// will be merged into a single CSS class, preferring the properties of the right-most
// CSS class.
//
// ⚠ Note: blend currently only supports classes with a single selector.
//
// Eg:
//
//	gorgeous.Blend(
//		gorgeous.Class(".redtext", gorgeous.CSSProps{
//			"color": "red",
//		}),
//		gorgeous.Class(".bluetext", gorgeous.CSSProps{
//			"color": "blue",
//		}),
//	)
//
// renders as:
//
//	`.redtext_bluetext {
//		color: blue;
//	}`
func Blend(classes ...CSSClass) {
	if len(classes) == 0 {
		return
	}

	if len(classes) == 1 {
		RawClass(classes[0].Name, renderCSSProps(classes[0].Name, classes[0].Class))
	}

	out := classes[0]
	for i := 1; i < len(classes); i++ {
		out = out.Blend(classes[i])
	}

	RawClass(out.Name, renderCSSProps(out.Name, out.Class))
}

func BlendFrom(names ...string) {
	blendClasses := []CSSClass{}
	for _, name := range names {
		blendClasses = append(blendClasses, classes[name])
	}

	Blend(blendClasses...)
}

// Create a new CSS class from two CSS classes, where the properties of the second class
// overwrite the properties of the first class.
func (a *CSSClass) Blend(b CSSClass) CSSClass {
	return CSSClass{
		Name:  blendClassNames(a.Name, b.Name),
		Class: blendCssProps(a.Class, b.Class),
	}
}

// Blends the CSS properties of two CSS classes into a new CSS class.
// Properties from the second class will overwrite properties from the first class.
func blendCssProps(a CSSProps, b CSSProps) CSSProps {
	props := CSSProps{}

	for k, v := range a {
		props[k] = v
	}

	for k, v := range b {
		props[k] = v
	}

	return props
}

// Blends the class names of two CSS classes into a new CSS class name.
// The class names are separated by a underscore, and contain only letters, numbers and underscores.
func blendClassNames(a string, b string) string {
	return fmt.Sprintf("%s_%s", stripNameForBlend(a), stripNameForBlend(b))
}

func stripNameForBlend(name string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\-_]`)
	return re.ReplaceAllString(name, "_")
}
