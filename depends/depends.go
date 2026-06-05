package depends

import "url-shortener/store"

var s = store.NewMemoryStore()

func GetStore() store.Store {
	if s == nil {
		panic("Error")
	}
	return s
}
