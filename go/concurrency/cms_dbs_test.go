package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testDBS(mode string) *httptest.ResponseRecorder {
	rurl := "/dbs?dataset=/ZMM*"
	action = mode
	req, _ := http.NewRequest("GET", rurl, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DBSHandler)
	handler.ServeHTTP(rr, req)
	return rr
}

func TestDBSSequential(t *testing.T) {
	initDBS()
	rr := testDBS("sequential")
	if status := rr.Code; status != http.StatusOK {
		t.Fatal("invalid HTTP request")
	}
}

func TestDBSConcurrent(t *testing.T) {
	initDBS()
	rr := testDBS("concurrent")
	if status := rr.Code; status != http.StatusOK {
		t.Fatal("invalid HTTP request")
	}
}

// BenchmarkTest
func BenchmarkTest(b *testing.B) {
	initDBS()
	action = "concurrent"
	records := getDatasets("/ZMM*")
	for i := 0; i < b.N; i++ {
		concurrentFunction(records)
	}
}
