package languages

type English struct {
}

func (e *English) Hello() string {
	return "Hello"
}

func (e *English) Goodbye() string {
	return "Goodbye"
}

func NewEnglish() *English {
	return &English{}
}
