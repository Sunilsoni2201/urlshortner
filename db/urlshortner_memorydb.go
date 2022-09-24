package db

type memoryDb struct {
	urls map[string]string
}

func NewMemoryDb() UrlShortnerDb {
	return &memoryDb{
		urls: make(map[string]string),
	}
}

func (m *memoryDb) Get(key string) (ourl string, err error) {
	ourl = m.urls[key]
	return
}

func (m *memoryDb) Set(key string, ourl string) (err error) {
	m.urls[key] = ourl
	return
}
