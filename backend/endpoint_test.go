package main

import (
	"net/http"
	"testing"
)

func TestEverything(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/country/1")
	if err != nil {
		t.Error(err.Error())
	}
	if resp.StatusCode != 200 {
		t.Error("Status code is not 200")
	}
}
