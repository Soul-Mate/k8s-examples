package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	ip := os.Getenv("WEBAPP_HELLOWORLD_SERVICE_HOST")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello World!-%d [Client: %s, Host: %s, IP: %s] \n",
			atomic.LoadInt64(&cnt), r.RemoteAddr, hostname, ip)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}
