package gorgeous

import (
	"fmt"

	"github.com/google/uuid"
)

func RenderDocument(document *HTMLElement) *RenderedHTML {
	renderedDocument := RenderElement(document)

	renderedHtml := &RenderedHTML{
		Document: renderedDocument.Document,
		Script:   collectServices() + renderedDocument.Script,
		Style:    collectClasses() + renderedDocument.Style,
	}

	return renderedHtml
}

func RenderElement(element *HTMLElement) *RenderedHTML {
	renderedHtml := &RenderedHTML{
		Document: "",
		Script:   "",
		Style:    "",
	}

	// Render text content and return early
	if element.Text != "" {
		renderedHtml.Document += renderTextContent(element)
		return renderedHtml
	}

	renderedHtml.Document += renderElementProps(element)

	for _, child := range element.Children {
		renderedChild := RenderElement(child)
		renderedHtml.Document += renderedChild.Document
		renderedHtml.Script += renderedChild.Script
		renderedHtml.Style += renderedChild.Style
	}

	renderedHtml.Document += HTML(element.CloseTag)
	renderedHtml.Script += JavaScript("") // TODO: add script rendering
	renderedHtml.Style += CSS("")         // TODO: add style rendering

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
		`%s id="%s" %s %s %s >`,
		element.OpenTag,
		renderElementId(element.EB.Id),
		renderStyles(element.EB.Style),
		renderClassList(element.EB.ClassList),
		renderTextProps(element.EB.Props),
	))
}

func renderElementId(id string) string {
	if id == "" {
		return uuid.NewString()
	}

	return fmt.Sprintf(`id="%s"`, id)
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
		textProps += fmt.Sprintf(`%s="%s"`, key, value)
	}

	return textProps
}
