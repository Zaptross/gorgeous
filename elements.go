package gorgeous

import "html"

func apply(element *HTMLElement, base EB) *HTMLElement {
	element.EB = base
	element.Text = base.Text
	return element
}

func createElement(element string) *HTMLElement {
	return &HTMLElement{
		Tag:      element,
		OpenTag:  "<" + element,
		CloseTag: "</" + element + ">",
	}
}
func createVoidElement(element string) *HTMLElement {
	return &HTMLElement{
		OpenTag:  "<" + element,
		CloseTag: "",
	}
}

//// HTML elements

// framework elements

// The text element is a special case, standing in as a placeholder for text
// content. It is not rendered as a HTML element, but rather as html escaped text
// content of the parent element. Eg:
//
//	Div(&EB{
//	    Children: CE{
//	      Text("Hello, world!"),
//	    },
//	})
//
// renders as:
//
//	`<div>Hello, world!</div>`
func Text(content string) *HTMLElement {
	return &HTMLElement{
		OpenTag:  "",
		CloseTag: "",
		Text:     html.EscapeString(content),
	}
}

// The raw text element is a special case, standing in as a placeholder for text
// content. It is not rendered as a HTML element, but rather as unescaped text content
// of the parent element.
//
// Eg:
//
//	Div(&EB{
//	    Children: CE{
//	      RawText("<b>Hello, world!</b>"),
//	    },
//	})
//
// renders as:
//
//	`<div><b>Hello, world!</b></div>`
func RawText(content string) *HTMLElement {
	return &HTMLElement{
		OpenTag:  "",
		CloseTag: "",
		Text:     content,
	}
}

// The empty element is a special case, standing in as a placeholder for an empty
// element. It is not rendered as a HTML element, but rather as an empty string.
//
// Eg:
//
//	Div(&EB{
//	    Children: CE{
//	      Empty(),
//	    },
//	})
//
// renders as:
//
//	`<div></div>`
func Empty() *HTMLElement {
	return &HTMLElement{
		OpenTag:  "",
		CloseTag: "",
	}
}

// The custom element is a special case, standing in as a placeholder for any element
// not already defined in the framework.
//
// ⚠ It is recommended to use this element sparingly, as it is not type safe, and does
// not provide any of the benefits of the framework.
//
// Eg:
//
//	CustomElement("my-custom-element", true, &EB{
//	    Children: CE{
//	      Text("Hello, world!"),
//	    },
//	})
//
// renders as:
//
//	`<my-custom-element>Hello, world!</my-custom-element>`
func CustomElement(tag string, hasClosingTag bool, e EB) *HTMLElement {
	closeTag := ""
	if hasClosingTag {
		closeTag = "</" + tag + ">"
	}

	return apply(&HTMLElement{
		OpenTag:  "<" + tag,
		CloseTag: closeTag,
	}, e)
}

// Document

// Creates a new HTML document.
// ⚠ This element should only be used once per document.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html
func Html(e EB) *HTMLElement {
	return apply(
		&HTMLElement{OpenTag: "<!DOCTYPE html>\n<html",
			CloseTag: "</html>",
		},
		e,
	)
}

// Creates a new Head element.
// ⚠ This element should only be used once per document.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
func Head(e EB) *HTMLElement { return apply(createElement("head"), e) }

// Creates a new Body element.
// ⚠ This element should only be used once per document.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func Body(e EB) *HTMLElement {
	b := apply(createElement("body"), e)
	b.Id = "body"
	b.EB.Id = "body"
	return b
}

// Head

// Creates a new Meta element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(e EB) *HTMLElement { return apply(createVoidElement("meta"), e) }

// Creates a new Link element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(e EB) *HTMLElement { return apply(createVoidElement("link"), e) }

// Creates a new Title element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(e EB) *HTMLElement { return apply(createElement("title"), e) }

// Creates a new Script element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(e EB) *HTMLElement { return apply(createElement("script"), e) }

