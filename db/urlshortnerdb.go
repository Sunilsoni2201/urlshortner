package db

type UrlShortnerDb interface {
	Get(string) (string, error)
	Set(string, string) error
}
