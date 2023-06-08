package gorgeous

type CSS string

func (c CSS) String() string { return string(c) }

type CSSClass struct {
	// The Selector property is used to specify a CSS selector (e.g. "div.my-class").
	// This is useful for specifying complex selectors such as pseudo-classes and
	// pseudo-elements.
	//
	// ⚠ WARNING: The Selector property is not validated, so you must ensure that
	// it is valid CSS.
	Selector string `required:"true"`

	// Whether to include the class in the rendered CSS, even if it is not used in
	// the document. Defaults to false.
	Include bool `default:"false"`

	// The CSS properties of the CSS class.
	Props CSSProps `required:"true"`

	// The raw CSS of the CSS class.
	Raw CSS
}

type HTML string

func (h HTML) String() string { return string(h) }

type JavaScript string

func (j JavaScript) String() string { return string(j) }

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

	// html element text content
	Text string

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

	Deferred bool
}

// HTMLElement is the base type for all HTML elements
type HTMLElement struct {
	EB

	// html element text content, if Text is present in the EB struct then
	// it will override this value.
	Text string

	// render properties

	// The HTML tag name of the element eg: div
	// Primarily used for disambiguating between different elements
	Tag string
	// The opening tag of the element including the angle brace eg: <div
	OpenTag string
	// The closing tag of the element including the angle brace eg: </div>
	CloseTag string

	// Optionally render the element in a deferred manner, this is useful for
	// rendering elements that are not visible on the page when it is first loaded
	// eg: elements that are only visible when the user scrolls down the page
	// or elements that are only visible when the user clicks a button.
	Deferred bool
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
