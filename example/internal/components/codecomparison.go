package components

import g "github.com/zaptross/gorgeous"

func CodeComparison() *g.HTMLElement {
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
	})
}
