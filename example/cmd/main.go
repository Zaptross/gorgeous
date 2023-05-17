package main

import (
	"os"

	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/gorgeous/example/internal/components"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func main() {
	prv.NewThemeProvider(c.SolarizedDark)
	registerClasses()
	rendered := g.RenderDocument(
		g.Document(
			c.Head(),
			getBody(),
		),
	)

	os.Mkdir("dist", 0755)
	os.WriteFile("dist/index.html", []byte(rendered.Document), 0644)
	os.WriteFile("dist/style.css", []byte(rendered.Style), 0644)
	os.WriteFile("dist/script.js", []byte(rendered.Script), 0644)
}

func getBody() *g.HTMLElement {
	theme := prv.ThemeProvider.GetTheme()

	return g.Body(g.EB{
		Style: g.CSSProps{
			"text-align": "center",
		},
		Children: g.CE{
			c.PageTitle(c.PageTitleProps{Title: "✨ Gorgeous ✨"}),
			g.Div(g.EB{Children: g.CE{
				g.P(g.EB{
					Children: g.CE{g.Text("Gorgeous is a server-side rendering library for Go, inspired by React and Flutter.")},
					Style:    g.CSSProps{"color": theme.Cyan, "font-weight": "bold"},
				}),
				g.P(g.EB{
					Children: g.CE{g.Text("The following is a code comparison between the title component on this page and the equivalent component in React.")},
					Style:    g.CSSProps{"color": theme.Base2},
				}),
			}}),
			c.CodeComparison(),
			g.P(g.EB{
				Children: g.CE{
					g.Text("It's still in early development, but you can check out the source code for this page on "),
					g.A(g.EB{
						Props: g.Props{
							"href":   "https://github.com/zaptross/gorgeous",
							"target": "_blank",
						},
						Children: g.CE{g.RawText("GitHub &#x1f855;")},
					}),
					g.Text("."),
				},
				Style: g.CSSProps{"color": theme.Base2, "margin-top": "2em"},
			}),
		},
	})
}
