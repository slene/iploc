package main

import (
	"fmt"
	"github.com/slene/iploc"
	"os"
	"path/filepath"
	. "testing"
)

func init() {
	// replace iplocFilePath to your iploc.dat path
	iplocFilePath, _ := filepath.Abs("src/github.com/slene/iploc/iploc.dat")

	// read iploc.dat into memory
	iploc.IpLocInit(iplocFilePath)
}

func testIp(ipAddr string) {
	ipinfo, err := iploc.GetIpInfo(ipAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(ipinfo.Ip)
	fmt.Println(ipinfo.IpLong)

	switch ipinfo.Flag {
	case iploc.FLAG_INUSE:
		if ipinfo.Code == "CN" {
			fmt.Println(ipinfo.Code)
			fmt.Println(ipinfo.Country)
			fmt.Println(ipinfo.Region)
			fmt.Println(ipinfo.City)
			fmt.Println(ipinfo.Isp)
		} else {
			fmt.Println(ipinfo.Code)
			fmt.Println(ipinfo.Country)
		}
	case iploc.FLAG_RESERVED:
		fmt.Println(ipinfo.Note)
	case iploc.FLAG_NOTUSE:
		fmt.Println(ipinfo.Note)
	}

	for i := 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func testSpeed() {
	r := Benchmark(func(b *B){
		ips := []string{
			"0.0.0.0",
			"127.0.0.1",
			"169.254.0.1",
			"192.168.1.1",
			"10.0.0.0",
			"255.255.255.255",
			"112.226.155.1",
			"121.18.72.0",
			"6.18.72.0",
			"200.18.72.0",
		}
		for i := 0; i < b.N; i++ {
			for _, ipAddr := range ips {
				iploc.GetIpInfo(ipAddr)
			}
		}
	})
	fmt.Println(r)
	fmt.Printf("10w次查询: %.1f 毫秒", float64(r.NsPerOp()) / 100000000 * 1000 / 10 * 100000)
}

func main() {
	testIp("0.0.0.0")
	testIp("127.0.0.1")
	testIp("169.254.0.1")
	testIp("192.168.1.1")
	testIp("10.0.0.0")
	testIp("255.255.255.255")
	testIp("112.226.155.1")
	testIp("121.18.72.0")

	testSpeed()
}
