package components

import (
	"html"
	"os"

	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

type CodeblockProps struct {
	FilePath string
	FileName string
	Language string
}

func Codeblock(props CodeblockProps) *g.HTMLElement {
	theme := prv.ThemeProvider.GetTheme()

	g.Class(&g.CSSClass{
		Selector: ".codeblock",
		Props: g.CSSProps{
			"background-color": theme.Base02 + " !important",
			"height":           "100%",
			"border":           "1px solid " + theme.Green,
			"padding":          "0.5rem",
			"font-family":      "monospace !important",
		},
	})

	return g.Div(
		g.EB{
			Children: g.CE{
				g.Pre(g.EB{Children: g.CE{
					g.CustomElement("code", true, g.EB{
						Children:  g.CE{g.RawText(code(props))},
						ClassList: []string{"language-" + props.Language},
					}),
				}}),
			},
			ClassList: []string{"codeblock"},
		},
	)
}

func code(props CodeblockProps) string {
	raw, err := os.ReadFile(props.FilePath)

	if err != nil {
		panic(err)
	}

	codeString := html.EscapeString(string(raw))

	return "// " + props.FileName + "\n" + codeString
}
