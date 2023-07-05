package gorgeous

import "testing"

func TestEmpty(t *testing.T) {
	output := RenderElement(Empty(), "")

	if output.Document.String() != "" {
		t.Errorf("Empty() did not render an empty string: %s", output.Document.String())
	}
}

func TestTextTableDriven(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"test", "test"},
		{"<test>", "&lt;test&gt;"},
		{"<test>test</test>", "&lt;test&gt;test&lt;/test&gt;"},
	}

	for _, test := range tests {
		output := RenderElement(Text(test.input), "")

		if output.Document.String() != test.expected {
			t.Errorf("Text(\"%s\") did not render \"%s\", got: \"%s\"", test.input, test.expected, output.Document.String())
		}
	}
}

func TestRawTextTableDriven(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"test", "test"},
		{"<test>", "<test>"},
		{"<test>test</test>", "<test>test</test>"},
	}

	for _, test := range tests {
		output := RenderElement(RawText(test.input), "")

		if output.Document.String() != test.expected {
			t.Errorf("RawText(\"%s\") did not render \"%s\", got: \"%s\"", test.input, test.expected, output.Document.String())
		}
	}
}

func TestBody(t *testing.T) {
	output := RenderElement(Body(EB{}), "")

	if output.Document.String() != "<body id=\"body\"></body>" {
		t.Errorf("Body() did not render \"<body></body>\", got: \"%s\"", output.Document.String())
	}
}
