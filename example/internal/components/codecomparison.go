package components

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/gorgeous/example/internal/provider"
	s "github.com/zaptross/gorgeous/example/internal/services"
)

func CodeComparison(titleElementId string) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	es := s.GetElementService()

	g.Class(".code-comparison-hover", g.CSSProps{
		"background-color":   theme.Base02,
		"box-shadow":         fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
		"-webkit-box-shadow": fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
		"-moz-box-shadow":    fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
	})

	return g.Div(g.EB{
		ClassList: []string{"code-comparison"},
		Children: g.CE{
			Codeblock(CodeblockProps{
				FilePath: "../internal/components/pageTitle.go",
				FileName: "pageTitle.go",
				Language: "go",
			}),
			Codeblock(CodeblockProps{
				FilePath: "../internal/components/reactPageTitle.tsx",
				FileName: "pageTitle.tsx",
				Language: "tsx",
			}),
		},
		Style: g.CSSProps{
			"display":         "flex",
			"flex-direction":  "row",
			"justify-content": "space-between",
			"text-align":      "left",
		},
		Script: g.JavaScript(g.JavaScript(getMouseEventScript(es, titleElementId))),
	})
}

func getMouseEventScript(es string, titleId string) string {
	titleElement := fmt.Sprintf(`%s('%s')`, es, titleId)
	return fmt.Sprintf(`
		thisElement.addEventListener('mouseover', () => {
			%s?.classList.add('code-comparison-hover');
		});
		thisElement.addEventListener('mouseout', () => {
			%s?.classList.remove('code-comparison-hover');
		});
	`, titleElement, titleElement)
}
