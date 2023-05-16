package gorgeous

import "html"

func apply(element *HTMLElement, base EB) *HTMLElement {
	element.EB = base
	return element
}

func createElement(element string) *HTMLElement {
	return &HTMLElement{
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
// content. It is not rendered as a HTML element, but rather as text content of
// the parent element. Eg:
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
func RawText(content string) *HTMLElement {
	return &HTMLElement{
		OpenTag:  "",
		CloseTag: "",
		Text:     content,
	}
}

// Document

func Html(e EB) *HTMLElement {
	return apply(
		&HTMLElement{OpenTag: "<!DOCTYPE html>\n<html",
			CloseTag: "</html>",
		},
		e,
	)
}
func Head(e EB) *HTMLElement { return apply(createElement("head"), e) }
func Body(e EB) *HTMLElement { return apply(createElement("body"), e) }

// Head

func Meta(e EB) *HTMLElement   { return apply(createVoidElement("meta"), e) }
func Link(e EB) *HTMLElement   { return apply(createVoidElement("link"), e) }
func Title(e EB) *HTMLElement  { return apply(createElement("title"), e) }
func Script(e EB) *HTMLElement { return apply(createElement("script"), e) }

// Headings

func H1(e EB) *HTMLElement { return apply(createElement("h1"), e) }
func H2(e EB) *HTMLElement { return apply(createElement("h2"), e) }
func H3(e EB) *HTMLElement { return apply(createElement("h3"), e) }
func H4(e EB) *HTMLElement { return apply(createElement("h4"), e) }
func H5(e EB) *HTMLElement { return apply(createElement("h5"), e) }

// Text

func P(e EB) *HTMLElement          { return apply(createElement("p"), e) }
func Div(e EB) *HTMLElement        { return apply(createElement("div"), e) }
func Span(e EB) *HTMLElement       { return apply(createElement("span"), e) }
func A(e EB) *HTMLElement          { return apply(createElement("a"), e) }
func Img(e EB) *HTMLElement        { return apply(createElement("img"), e) }
func Blockquote(e EB) *HTMLElement { return apply(createElement("blockquote"), e) }

// Interactive

func Button(e EB) *HTMLElement   { return apply(createElement("button"), e) }
func Input(e EB) *HTMLElement    { return apply(createElement("input"), e) }
func Form(e EB) *HTMLElement     { return apply(createElement("form"), e) }
func Label(e EB) *HTMLElement    { return apply(createElement("label"), e) }
func Select(e EB) *HTMLElement   { return apply(createElement("select"), e) }
func Option(e EB) *HTMLElement   { return apply(createElement("option"), e) }
func Textarea(e EB) *HTMLElement { return apply(createElement("textarea"), e) }

// Table

func Table(e EB) *HTMLElement { return apply(createElement("table"), e) }
func Tr(e EB) *HTMLElement    { return apply(createElement("tr"), e) }
func Td(e EB) *HTMLElement    { return apply(createElement("td"), e) }
func Th(e EB) *HTMLElement    { return apply(createElement("th"), e) }
func Thead(e EB) *HTMLElement { return apply(createElement("thead"), e) }
func Tbody(e EB) *HTMLElement { return apply(createElement("tbody"), e) }
func Tfoot(e EB) *HTMLElement { return apply(createElement("tfoot"), e) }

// Lists

func Ul(e EB) *HTMLElement { return apply(createElement("ul"), e) }
func Ol(e EB) *HTMLElement { return apply(createElement("ol"), e) }
func Li(e EB) *HTMLElement { return apply(createElement("li"), e) }
