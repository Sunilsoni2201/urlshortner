package db

import (
	"fmt"
)

type memoryDb struct {
	urls map[string]string
}

func NewMemoryDb() UrlShortnerDb {
	m := &memoryDb{
		urls: make(map[string]string),
	}
	return m
}

func (m *memoryDb) Get(key string) (ourl string, err error) {
	ok := false
	if ourl, ok = m.urls[key]; !ok {
		err = fmt.Errorf("key %s not found in db", key)
	}
	return
}

func (m *memoryDb) Set(key string, ourl string) (err error) {
	m.urls[key] = ourl
	return
}
