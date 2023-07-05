package gorgeous

import "testing"

func TestRenderElementId(t *testing.T) {
	output := RenderElement(Div(EB{Id: "test"}), "")

	if output.Document.String() != "<div id=\"test\"></div>" {
		t.Errorf("RenderElement(Empty(), \"test\") did not render \"<div id=\"test\"></div>\", got: \"%s\"", output.Document.String())
	}
}

func TestRenderStylesTableDriven(t *testing.T) {
	tests := []struct {
		input    CSSProps
		expected string
	}{
		{CSSProps{}, ""},
		{CSSProps{"color": "red"}, "style=\"color: red;\""},
		{CSSProps{"color": "red", "font-size": "12px"}, "style=\"color: red;font-size: 12px;\""},
	}

	for _, test := range tests {
		output := renderStyles(test.input)

		if output != test.expected {
			t.Errorf("renderStyles(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderClassListTableDriven(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"test"}, "class=\"test\""},
		{[]string{"test", "test2"}, "class=\"test test2\""},
	}

	for _, test := range tests {
		output := renderClassList(test.input)

		if output != test.expected {
			t.Errorf("renderClassList(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderTextPropsTableDriven(t *testing.T) {
	tests := []struct {
		input    Props
		expected string
	}{
		{Props{}, ""},
		{Props{"test": "test"}, "test=\"test\""},
		{Props{"test": "test", "test2": "test2"}, "test=\"test\" test2=\"test2\""},
		// tests that the rendered props are sorted alphabetically
		{Props{"zzz": "zzz", "aaa": "aaa"}, "aaa=\"aaa\" zzz=\"zzz\""},
	}

	for _, test := range tests {
		output := renderTextProps(test.input)

		if output != test.expected {
			t.Errorf("renderTextProps(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderCSSPropsTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    CSSProps
		expected CSS
	}{
		{"", CSSProps{}, CSS("")},
		{"test", CSSProps{"color": "red"}, CSS("test {\n\tcolor: red;\n}\n")},
		// tests that the rendered CSS is sorted alphabetically
		{"test", CSSProps{"font-size": "12px", "color": "red"}, CSS("test {\n\tcolor: red;\n\tfont-size: 12px;\n}\n")},
	}

	for _, test := range tests {
		output := renderCSSProps(test.name, test.input)

		if output != test.expected {
			t.Errorf("renderCSSProps(%s, %v) did not render \"%s\", got: \"%s\"", test.name, test.input, test.expected, output)
		}
	}
}

func TestRenderElementScript(t *testing.T) {
	output := RenderElement(Script(EB{
		Id: "test",
		Children: CE{
			Text("test"),
		},
	}), "")

	if output.Document.String() != "<script id=\"test\">test</script>" {
		t.Errorf("RenderElement(Script()) did not render \"<script id=\"test\">test</script>\", got: \"%s\"", output.Document.String())
	}
}

func TestRenderElementTableDriven(t *testing.T) {
	tests := []struct {
		input    *HTMLElement
		expected string
	}{
		{Empty(), ""},
		{Div(EB{Id: "a"}), "<div id=\"a\"></div>"},
		{P(EB{Id: "a", Style: CSSProps{"color": "red"}}), "<p id=\"a\" style=\"color: red;\"></p>"},
		{P(EB{Id: "a", ClassList: []string{"test"}}), "<p id=\"a\" class=\"test\"></p>"},
		{P(EB{Id: "a", Props: Props{"test": "test"}}), "<p id=\"a\" test=\"test\"></p>"},
		{P(EB{Id: "a", Script: "test"}), "<p id=\"a\"></p><script id=\"script-a\">\n\t\t\t((thisElement) => { test })(document.getElementById('a'));\n\t\t\tdocument.getElementById('script-a').remove();\n\t\t</script>"},
	}

	for _, test := range tests {
		output := RenderElement(test.input, "")

		if output.Document.String() != test.expected {
			t.Errorf("RenderElement(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output.Document.String())
		}
	}
}

func TestRemoveEmptyStringsTableDriven(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"test"}, []string{"test"}},
		{[]string{"test", ""}, []string{"test"}},
		{[]string{"", "test"}, []string{"test"}},
		{[]string{"", "test", ""}, []string{"test"}},
		{[]string{"test", "", "test"}, []string{"test", "test"}},
		{[]string{"", "test", "", "test", ""}, []string{"test", "test"}},
	}

	for _, test := range tests {
		output := removeEmptyStrings(test.input)

		if len(output) != len(test.expected) {
			t.Errorf("removeEmptyStrings(%v) did not return %v, got: %v", test.input, test.expected, output)
		}

		for i, str := range output {
			if str != test.expected[i] {
				t.Errorf("removeEmptyStrings(%v) did not return %v, got: %v", test.input, test.expected, output)
			}
		}
	}
}
