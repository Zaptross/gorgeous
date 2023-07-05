package gorgeous

import (
	"fmt"
	"sort"
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
		Style:    collectClasses(renderedDocument.Document) + renderedDocument.Style,
	}

	return renderedHtml
}

// Renders as a combined document, embedding css and js into the html document head.
// Note: this assumes that the document contains a head and body at the first level of children,
// eg: using gorgeous.Document()
func RenderPage(document *HTMLElement) *HTML {
	body, head := findBodyAndHead(document)

	replaceText := "{{" + uuid.NewString() + "}}"

	head.Children = append(
		head.Children,
		// Compile the services and classes into the html
		Script(EB{Text: collectServices().String()}),
		Style(EB{Text: replaceText}),
	)

	classes := collectClasses("").String()

	doc := RenderStatic(
		Document(
			head,
			body,
		),
	).Document

	doc = HTML(strings.ReplaceAll(doc.String(), replaceText, classes))

	return &doc
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

	if element.OpenTag == "" && element.CloseTag == "" && element.Text == "" {
		// Render empty element and return early
		return renderedHtml
	}

	if element.OpenTag == "" && element.CloseTag == "" && element.Text != "" {
		// Render text content and return early
		renderedHtml.Document += renderTextContent(element)
		return renderedHtml
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
	if element.OpenTag == "" && element.Text != "" {
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
	text := ""
	if element.Text != "" {
		text = element.Text
	}

	return HTML(fmt.Sprintf(
		`%s %s>%s`,
		element.OpenTag,
		strings.TrimSpace(
			strings.Join(
				removeEmptyStrings(
					[]string{
						renderElementId(element),
						renderStyles(element.EB.Style),
						renderClassList(element.EB.ClassList),
						renderTextProps(element.EB.Props),
					},
				),
				" ",
			),
		),
		text,
	))
}

func removeEmptyStrings(strings []string) []string {
	var result []string

	for _, str := range strings {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
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

	return `style="` + style + `"`
}

func renderClassList(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	return `class="` + strings.Join(classList, " ") + `"`
}

func renderCSSProps(name string, class CSSProps) CSS {
	if len(class) == 0 {
		return CSS("")
	}

	style := CSS("")

	keys := make([]string, len(class))
	for key := range class {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key == "" {
			continue
		}
		style += CSS(fmt.Sprintf("\t%s: %s;\n", key, class[key]))
	}

	return CSS(fmt.Sprintf("%s {\n%s}\n", name, style))
}

func renderTextProps(props Props) string {
	textProps := ""

	if len(props) == 0 {
		return textProps
	}

	keys := []string{}
	for key := range props {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	renderedProps := []string{}

	for _, key := range keys {
		renderedProps = append(renderedProps, fmt.Sprintf(`%s="%s"`, key, strings.ReplaceAll(props[key], `"`, `'`)))
	}

	return strings.Join(renderedProps, " ")
}
