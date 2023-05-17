package main

import (
	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func registerClasses() {
	theme := prv.ThemeProvider.GetTheme()

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
		"background-color": theme.Base02 + " !important",
		"height":           "100%",
		"border":           "1px solid " + theme.Green,
		"padding":          "0.5rem",
		"font-family":      "monospace !important",
	})

	g.Class("pre", g.CSSProps{
		"background":  "none !important",
		"text-shadow": "none !important",
	})
	g.Class("code", g.CSSProps{
		"color":       theme.Base1 + " !important",
		"background":  "none !important",
		"text-shadow": "none !important",
	})
	g.Class(".operator", g.CSSProps{
		"background": "none !important",
	})

	g.Class("a", g.CSSProps{
		"color":           theme.Cyan,
		"text-decoration": "none",
	})

	g.Class("a:hover", g.CSSProps{
		"color":           theme.Violet,
		"text-decoration": "none",
	})

	g.Class(".code-comparison > div:first-child", g.CSSProps{
		"margin-right": "1em",
	})
}
