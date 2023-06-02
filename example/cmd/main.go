package main

import (
	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/gorgeous/example/internal/components"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

func main() {
	prv.NewThemeProvider(c.SolarizedDark)
	registerClasses()
	rendered := g.RenderPage(
		g.Document(
			c.Head(),
			c.Body(),
		),
	)

	// Render distribution files
	createDistDirectories()
	writeRenderedHTML(rendered)
	copyPublicToDist()
}
