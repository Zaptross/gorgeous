package main

import (
	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func registerClasses() {
	theme := prv.ThemeProvider.GetTheme()

	// Because this class is not used anywhere, it will be tree-shaken from the final css.
	g.Class(&g.CSSClass{
		Selector: ".to-be-tree-shaken",
		Props: g.CSSProps{
			"color": theme.Red,
		},
	})

	g.Class(&g.CSSClass{
		Selector: "html, body",
		Include:  true,
		Props: g.CSSProps{
			"background-color": theme.Base03,
			"height":           "100%",
			"display":          "flex",
			"flex-direction":   "column",
			"align-items":      "center",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "body",
		Props: g.CSSProps{
			"width": "60%",
		},
	})

	// Customise the codeblock styling
	g.Class(&g.CSSClass{
		Selector: "pre",
		Include:  true,
		Props: g.CSSProps{
			"background":  "none !important",
			"text-shadow": "none !important",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "code",
		Include:  true,
		Props: g.CSSProps{
			"color":       theme.Base1 + " !important",
			"background":  "none !important",
			"text-shadow": "none !important",
		},
	})
	g.Class(&g.CSSClass{
		Selector: ".operator",
		Include:  true,
		Props: g.CSSProps{
			"background": "none !important",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a",
		Props: g.CSSProps{
			"color":           theme.Cyan,
			"text-decoration": "none",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a:hover",
		Props: g.CSSProps{
			"color":           theme.Violet,
			"text-decoration": "none",
		},
	})

	g.Class(&g.CSSClass{
		Selector: ".code-comparison > div:first-child",
		Props: g.CSSProps{
			"margin-right": "1em",
		},
	})
}
