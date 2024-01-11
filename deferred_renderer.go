package gorgeous

import (
	"fmt"
	"sort"
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
		%s
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
		renderDeferredOnClick(element.EB.OnClick, element.EB.Props),
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

	keys := make([]string, 0, len(styles))
	for key := range styles {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var style string

	for _, key := range keys {
		style += fmt.Sprintf(`%s: %s;`, key, styles[key])
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

	keys := make([]string, 0, len(props))
	for key := range props {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var script JavaScript

	for _, key := range keys {
		if key == "onclick" {
			continue // onclick is handled separately
		}
		script += JavaScript(fmt.Sprintf(`ele.%s = "%s";\n`, key, props[key]))
	}

	return script
}

func renderDeferredOnClick(onClick JavaScript, props Props) string {
	if onClick == "" && props != nil && props["onclick"] == "" {
		return ""
	}

	onClickScript := ""
	if props != nil && props["onclick"] != "" {
		onClickScript = props["onclick"]
	}
	if onClickScript != "" && onClick != "" {
		onClickScript += ";"
	}
	if onClick != "" {
		onClickScript += string(onClick)
	}

	return fmt.Sprintf(`		ele.onclick = "%s";`, onClickScript)
}
