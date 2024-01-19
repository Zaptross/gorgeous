package gorgeous

import (
	"strings"
	"testing"
)

func clearClasses() {
	classes = map[string]CSSClass{}
}
func clearMedia() {
	media = map[string]map[string]CSS{}
}

func TestClass(t *testing.T) {
	clearClasses()

	selector := ".my-class"

	Class(&CSSClass{
		Selector: selector,
		Props: CSSProps{
			"color": "red",
		},
	})

	if len(classes) != 1 {
		t.Errorf("Class did not add the class to the classes map")
	}

	if classes[".my-class"].Selector != ".my-class" {
		t.Errorf("Class did not set the selector correctly")
	}

	if classes[".my-class"].Props["color"] != "red" {
		t.Errorf("Class did not set the props correctly")
	}

	// test panic when trying to register a different class with the same selector
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("collectClasses did not panic when registering a different class with the same selector")
		}
	}()

	Class(&CSSClass{
		Selector: selector,
		Props: CSSProps{
			"color": "blue",
		},
	})
}

func TestRawClass(t *testing.T) {
	clearClasses()

	tc := struct {
		className   string
		rawCSSClass string
	}{
		className: "my-class",
		rawCSSClass: `.my-class {
			color: red;
		}`,
	}

	RawClass(tc.className, CSS(tc.rawCSSClass))

	if len(classes) != 1 {
		t.Errorf("RawClass did not add the class to the classes map")
	}

	if classes[tc.className].Raw != CSS(tc.rawCSSClass) {
		t.Errorf("RawClass did not set the raw css correctly, expected: %s, got: %s", tc.rawCSSClass, classes[".my-class"].Raw)
	}
}

func TestMedia(t *testing.T) {
	clearMedia()

	tc := struct {
		query            string
		className        string
		key              string
		col              string
		expectedSelector string
		expectedProp     string
	}{
		query:            "(max-width: 600px)",
		className:        ".my-class",
		key:              "color",
		col:              "red",
		expectedSelector: ".my-class",
		expectedProp:     "color: red;",
	}

	Media(tc.query, tc.className, CSSProps{
		tc.key: tc.col,
	})

	if len(media) != 1 {
		t.Errorf("Media did not add the media query to the media map")
	}

	if len(media[tc.query]) != 1 {
		t.Errorf("Media did not add the class to the media query")
	}

	rawClass := media[tc.query][tc.expectedSelector]
	if strings.Contains(rawClass.String(), tc.expectedSelector) == false {
		t.Errorf("Media did not add the class to the media query, expected: %s, got: %s", tc.expectedSelector, rawClass.String())
	}

	if strings.Contains(rawClass.String(), tc.expectedProp) == false {
		t.Errorf("Media did not add the class to the media query, expected: %s, got: %s", tc.expectedProp, rawClass.String())
	}

	// test panic when trying to register a different class with the same selector
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("collectClasses did not panic when registering a different class with the same selector")
		}
	}()

	Media(tc.query, tc.className, CSSProps{
		tc.key: "blue",
	})
}

func TestRawMedia(t *testing.T) {
	clearMedia()

	tc := struct {
		query       string
		selector    string
		rawCSSClass string
		expected    string
	}{
		query:    "(max-width: 600px)",
		selector: ".my-class",
		rawCSSClass: `.my-class {
			color: red;
		}`,
		expected: `.my-class {
			color: red;
		}`,
	}

	RawMedia(tc.query, tc.selector, CSS(tc.rawCSSClass))

	if len(media) != 1 {
		t.Errorf("RawMedia did not add the media query to the media map")
	}

	if len(media[tc.query]) != 1 {
		t.Errorf("RawMedia did not add the class to the media query")
	}

	rawClass := media[tc.query][tc.selector]
	if rawClass.String() != tc.expected {
		t.Errorf("RawMedia did not add the class to the media query, expected: %s, got: %s", tc.expected, rawClass.String())
	}
}

func TestCollectClasses(t *testing.T) {
	clearClasses()
	clearMedia()

	className := "my-class"
	Class(&CSSClass{
		Selector: "." + className,
		Props: CSSProps{
			"color": "red",
		},
	})
	Class(&CSSClass{
		Selector: "." + className,
		Props: CSSProps{
			"color": "red",
		},
	})
	Media("(max-width: 600px)", "."+className, CSSProps{
		"color": "blue",
	})
	Media("(max-width: 600px)", "."+className, CSSProps{
		"color": "blue",
	})

	classes := collectClasses("document doesn't include classes")

	if strings.Index(classes.String(), "@media") > 0 {
		t.Errorf("collectClasses did not tree-shake the classes, expected: 0, got: %d", len(classes))
	}

	classes = collectClasses(HTML(`<div class="` + className + `"></div>`))

	if strings.Index(classes.String(), "."+className) > 0 {
		t.Errorf("collectClasses did not collect the classes, expected: 0, got: %d", len(classes))
	}

	// test that the normal classes are included
	if strings.Index(classes.String(), "@media") <= 0 {
		t.Errorf("collectClasses did not collect the classes, expected: >0, got: %d", len(classes))
	}

	if !strings.Contains(classes.String(), "color: red;") && !strings.Contains(classes.String(), "color: blue;") {
		t.Errorf("collectClasses did not collect the classes, expected: %s, got: %s", "color: red;", classes)
	}

	// test that the normal class is not duplicated
	if len(strings.Split(classes.String(), "color: red;")) != 2 {
		t.Errorf("collectClasses did not tree-shake the classes, expected: 2, got: %d", len(strings.Split(classes.String(), "color: red;")))
	}

	// test that the media class is not duplicated
	if len(strings.Split(classes.String(), "color: blue;")) != 2 {
		t.Errorf("collectClasses did not tree-shake the media classes, expected: 2, got: %d", len(strings.Split(classes.String(), "color: blue;")))
	}
}
