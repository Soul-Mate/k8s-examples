package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

func main() {
	http.HandleFunc("/arith", func(w http.ResponseWriter, r *http.Request) {
		addr := os.Getenv("WEBAPP_ARITH_SERVICE_SERVICE_HOST")
		port := os.Getenv("WEBAPP_ARITH_SERVICE_SERVICE_PORT")
		if addr == "" || port == "" {
			log.Fatal("serivce found error")
		}

		client, err := rpc.DialHTTP("tcp", addr+":"+port)
		if err != nil {
			log.Fatal("dialing:", err)
		}

		var (
			add, mul, sub int
			div, sum      float64
		)

		wg := sync.WaitGroup{}

		wg.Add(4)
		// add
		go func() {
			defer wg.Done()
			args := &Args{7, 8}
			reply := &ArithResult{}
			if err := client.Call("Arith.Add", args, &reply); err != nil {
				log.Fatal("arith add error:", err)
			}
			add = reply.Value
		}()

		// sub
		go func() {
			defer wg.Done()
			args := &Args{7, 8}
			reply := &ArithResult{}
			if err := client.Call("Arith.Sub", args, &reply); err != nil {
				log.Fatal("arith add error:", err)
			}
			sub = reply.Value
		}()

		// mul
		go func() {
			defer wg.Done()
			args := &Args{7, 8}
			reply := &ArithResult{}
			if err := client.Call("Arith.Mul", args, &reply); err != nil {
				log.Fatal("arith add error:", err)
			}
			mul = reply.Value
		}()

		// div
		go func() {
			defer wg.Done()
			args := &Args{7, 8}
			reply := &ArithResult{}
			if err := client.Call("Arith.Div", args, &reply); err != nil {
				log.Fatal("arith add error:", err)
			}
			div = reply.Quotient
		}()

		wg.Wait()

		sum = float64(add+sub+mul) + div

		fmt.Fprintf(w, "Client: %s Request (Host: %s, IP: %s) Arith Get Result: %f\n",
			r.RemoteAddr, GetHostname(), GetInternetIP(), sum)
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
