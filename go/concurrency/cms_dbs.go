package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/vkuznet/x509proxy"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// global variable to use
var (
	dbsUrl string
	action string
)

// HttpClient is HTTP client for our server
func HttpClient() (*http.Client, error) {
	uproxy := os.Getenv("X509_USER_PROXY")
	uckey := os.Getenv("X509_USER_KEY")
	ucert := os.Getenv("X509_USER_CERT")
	var err error
	var x509cert tls.Certificate
	if uproxy != "" {
		x509cert, err = x509proxy.LoadX509Proxy(uproxy)
	} else if uckey != "" {
		x509cert, err = tls.LoadX509KeyPair(ucert, uckey)
	} else {
		return nil, errors.New("Neither X509_USER_PROXY or X509_USER_KEY is set")
	}
	if err != nil {
		return nil, err
	}
	certs := []tls.Certificate{x509cert}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: certs,
			InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}, nil
}

// Record represents DBS record
type Record struct {
	Dataset string `json:"dataset"`
}

func httpRequest(rurl string) []byte {
	// use custom http client
	req, err := http.NewRequest("GET", rurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	client, err := HttpClient()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read body of HTTP response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getDatasets(dataset string) []Record {
	// construct query to DBS server to fetch datasets info
	rurl := fmt.Sprintf("%s/datasets?dataset=%s", dbsUrl, dataset)
	fmt.Println("call", rurl)

	// we cannot use http.Get as it is not handles X509 certs
	//     resp, err := http.Get(rurl)
	// instead we will use our custom httpRequest function
	body := httpRequest(rurl)
	fmt.Println("HTTP response body", string(body))
	var records []Record
	err := json.Unmarshal(body, &records)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("parsed response", records)
	return records
}

// /dbs?dataset=/ZMM*
func DBSHandler(w http.ResponseWriter, r *http.Request) {
	dataset := r.URL.Query().Get("dataset")
	records := getDatasets(dataset)
	if action == "concurrent" {
		concurrentFunction(records)
	} else {
		sequentialFunction(records)
	}

	w.WriteHeader(http.StatusOK)
}

// timing function will provide elapsed time
func timing(start time.Time, msg string) {
	fmt.Printf("%s elapsed time %v\n", msg, time.Since(start))
}

func sequentialFunction(records []Record) {
	defer timing(time.Now(), "sequential")
	for _, rec := range records {
		rurl := fmt.Sprintf("%s/datasets?dataset=%s", dbsUrl, rec.Dataset)
		body := httpRequest(rurl)
		fmt.Printf("%s\n%+v", rurl, string(body))
	}
}

func concurrentFunction(records []Record) {
	defer timing(time.Now(), "concurrent")
	var wg sync.WaitGroup
	for _, rec := range records {
		wg.Add(1) // for each goroutine we'll add counter
		go func() {
			msg := fmt.Sprintf("%s call", rec.Dataset)
			defer timing(time.Now(), msg)
			defer wg.Done() // always exit goroutine, regardless of errors
			rurl := fmt.Sprintf("%s/datasets?dataset=%s", dbsUrl, rec.Dataset)
			body := httpRequest(rurl)
			fmt.Printf("%s\n%+v", rurl, string(body))
		}()
	}
	// wait for all goroutines to complete their job
	wg.Wait()
}

func initDBS() {
	dbsUrl = os.Getenv("DBS_URL")
	if dbsUrl == "" {
		log.Fatal("No DBS_URL found in your environment")
	}
}

func main() {
	// input parameters for our executable
	flag.StringVar(&action, "action", "", "dbs action URL")
	flag.Parse()

	// init our DBS url
	initDBS()

	// verbose log output
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// web handlers
	http.HandleFunc("/", RequestHandler)
	http.HandleFunc("/dbs", DBSHandler)

	// web server
	http.ListenAndServe(":8888", nil)
}
