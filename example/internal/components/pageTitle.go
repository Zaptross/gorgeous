package components

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/gorgeous/example/internal/provider"
)

type PageTitleProps struct {
	Title string
	Id    string
}

func PageTitle(props PageTitleProps) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	return g.H1(g.EB{
		Children: g.CE{
			g.Text(props.Title),
		},
		Style: g.CSSProps{
			"color":   theme.Base2,
			"margin":  "0 0",
			"padding": "0.67em 0.67em",
		},
		Id: props.Id,
	})
}
