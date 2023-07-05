package gorgeous

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func renderElementDeferred(element *HTMLElement, parentId string) *RenderedHTML {
	rendered := &RenderedHTML{
		Document: "",
		Script:   "",
		Services: "",
	}

	if element.OpenTag == "" && element.CloseTag == "" && element.Text == "" {
		// Render empty element and return early
		return rendered
	}

	if element.EB.Id == "" {
		element.EB.Id = uuid.New().String()
	}

	for _, child := range element.Children {
		if child.Text != "" {
			rendered.Document += HTML(child.Text)
			continue
		}

		renderedChild := renderElementDeferred(child, element.EB.Id)
		rendered.Script += renderedChild.Script
	}

	rendered.Script = renderElementDeferredProps(element, parentId, rendered.Document) + rendered.Script
	rendered.Document = ""

	return rendered
}

func renderElementDeferredProps(element *HTMLElement, parentId string, innerText HTML) JavaScript {
	if element.EB.Id == "" {
		element.EB.Id = uuid.New().String()
	}

	// TODO - this should probably be refactored into a default service function
	// TODO - this needs some way to handle rerendering elements of the sub tree
	return JavaScript(fmt.Sprintf(
		`{
		const ele = document.createElement("%s");
		ele.id = "%s";
		ele.onclick = "%s";
		ele.style = "%s";
		%s
		%s
		ele.innerText = %s;
		document.getElementById("%s").appendChild(ele);

		addEventListener("load", () => ((thisElement) => { %s })(document.getElementById('%s')));
		}
		`,
		element.Tag,
		element.EB.Id,
		element.EB.OnClick,
		renderDeferredStyles(element.EB.Style),
		renderDeferredClassList(element.EB.ClassList),
		renderElementDeferredTextProps(element.EB.Props),
		"`"+innerText+"`",
		parentId,
		element.EB.Script,
		element.EB.Id,
	))
}

func renderDeferredStyles(styles CSSProps) string {
	if len(styles) == 0 {
		return ""
	}

	var style string

	for key, value := range styles {
		style += fmt.Sprintf(`%s: %s;`, key, value)
	}

	return style
}

func renderDeferredClassList(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	return fmt.Sprintf(`		"%s".split(" ").map(c => ele.classList.add(c));`, strings.Join(classList, " "))
}

func renderElementDeferredTextProps(props Props) JavaScript {
	if len(props) == 0 {
		return ""
	}

	var script JavaScript

	for key, value := range props {
		script += JavaScript(fmt.Sprintf(`ele.%s = "%s";\n`, key, value))
	}

	return script
}
