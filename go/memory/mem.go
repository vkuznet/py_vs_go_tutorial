package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"

	validator "github.com/go-playground/validator/v10"
)

func memUsage(m1, m2 *runtime.MemStats) {
	fmt.Printf("Sys: %d (MB) total alloc %d (MB), heap alloc %d (MB)\n",
		(m2.Sys-m1.Sys)/(1024*1024),
		(m2.TotalAlloc-m1.TotalAlloc)/(1024*1024),
		(m2.HeapAlloc-m1.HeapAlloc)/(1024*1024))
}

func readJSON(dst string) {
	fname := os.Args[1]
	fileInfo, _ := os.Stat(fname)
	var r0, r1 runtime.MemStats

	fmt.Printf("# File %s size %d (MB)\n", fname, fileInfo.Size()/(1024*1024))
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// read content of the given file
	body, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if dst == "map" {
		runtime.ReadMemStats(&r0)
		var data map[string]any
		err = json.Unmarshal(body, &data)
		runtime.ReadMemStats(&r1)
		t := reflect.TypeOf(data)
		fmt.Printf("Data %T, %d bytes\n", data, t.Size())
	} else {
		runtime.ReadMemStats(&r0)
		var data Data
		err = json.Unmarshal(body, &data)
		runtime.ReadMemStats(&r1)
		if err == nil {
			validate := validator.New(validator.WithRequiredStructEnabled())
			// explain data struct validation and different scope for err variable
			if err := validate.Struct(data); err != nil {
				log.Fatal(err)
			}
		}
		t := reflect.TypeOf(data)
		fmt.Printf("Data %T, %d bytes\n", data, t.Size())
		//         fmt.Printf("JSON data:\n%+v\n", data)
	}
	if err != nil {
		log.Fatal(err)
	}
	memUsage(&r0, &r1)
}

func main() {
	fmt.Println("\n### read JSON into generic map")
	readJSON("map")
	fmt.Println("\n### read JSON into data struct")
	readJSON("data")
}
