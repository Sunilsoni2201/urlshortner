package cache

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type URLCacher interface {
	GetActualURL(string) (string, error)
	CreateShortURL(string) (string, error)
}

type urlCache struct {
	urls map[string]string
}

func Init() URLCacher {
	return &urlCache{
		urls: make(map[string]string),
	}
}

func (u *urlCache) GetActualURL(surl string) (ourl string, err error) {
	if surl == "" {
		err = fmt.Errorf("Empty short url not allowed")
		return
	}
	ourl = u.urls[surl]

	if ourl == "" {
		err = fmt.Errorf("Invalid short url")
		return
	}

	return
}

func (u *urlCache) CreateShortURL(ourl string) (surl string, err error) {

	if ourl == "" {
		err = fmt.Errorf("Empty url not allowed")
		return
	}

	retryCount := 100

	i := 0
	for ; i < retryCount; i++ {
		surl = createURLHash(ourl, 6)
		if _, ok := u.urls[surl]; !ok {
			u.urls[surl] = ourl
			break
		}
	}

	if i == retryCount {
		return surl, fmt.Errorf("max retry limit exceeded to generate short url")
	}

	return
}

func createURLHash(ourl string, len int) string {
	if len == 0 {
		return ""
	}

	salt := time.Now().String()
	sha := sha256.Sum256([]byte(ourl + salt))
	return fmt.Sprintf("%x", sha)[:len]
}
