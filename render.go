package gorgeous

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func RenderDocument(document *HTMLElement) *RenderedHTML {
	renderedDocument := RenderElement(document)

	renderedHtml := &RenderedHTML{
		Document: renderedDocument.Document,
		Script:   collectServices(),
		Style:    collectClasses() + renderedDocument.Style,
	}

	return renderedHtml
}

func RenderElement(element *HTMLElement) *RenderedHTML {
	renderedHtml := &RenderedHTML{
		Document: "",
		Style:    "",
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
		renderedChild := RenderElement(child)
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
		renderStyles(element.EB.Style),
		renderClassList(element.EB.ClassList),
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

	return fmt.Sprintf(`style="%s"`, style)
}

func renderClassList(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	var classes string

	for _, class := range classList {
		classes += fmt.Sprintf(`%s `, class)
	}

	return fmt.Sprintf(`class="%s"`, classes)
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
