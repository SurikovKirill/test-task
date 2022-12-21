package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_get(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "With values in url",
			url:  "/exchange/CHF",
			want: want{
				code: 200,
			},
		},
		{
			name: "Without values in url",
			url:  "/exchange/",
			want: want{
				code: 200,
			},
		},
		{
			name: "Bad values in url",
			url:  "/exchange/URL",
			want: want{
				code: 500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.url, nil)
			request.Header.Set("X-API-KEY", "123321")
			w := httptest.NewRecorder()
			Get(w, request)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != tt.want.code {
				t.Log(tt.want.code, w.Code)
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
		})
	}
}

func Test_curVal(t *testing.T) {
	type args struct {
		tmp []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := curVal(tt.args.tmp); got != tt.want {
				t.Errorf("curVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
