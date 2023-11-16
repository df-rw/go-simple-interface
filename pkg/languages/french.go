package languages

type French struct {
}

func (e *French) Hello() string {
	return "Bon jour"
}

func (e *French) Goodbye() string {
	return "Au revoir"
}

func NewFrench() *French {
	return &French{}
}
