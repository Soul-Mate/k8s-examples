package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The %s (%s) Say: Hello World %s!\n", GetHostname(), GetInternetIP(), r.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}

var (
	IP           string
	ipLock       sync.RWMutex
	Hostname     string
	hostnameLock sync.RWMutex
)

// GetInternetIP 用于自动查找本机IP地址
func GetInternetIP() string {
	ipLock.RLock()
	if IP != "" {
		ipLock.RUnlock()
		return IP

	}

	ipLock.RUnlock()
	// 查找本机IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				if ip4[0] == 10 {
					// 赋值新的IP
					ipLock.Lock()
					IP = ip4.String()
					ipLock.Unlock()
				}
			}
		}
	}

	return IP
}

// GetHostname 用于自动获取本机Hostname信息
func GetHostname() string {
	hostnameLock.RLock()
	if Hostname != "" {
		hostnameLock.RUnlock()
		return Hostname
	}

	hostnameLock.RUnlock()
	// 查找本机hostname
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	hostnameLock.Lock()
	Hostname = hostname
	hostnameLock.Unlock()

	return Hostname
}
