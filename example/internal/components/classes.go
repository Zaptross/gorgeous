package components

import (
	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/gorgeous/example/internal/provider"
)

type ClassPallette struct {
	BlueText g.CSSClass
	BoldText g.CSSClass
}

func GetClassPallette() *ClassPallette {
	return &ClassPallette{
		BlueText: g.CSSClass{
			Selector: ".blue-text",
			Props:    g.CSSProps{"color": prv.ThemeProvider.GetTheme().Cyan},
		},
		BoldText: g.CSSClass{
			Selector: ".bold-text",
			Props:    g.CSSProps{"font-weight": "bold"},
		},
	}
}
