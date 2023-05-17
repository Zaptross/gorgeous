package components

import g "github.com/zaptross/gorgeous"

func Favicon() []*g.HTMLElement {
	return []*g.HTMLElement{
		// Add favicon
		g.Link(g.EB{
			Props: g.Props{
				"rel":   "apple-touch-icon",
				"sizes": "180x180",
				"href":  "/public/apple-touch-icon.png",
			},
		}),
		g.Link(g.EB{
			Props: g.Props{
				"rel":   "icon",
				"type":  "image/png",
				"sizes": "32x32",
				"href":  "/public/favicon-32x32.png",
			},
		}),
		g.Link(g.EB{
			Props: g.Props{
				"rel":   "icon",
				"type":  "image/png",
				"sizes": "16x16",
				"href":  "/public/favicon-16x16.png",
			},
		}),
		g.Link(g.EB{
			Props: g.Props{
				"rel":  "manifest",
				"href": "/public/site.webmanifest",
			},
		}),
	}
}
