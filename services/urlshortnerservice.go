package services

import (
	"crypto/sha256"
	"fmt"
	"time"
	"urlshortner/db"
)

type urlShortner struct {
	db db.UrlShortnerDb
}

type URLShortner interface {
	GetActualURL(string) (string, error)
	CreateShortURL(string) (string, error)
}

func NewURLShortner(db db.UrlShortnerDb) URLShortner {
	return &urlShortner{
		db: db,
	}
}

func (u *urlShortner) GetActualURL(surl string) (ourl string, err error) {
	if surl == "" {
		err = fmt.Errorf("empty short url not allowed")
		return
	}
	ourl, err = u.db.Get(surl)
	if err != nil {
		err = fmt.Errorf("get error from db")
		return
	}

	if ourl == "" {
		err = fmt.Errorf("invalid short url")
		return
	}

	return
}

func (u *urlShortner) CreateShortURL(ourl string) (surl string, err error) {

	if ourl == "" {
		err = fmt.Errorf("empty url not allowed")
		return
	}

	retryCount := 100

	i := 0
	for ; i < retryCount; i++ {
		surl = createURLHash(ourl, 6)
		if _, err1 := u.db.Get(surl); err1 != nil {
			err = u.db.Set(surl, ourl)
			if err != nil {
				err = fmt.Errorf("failed to set in db, %s : %s", surl, ourl)
			}
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
