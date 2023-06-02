package gorgeous

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// Deprecated: use RenderStatic instead
func RenderDocument(document *HTMLElement) *RenderedHTML {
	panic("RenderDocument is deprecated, use RenderStatic instead")
}

// Renders as three separate file contents for html, css, and js
func RenderStatic(document *HTMLElement) *RenderedHTML {
	renderedDocument := RenderElement(document, "body")

	renderedHtml := &RenderedHTML{
		Document: renderedDocument.Document,
		Script:   collectServices(),
		Style:    collectClasses() + renderedDocument.Style,
	}

	return renderedHtml
}

// Renders as a combined document, embedding css and js into the html document head.
// Note: this assumes that the document contains a head and body at the first level of children,
// eg: using gorgeous.Document()
func RenderPage(document *HTMLElement) *HTML {
	body, head := findBodyAndHead(document)

	head.Children = append(
		head.Children,
		Script(
			EB{
				Children: CE{Text(string(collectServices()))},
			},
		),
		Style(
			EB{
				Children: CE{Text(string(collectClasses()))},
			},
		),
	)

	return &RenderStatic(
		Document(
			head,
			body,
		),
	).Document
}

// Search the top level of the document to render for the head and body elements
// otherwise assume the document is a body element
func findBodyAndHead(document *HTMLElement) (*HTMLElement, *HTMLElement) {
	var body, head *HTMLElement

	for _, child := range document.Children {
		if child.Tag == "body" {
			body = child
		} else if child.Tag == "head" {
			head = child
		}
	}

	return body, head
}

func RenderElement(element *HTMLElement, parentId string) *RenderedHTML {
	renderedHtml := &RenderedHTML{
		Document: "",
		Style:    "",
	}

	if element.EB.Id == "" {
		element.EB.Id = uuid.New().String()
	}

	// Render deferred element and return early
	if element.EB.Deferred {
		deferredElement := renderElementDeferred(element, parentId)
		renderedHtml.Document += HTML(fmt.Sprintf(`<script id="script-%s">
		%s
		document.getElementById('script-%s').remove();
		</script>`, element.EB.Id, deferredElement.Script, element.EB.Id))
		return renderedHtml
	}

	// Render text content and return early
	if element.Text != "" {
		renderedHtml.Document += renderTextContent(element)
		return renderedHtml
	}

	renderedChildren := &RenderedHTML{
		Document: "",
		Style:    "",
	}

	for _, child := range element.Children {
		renderedChild := RenderElement(child, element.EB.Id)
		renderedChildren.Document += renderedChild.Document
		renderedChildren.Style += renderedChild.Style
	}

	renderedHtml.Document += renderElementProps(element) + renderedChildren.Document + HTML(element.CloseTag) + renderElementScript(element)

	renderedHtml.Style += renderedChildren.Style

	return renderedHtml
}

func renderTextContent(element *HTMLElement) HTML {
	return HTML(fmt.Sprintf(
		`%s%s%s`,
		element.OpenTag,
		element.Text,
		element.CloseTag,
	))
}

func renderElementProps(element *HTMLElement) HTML {
	return HTML(fmt.Sprintf(
		`%s %s %s %s %s >`,
		element.OpenTag,
		renderElementId(element),
		`style="`+renderStyles(element.EB.Style)+`"`,
		`class="`+renderClassList(element.EB.ClassList)+`"`,
		renderTextProps(element.EB.Props),
	))
}

func renderElementScript(element *HTMLElement) HTML {
	if element.Script == "" {
		return ""
	}

	if element.EB.Id == "" {
		element.EB.Id = uuid.NewString()
	}

	return HTML(fmt.Sprintf(
		`<script id="script-%s">
			((thisElement) => { %s })(document.getElementById('%s'));
			document.getElementById('script-%s').remove();
		</script>`,
		element.Id,
		element.Script,
		element.Id,
		element.Id,
	))
}

func renderElementId(element *HTMLElement) string {
	if element.EB.Id == "" {
		element.EB.Id = uuid.NewString()
	}

	return fmt.Sprintf(`id="%s"`, element.EB.Id)
}

func renderStyles(styles CSSProps) string {
	if len(styles) == 0 {
		return ""
	}

	var style string

	for key, value := range styles {
		style += fmt.Sprintf(`%s: %s;`, key, value)
	}

	return style
}

func renderClassList(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	var classes string

	for _, class := range classList {
		classes += fmt.Sprintf(`%s `, class)
	}

	return classes
}

func renderCSSProps(name string, class CSSProps) CSS {
	style := CSS("")

	for key, value := range class {
		style += CSS(fmt.Sprintf("\t%s: %s;\n", key, value))
	}

	return CSS(fmt.Sprintf("%s {\n%s}\n", name, style))
}

func renderTextProps(props Props) string {
	textProps := ""

	if len(props) == 0 {
		return textProps
	}

	for key, value := range props {
		textProps += fmt.Sprintf(` %s="%s"`, key, strings.ReplaceAll(value, `"`, `'`))
	}

	return textProps
}
