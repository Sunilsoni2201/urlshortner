package db

import (
	"os"
	"strconv"
	"testing"
)

func TestNewFileDb(t *testing.T) {
	db := NewFileDb("~/filedb.json")
	_, err := db.Get("123abc")
	if err == nil {
		t.Errorf("error must be nil")
	}
	os.Remove("~/filedb.json")
}

func TestFileDbSet(t *testing.T) {

	testcases := []struct {
		name string
		key  string
		ourl string

		getkey    string
		want      string
		wantError bool
	}{
		{
			name: "valid value",
			key:  "abcdef",
			ourl: "https://www.hamaraurl.com",

			getkey:    "abcdef",
			want:      "https://www.hamaraurl.com",
			wantError: false,
		},
		{
			name: "invalid value",
			key:  "abc123",
			ourl: "https://www.hamaraurl.com",

			getkey:    "111111",
			want:      "",
			wantError: true,
		},
	}

	db := NewFileDb("")

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			err := db.Set(tc.key, tc.ourl)
			if err != nil {
				t.Errorf("error in set key : %v, url: %v", tc.key, tc.ourl)
			}
			got, err := db.Get(tc.getkey)
			gotError := err != nil
			if tc.wantError != gotError {
				//fail the test
				t.Errorf("got:%v, want:%v", got, tc.want)
			}

			if tc.want != got {
				//fail the test
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		})
	}
}

func BenchmarkFileDbSet(b *testing.B) {

	db := NewFileDb("")

	for i := 0; i < b.N; i++ {
		//_ = db.Set(strconv.Itoa(i), fmt.Sprintf("www.google%v.com", i))
		_ = db.Set(strconv.Itoa(i), "www.google"+strconv.Itoa(i)+".com")
	}
}
