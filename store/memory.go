package store

type MemoryStore struct {
	urls map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls: map[string]string{},
	}
}
func (m *MemoryStore) Save(s Shortener) {
	m.urls[s.Code] = s.URL
}

func (m *MemoryStore) Get(code string) *Shortener {
	url, ok := m.urls[code]

	if !ok {
		return nil
	}

	return &Shortener{
		Code: code,
		URL:  url,
	}
}

func (m *MemoryStore) Delete(code string) {
	delete(m.urls, code)
}

func (m *MemoryStore) GetAll() []Shortener {
	s := []Shortener{}
	for key, value := range m.urls {
		s = append(s, Shortener{
			Code: key,
			URL:  value,
		})
	}
	return s
}
