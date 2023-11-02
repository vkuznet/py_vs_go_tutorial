package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/struCoder/pidusage"
)

func sysInfoStat() *pidusage.SysInfo {
	sysInfo, err := pidusage.GetStat(os.Getpid())
	if err != nil {
		log.Fatal(err)
	}
	//     fmt.Printf("# sysInfo %+v\n", sysInfo)
	return sysInfo
}

func mem_usage() (*mem.VirtualMemoryStat, *mem.SwapMemoryStat) {
	//     process, perr := process.NewProcess(int32(os.Getpid()))
	m, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	//     fmt.Printf("# Virtual memory %+v\n# Swap memory: %+v\n", m, s)
	return m, s
}

func readJSON(dst string) {
	mem_usage()
	s0 := sysInfoStat()
	fname := os.Args[1]
	fileInfo, _ := os.Stat(fname)
	m0, _ := mem_usage()
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
		var data map[string]any
		err = json.Unmarshal(body, &data)
	} else {
		var data Data
		err = json.Unmarshal(body, &data)
	}
	if err != nil {
		log.Fatal(err)
	}
	m1, _ := mem_usage()
	fmt.Printf("### Used memory by gopsutil to load JSON %d (MB)\n", (m1.Used-m0.Used)/(1024*1024))
	s1 := sysInfoStat()
	fmt.Printf("### Used memory by pidusage to load JSON %v (MB)\n", (s1.Memory-s0.Memory)/(1024*1024))
}

func main() {
	fmt.Println("### read JSON into generic map")
	readJSON("map")
	fmt.Println("### read JSON into data struct")
	readJSON("data")
}
