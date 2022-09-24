package db

import (
	"encoding/json"
	"io/ioutil"
)

type memoryDb struct {
	Urls     map[string]string
	filePath string
}

func NewMemoryDb(filepath string) UrlShortnerDb {
	m := &memoryDb{
		Urls:     make(map[string]string),
		filePath: filepath,
	}
	file, _ := ioutil.ReadFile(filepath)

	_ = json.Unmarshal([]byte(file), m)

	return m
}

func (m *memoryDb) Get(key string) (ourl string, err error) {
	ourl = m.Urls[key]

	return
}

func (m *memoryDb) Set(key string, ourl string) (err error) {
	m.Urls[key] = ourl
	file, _ := json.MarshalIndent(m, "", " ")
	_ = ioutil.WriteFile(m.filePath, file, 0644)
	return
}
