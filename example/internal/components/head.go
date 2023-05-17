package components

import g "github.com/zaptross/gorgeous"

func Head() *g.HTMLElement {
	return g.Head(g.EB{
		Children: g.CE{
			g.Meta(g.EB{
				Props: g.Props{
					"charset": "UTF-8",
				},
			}),
			g.Meta(g.EB{
				Props: g.Props{
					"name":    "viewport",
					"content": "width=device-width, initial-scale=1.0",
				},
			}),
			g.Title(g.EB{
				Children: g.CE{
					g.Text("Gorgeous"),
				},
			}),
			g.Link(g.EB{
				Props: g.Props{
					"rel":  "stylesheet",
					"type": "text/css",
					"href": "style.css",
				},
			}),
			g.Script(g.EB{
				Props: g.Props{
					"type": "text/javascript",
					"src":  "script.js",
				},
			}),

			// Add syntax highlighting for code blocks
			g.Link(g.EB{
				Props: g.Props{
					"rel":  "stylesheet",
					"type": "text/css",
					"href": "https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/themes/prism.min.css",
				},
			}),
			loadScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/components/prism-core.min.js"),
			loadScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/plugins/autoloader/prism-autoloader.min.js"),
		},
	})
}

func loadScript(src string) *g.HTMLElement {
	return g.Script(g.EB{
		Props: g.Props{
			"type": "text/javascript",
			"src":  src,
		},
	})
}
