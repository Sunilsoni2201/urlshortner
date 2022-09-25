package services

import (
	"testing"
	"urlshortner/db"
)

func BenchmarkCreateShortURL(b *testing.B) {
	svc := NewURLShortner(db.NewMemoryDb())

	for i := 0; i < b.N; i++ {
		_, _ = svc.CreateShortURL("https://www.google.com")
	}
}
