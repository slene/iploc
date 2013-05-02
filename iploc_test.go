package iploc

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"
	"time"
)

var seedrand = rand.New(rand.NewSource(time.Now().UnixNano()))

func getIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", seedrand.Intn(255), seedrand.Intn(255), seedrand.Intn(255), seedrand.Intn(255))
}

func init() {
	iplocFilePath, _ := filepath.Abs("iploc.dat")
	IpLocInit(iplocFilePath, true)
}

func Test_GetIpInfo(t *testing.T) {
	GetIpInfo(getIP())
}

func Benchmark_GetIpInfo(b *testing.B) {
	ipAddr := getIP()
	for i := 0; i < b.N; i++ {
		GetIpInfo(ipAddr)
	}
}
