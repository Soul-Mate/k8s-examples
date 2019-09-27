package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/external", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]interface{}{
			"code": 200,
			"data": "api-service",
		}

		bs, _ := json.Marshal(data)
		w.Write(bs)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:31000", nil))
}
