package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
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

	ctx := context.Background()
	for _, test := range tests {
		ctx_modified := context.WithValue(ctx, contextUserKey, test.IP)
		res := app.ipFromContext(ctx_modified)
		if res != test.IP {
			t.Error("IP from context does not match")
		}
	}

}

func Test_auth(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	var tests = []struct {
		name   string
		isAuth bool
	}{
		{"logged in", true},
		{"not logged in", false},
	}

	for _, test := range tests {
		handlerToTest := app.auth(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if test.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}

		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if test.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s: expected status code of 200 but got %d", test.name, rr.Code)
		}

		if !test.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("%s: expected redirect but got status code %d", test.name, rr.Code)
		}
	}
}
