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
		renderElementTagFromOpenTag(element.OpenTag),
		element.EB.Id,
		element.EB.OnClick,
		renderStyles(element.EB.Style),
		renderDeferredClassList(element.EB.ClassList),
		renderElementDeferredTextProps(element.EB.Props),
		"`"+innerText+"`",
		parentId,
		element.EB.Script,
		element.EB.Id,
	))
}

func renderElementTagFromOpenTag(openTag string) string {
	if openTag == "" {
		return ""
	}
	return strings.ReplaceAll(strings.ReplaceAll(openTag, "<", ""), " ", "")
}

func renderDeferredClassList(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	return fmt.Sprintf(`		"%s".split(" ").map(c => ele.classList.add(c));`, renderClassList(classList))
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
