package components

import (
	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

type PageTitleProps struct {
	Title string
}

func PageTitle(props PageTitleProps) *g.HTMLElement {
	theme := prv.ThemeProvider.GetTheme()

	return g.H1(g.EB{
		Children: g.CE{
			g.Text(props.Title),
		},
		Style: g.CSSProps{
			"color": theme.Base2,
		},
	})
}
