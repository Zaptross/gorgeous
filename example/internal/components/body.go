package components

import (
	"time"

	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func Body() *g.HTMLElement {
	theme := prv.ThemeProvider.GetTheme()
	classes := GetClassPallette()

	titleRef := g.CreateRef("title-element")

	// This is a simple example of how to use the Blend function to create a new CSS class from two existing classes.
	// The new class will have the properties of both classes, with the properties of the second class taking precedence.
	// The base classes are not modified, and because they aren't used elsewhere, they will be tree-shaken from the final CSS.
	boldCyan := g.Blend(classes.BlueText, classes.BoldText)

	return g.Body(g.EB{
		Style: g.CSSProps{
			"text-align":    "center",
			"margin-bottom": "10px",
		},
		Children: g.CE{
			g.Empty(), // This should not render anything in the final HTML
			titleRef.Get(PageTitle(PageTitleProps{Title: "✨ Gorgeous ✨"})),
			g.Div(g.EB{Children: g.CE{
				g.P(g.EB{
					Children:  g.CE{g.Text("Gorgeous is a server-side rendering library for Go, inspired by React and Flutter.")},
					ClassList: []string{boldCyan},
				}),
				g.P(g.EB{
					Children: g.CE{g.Text("The following is a code comparison between the title component on this page and the equivalent component in React.")},
					Style:    g.CSSProps{"color": theme.Base2},
				}),
			}}),
			CodeComparison(titleRef),
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
				Style:    g.CSSProps{"color": theme.Base2, "padding-bottom": "2em"},
			}),
		},
	})
}
