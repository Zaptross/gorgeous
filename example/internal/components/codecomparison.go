package components

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/gorgeous/example/internal/provider"
)

func CodeComparison(titleRef *g.Reference) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	g.Class(&g.CSSClass{
		Selector: ".code-comparison-hover",
		Props: g.CSSProps{
			"background-color":   theme.Base02,
			"box-shadow":         fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
			"-webkit-box-shadow": fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
			"-moz-box-shadow":    fmt.Sprintf("inset 0 0 0 2px %s", theme.Green),
		},
	})

	g.Class(&g.CSSClass{
		Selector: ".code-comparison", Props: g.CSSProps{
			"display":         "flex",
			"flex-direction":  "row",
			"justify-content": "space-between",
			"text-align":      "left",
		},
	})

	mediaSquare := "(max-aspect-ratio: 1.2)"

	g.Media(mediaSquare, ".code-comparison", g.CSSProps{
		"flex-direction": "column",
	})

	g.Media(mediaSquare, ".code-comparison div", g.CSSProps{
		"width": "100%",
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
		Style:  g.CSSProps{},
		Script: g.JavaScript(g.JavaScript(getMouseEventScript(titleRef))),
	})
}

func getMouseEventScript(title *g.Reference) string {
	return fmt.Sprintf(`
		thisElement.addEventListener('mouseover', () => {
			%s?.classList.add('code-comparison-hover');
		});
		thisElement.addEventListener('mouseout', () => {
			%s?.classList.remove('code-comparison-hover');
		});
	`, title.Javascript(), title.Javascript())
}
