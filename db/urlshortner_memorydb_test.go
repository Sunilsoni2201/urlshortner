package db

import "testing"

func TestNewMemoryDb(t *testing.T) {

	db := NewMemoryDb()
	_, err := db.Get("abc123")
	if err == nil {
		t.Errorf("error must be nil")
	}
}

func TestMemoryDBSet(t *testing.T) {

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

	db := NewMemoryDb()

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

