package components

import g "github.com/zaptross/gorgeous"

type PageTitleProps struct {
	Title string
	Theme Pallette
}

func PageTitle(props PageTitleProps) *g.HTMLElement {
	return g.H1(g.EB{
		Children: g.CE{
			g.Text(props.Title),
		},
		Style: g.CSSProps{
			"color": props.Theme.Base2,
		},
	})
}
