package gorgeous

import (
	"testing"
)

func TestRenderDeferredStylesTableDriven(t *testing.T) {
	tests := []struct {
		input    CSSProps
		expected string
	}{
		{CSSProps{}, ""},
		{CSSProps{"color": "red"}, "color: red;"},
		{CSSProps{"color": "red", "font-size": "12px"}, "color: red;font-size: 12px;"},
	}

	for _, test := range tests {
		output := renderDeferredStyles(test.input)

		if output != test.expected {
			t.Errorf("renderDeferredStyles(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderDeferredClassListTableDriven(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"test"}, `		"test".split(" ").map(c => ele.classList.add(c));`},
		{[]string{"test", "test2"}, `		"test test2".split(" ").map(c => ele.classList.add(c));`},
	}

	for _, test := range tests {
		output := renderDeferredClassList(test.input)

		if output != test.expected {
			t.Errorf("renderDeferredClassList(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderElementDeferredTextPropsTableDriven(t *testing.T) {
	tests := []struct {
		input    Props
		expected string
	}{
		{Props{}, ""},
		{Props{"id": "test"}, `ele.id = "test";\n`},
		// test that onclick is handled separately
		{Props{"id": "test", "onclick": "test"}, `ele.id = "test";\n`},
	}

	for _, test := range tests {
		output := renderElementDeferredTextProps(test.input)

		if output.String() != test.expected {
			t.Errorf("renderElementDeferredTextProps(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}

func TestRenderDeferredOnClickTableDriven(t *testing.T) {
	tests := []struct {
		OnClick      string
		PropsOnClick string
		expected     string
	}{
		{"", "", ""},
		{"clip()", "", `		ele.onclick = "clip()";`},
		{"", "clip()", `		ele.onclick = "clip()";`},
		{"clip()", "clip()", `		ele.onclick = "clip();clip()";`},
	}

	for _, test := range tests {
		output := renderDeferredOnClick(JavaScript(test.OnClick), Props{"onclick": test.PropsOnClick})

		if output != test.expected {
			t.Errorf("renderDeferredOnClick(%v) did not render \"%s\", got: \"%s\"", test.OnClick, test.expected, output)
		}
	}
}

func TestRenderElementDeferredTableDriven(t *testing.T) {
	tests := []struct {
		input    *HTMLElement
		expected string
	}{
		{Empty(), ""},
		// {
		// 	Div(EB{Id: "test"}),
		// 	`const ele = document.createElement("div");`,
		// },
	}

	for _, test := range tests {
		output := renderElementDeferred(test.input, "")

		if output.Script.String() != test.expected {
			t.Errorf("renderElementDeferred(%v) did not render \"%s\", got: \"%s\"", test.input, test.expected, output)
		}
	}
}
