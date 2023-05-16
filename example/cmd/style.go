package main

import (
	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/gorgeous/example/cmd/internal/components"
)

func registerClasses(theme c.Pallette) {
	g.Class("html, body", g.CSSProps{
		"background-color": theme.Base03,
		"height":           "100%",
		"display":          "flex",
		"flex-direction":   "column",
		"align-items":      "center",
	})
	g.Class("body", g.CSSProps{
		"width": "60%",
	})
	g.Class(".codeblock", g.CSSProps{
		"background-color": theme.Base02,
		"color":            theme.Base0,
		"height":           "100%",
		"border":           "1px solid " + theme.Green,
		"padding":          "0.5rem",
		"font-family":      "monospace",
	})
	g.Class("a", g.CSSProps{
		"color":           theme.Cyan,
		"text-decoration": "none",
	})
	g.Class("a:hover", g.CSSProps{
		"color":           theme.Violet,
		"text-decoration": "none",
	})
}
