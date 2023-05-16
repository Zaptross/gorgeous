package components

import (
	"html"
	"os"
	"strings"

	g "github.com/zaptross/gorgeous"
)

type CodeblockProps struct {
	FilePath string
	FileName string
}

const nbsp = "&nbsp;"

func Codeblock(props CodeblockProps) *g.HTMLElement {
	return g.Div(
		g.EB{
			Children: g.CE{
				g.Div(
					g.EB{
						Children:  g.CE{g.RawText(code(props))},
						ClassList: []string{"codeblock"},
					},
				),
			},
		},
	)
}

func code(props CodeblockProps) string {
	raw, err := os.ReadFile(props.FilePath)

	if err != nil {
		panic(err)
	}

	codeString := html.EscapeString(string(raw))
	codeString = strings.ReplaceAll(codeString, "\t", "    ")
	codeString = strings.ReplaceAll(codeString, " ", nbsp)
	codeString = strings.ReplaceAll(codeString, "\n", "<br>")

	return "// " + props.FileName + "<br><br>" + codeString
}
