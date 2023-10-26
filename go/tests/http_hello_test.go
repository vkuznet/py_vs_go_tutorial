package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatal("invalid HTTP request")
	}
}

// BenchmarkTest
func BenchmarkTest(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a = i
	}
	fmt.Println("a", a)
}
