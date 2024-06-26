package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	app.DB = &dbrepo.TestDBRepo{}
	os.Exit(m.Run())
}
