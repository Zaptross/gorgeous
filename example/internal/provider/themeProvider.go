package provider

import "github.com/zaptross/gorgeous/example/internal/types"

type ThemeProviderType struct {
	Theme types.Pallette
}

var ThemeProvider = &ThemeProviderType{}

func NewThemeProvider(theme types.Pallette) *ThemeProviderType {
	ThemeProvider = &ThemeProviderType{
		Theme: theme,
	}

	return ThemeProvider
}

func (tp *ThemeProviderType) GetTheme() types.Pallette {
	return ThemeProvider.Theme
}

func (tp *ThemeProviderType) SetTheme(theme types.Pallette) {
	ThemeProvider.Theme = theme
}
