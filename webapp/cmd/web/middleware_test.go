package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})
	var app application

	for _, t := range tests {
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if t.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(t.headerName) > 0 {
			req.Header.Add(t.headerName, t.headerValue)
		}

		if len(t.addr) > 0 {
			req.RemoteAddr = t.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_ipFromContext(t *testing.T) {

	tests := []struct {
		IP string
	}{
		{"192.3.2.1"},
		{""},
	}

	var app application

	ctx := context.Background()
	for _, test := range tests {
		ctx_modified := context.WithValue(ctx, contextUserKey, test.IP)
		res := app.ipFromContext(ctx_modified)
		if res != test.IP {
			t.Error("IP from context does not match")
		}
	}

}