// Creates a new Style element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(e EB) *HTMLElement { return apply(createElement("style"), e) }

// Headings

// Creates a new H1 element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H1(e EB) *HTMLElement { return apply(createElement("h1"), e) }

// Creates a new H2 element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H2(e EB) *HTMLElement { return apply(createElement("h2"), e) }

// Creates a new H3 element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H3(e EB) *HTMLElement { return apply(createElement("h3"), e) }

// Creates a new H4 element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H4(e EB) *HTMLElement { return apply(createElement("h4"), e) }

// Creates a new H5 element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H5(e EB) *HTMLElement { return apply(createElement("h5"), e) }

// Text

// Creates a new Paragraph element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func P(e EB) *HTMLElement { return apply(createElement("p"), e) }

// Creates a new Div element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(e EB) *HTMLElement { return apply(createElement("div"), e) }

// Creates a new Span element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(e EB) *HTMLElement { return apply(createElement("span"), e) }

// Creates a new Pre element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Pre(e EB) *HTMLElement { return apply(createElement("pre"), e) }

// Creates a new Br element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Br(e EB) *HTMLElement { return apply(createVoidElement("br"), e) }

// Creates a new Strong element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(e EB) *HTMLElement { return apply(createElement("strong"), e) }

// Creates a new Em element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Em(e EB) *HTMLElement { return apply(createElement("em"), e) }

// Creates a new A element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func A(e EB) *HTMLElement { return apply(createElement("a"), e) }

// Creates a new I element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func I(e EB) *HTMLElement { return apply(createElement("i"), e) }

// Creates a new Img element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Img(e EB) *HTMLElement { return apply(createElement("img"), e) }

// Creates a new Blockquote element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func Blockquote(e EB) *HTMLElement { return apply(createElement("blockquote"), e) }

// Interactive

// Creates a new Button element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(e EB) *HTMLElement { return apply(createElement("button"), e) }

// Creates a new Input element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(e EB) *HTMLElement { return apply(createElement("input"), e) }

// Creates a new Form element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(e EB) *HTMLElement { return apply(createElement("form"), e) }

// Creates a new Label element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(e EB) *HTMLElement { return apply(createElement("label"), e) }

// Creates a new Select element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(e EB) *HTMLElement { return apply(createElement("select"), e) }

// Creates a new Option element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(e EB) *HTMLElement { return apply(createElement("option"), e) }

// Creates a new Textarea element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func Textarea(e EB) *HTMLElement { return apply(createElement("textarea"), e) }

// Creates a new Fieldset element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func Fieldset(e EB) *HTMLElement { return apply(createElement("fieldset"), e) }

// Creates a new Details element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(e EB) *HTMLElement { return apply(createElement("details"), e) }

// Creates a new Summary element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(e EB) *HTMLElement { return apply(createElement("summary"), e) }

// Table

// Creates a new Table element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(e EB) *HTMLElement { return apply(createElement("table"), e) }

// Creates a new Tr element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func Tr(e EB) *HTMLElement { return apply(createElement("tr"), e) }

// Creates a new Td element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func Td(e EB) *HTMLElement { return apply(createElement("td"), e) }

// Creates a new Th element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func Th(e EB) *HTMLElement { return apply(createElement("th"), e) }

// Creates a new Thead element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func Thead(e EB) *HTMLElement { return apply(createElement("thead"), e) }

// Creates a new Tbody element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func Tbody(e EB) *HTMLElement { return apply(createElement("tbody"), e) }

// Creates a new Tfoot element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func Tfoot(e EB) *HTMLElement { return apply(createElement("tfoot"), e) }

// Lists

// Creates a new Ul element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func Ul(e EB) *HTMLElement { return apply(createElement("ul"), e) }

// Creates a new Ol element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func Ol(e EB) *HTMLElement { return apply(createElement("ol"), e) }

// Creates a new Li element.
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func Li(e EB) *HTMLElement { return apply(createElement("li"), e) }
