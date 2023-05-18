package gorgeous

type HTML string
type JavaScript string
type CSS string
type Props map[string]string

// CE stands for Child Elements - it is a slice of HTMLElement pointers
// It has been shortened to CE to reduce the amount of typing required to create
type CE = []*HTMLElement

// CSSProps is a map of CSS property names to CSS property values
type CSSProps = map[string]string

// EB stands for Element Base - it contains the common properties of all elements
// It has been shortened to EB to reduce the amount of typing required to create
type EB struct {
	// html element attributes
	Id        string
	OnClick   string
	Style     CSSProps
	ClassList []string

	// html element properties (e.g. type="text" for input elements)
	Props Props

	// Client-side code specific to the element, run when the document is loaded
	// in the browser.
	//
	// Eg:
	//
	//	Id: "element-example",
	//	Script: `thisElement.value = 'Hello, world!'`
	//
	// renders as:
	//
	//	<script id="script-element-example">
	//		((thisElement) => {
	//			thisElement.value = 'Hello, world!';
	//		})(document.getElementById('element-example'));
	//		document.getElementById('script-element-example').remove();
	//	</script>
	Script JavaScript

	// html element children
	Children CE
}

// HTMLElement is the base type for all HTML elements
type HTMLElement struct {
	EB

	// html element text content
	Text string

	// render properties
	OpenTag  string
	CloseTag string
}

// RenderedHTML is the result of calling the Render function, it contains the
// rendered HTML, CSS and JavaScript. These should be written to files or served
// to the client via the appropriate api endpoints.
type RenderedHTML struct {
	Document HTML
	Script   JavaScript
	Services JavaScript
	Style    CSS
}
