package translator

import "languages/pkg/languages"

type Translator interface {
	Hello() string
	Goodbye() string
}

type LanguageTranslator struct {
	Translator
}

func New(language string) *LanguageTranslator {
	switch language {
	case "english":
		return &LanguageTranslator{languages.NewEnglish()}
	case "french":
		return &LanguageTranslator{languages.NewFrench()}
	default:
		return nil
	}
}
