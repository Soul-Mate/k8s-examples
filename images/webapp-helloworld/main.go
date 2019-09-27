package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := os.Getenv("HOSTNAME")
		podName := os.Getenv("MY_POD_NAME")
		podNamespace := os.Getenv("MY_POD_NAMESPACE")
		podIP := os.Getenv("MY_POD_IP")
		bs, _ := json.Marshal(&map[string]interface{}{
			"pod": map[string]string{
				"hostname":      hostname,
				"pod_name":      podName,
				"pod_namespace": podNamespace,
				"pod_ip":        podIP,
			},
			"code": 200,
			"data": "Hello World",
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(bs)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
