package gorgeous

import (
	"strings"
	"testing"
)

func clearClasses() {
	classes = map[string]CSSClass{}
}
func clearMedia() {
	media = map[string][]CSS{}
}

func TestClass(t *testing.T) {
	clearClasses()

	Class(&CSSClass{
		Selector: ".my-class",
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

	rawClass := media[tc.query][0]
	if strings.Contains(rawClass.String(), tc.expectedSelector) == false {
		t.Errorf("Media did not add the class to the media query, expected: %s, got: %s", tc.expectedSelector, rawClass.String())
	}

	if strings.Contains(rawClass.String(), tc.expectedProp) == false {
		t.Errorf("Media did not add the class to the media query, expected: %s, got: %s", tc.expectedProp, rawClass.String())
	}
}

func TestRawMedia(t *testing.T) {
	clearMedia()

	tc := struct {
		query       string
		rawCSSClass string
		expected    string
	}{
		query: "(max-width: 600px)",
		rawCSSClass: `.my-class {
			color: red;
		}`,
		expected: `.my-class {
			color: red;
		}`,
	}

	RawMedia(tc.query, CSS(tc.rawCSSClass))

	if len(media) != 1 {
		t.Errorf("RawMedia did not add the media query to the media map")
	}

	if len(media[tc.query]) != 1 {
		t.Errorf("RawMedia did not add the class to the media query")
	}

	rawClass := media[tc.query][0]
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

	if strings.Index(classes.String(), "@media") <= 0 {
		t.Errorf("collectClasses did not collect the classes, expected: >0, got: %d", len(classes))
	}

	if !strings.Contains(classes.String(), "color: red;") && !strings.Contains(classes.String(), "color: blue;") {
		t.Errorf("collectClasses did not collect the classes, expected: %s, got: %s", "color: red;", classes)
	}
}
