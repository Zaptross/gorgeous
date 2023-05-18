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
		Script:   "",
		Style:    "",
	}

	// Render text content and return early
	if element.Text != "" {
		renderedHtml.Document += renderTextContent(element)
		return renderedHtml
	}

	renderedChildren := &RenderedHTML{
		Document: "",
		Script:   "",
		Style:    "",
	}

	for _, child := range element.Children {
		renderedChild := RenderElement(child)
		renderedChildren.Document += renderedChild.Document
		renderedChildren.Script += renderedChild.Script
		renderedChildren.Style += renderedChild.Style
	}

	if strings.Contains(element.OpenTag, "body") {
		if element.EB.Props == nil {
			element.EB.Props = Props{}
		}

		element.EB.Props["onload"] = fmt.Sprintf(`(() => { %s %s })()`, renderElementScript(element), renderedChildren.Script)
		element.EB.Script = ""
		renderedChildren.Script = ""
	}

	renderedHtml.Script += renderElementScript(element) + renderedChildren.Script
	renderedHtml.Document += renderElementProps(element) + renderedChildren.Document + HTML(element.CloseTag)

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
		renderElementId(element.EB.Id),
		renderStyles(element.EB.Style),
		renderClassList(element.EB.ClassList),
		renderTextProps(element.EB.Props),
	))
}

func renderElementScript(element *HTMLElement) JavaScript {
	if element.Script == "" {
		return ""
	}

	if element.EB.Id == "" {
		element.EB.Id = uuid.NewString()
	}

	return JavaScript(fmt.Sprintf(
		`((thisElement) => { %s })(document.getElementById('%s'));`,
		element.Script,
		element.Id,
	))
}

func renderElementId(id string) string {
	if id == "" {
		id = uuid.NewString()
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
		textProps += fmt.Sprintf(` %s="%s"`, key, strings.ReplaceAll(value, `"`, `'`))
	}

	return textProps
}
