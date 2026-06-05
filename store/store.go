package store

type Store interface {
	Save(Shortener)
	Get(code string) *Shortener
	Delete(code string)
	GetAll() []Shortener
}

type Shortener struct {
	Code string
	URL  string
}
