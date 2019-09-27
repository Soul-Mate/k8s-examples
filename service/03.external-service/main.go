package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	ip := os.Getenv("WEBAPP_HELLOWORLD_SERVICE_HOST")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://api-service:31000/external")
		if err != nil {
			fmt.Fprintf(w, "request api-service error: %s", err.Error())
			return
		}

		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")

		// read api-service response
		data, _ := ioutil.ReadAll(resp.Body)
		m := make(map[string]interface{})
		json.Unmarshal(data, &m)

		m["data"] = fmt.Sprintf("webapp-external-service call %s", m["data"])
		m["ip"] = ip
		m["hostname"] = hostname

		if resp, err := http.Get("http://markhub-external"); err != nil {
			fmt.Fprintf(w, "request markhub-external error: %s", err.Error())
			return
		} else {
			data, _ = ioutil.ReadAll(resp.Body)
			m["www.markhub.cn-webpage"] = string(data)
		}

		bs, _ := json.Marshal(m)

		w.Write(bs)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
