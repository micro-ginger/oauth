package app

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func (a *app) initializeLanguage() {
	a.language.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, lang := range a.config.Gateway.Language.Languages {
		a.language.MustLoadMessageFile(filepath.Join(a.config.Gateway.Language.Dir, lang))
	}
}
