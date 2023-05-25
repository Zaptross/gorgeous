package components

import (
	"time"

	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func Body() *g.HTMLElement {
	theme := prv.ThemeProvider.GetTheme()

	titleId := "title-element"

	return g.Body(g.EB{
		Style: g.CSSProps{
			"text-align": "center",
		},
		Children: g.CE{
			PageTitle(PageTitleProps{Id: titleId, Title: "✨ Gorgeous ✨"}),
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
			CodeComparison(titleId),
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
			g.P(g.EB{
				Deferred: true,
				Children: g.CE{g.Text("This document was rendered at " + time.Now().Format("03:04:05 PM") + ". This page was rendered at ${new Date().toLocaleTimeString()}")},
				Script:   g.JavaScript("thisElement.innerText += `, appended by script at: ${new Date().toLocaleTimeString()}.`"),
				Style:    g.CSSProps{"color": theme.Base2, "margin-block-start": "0"},
			}),
		},
	})
}
