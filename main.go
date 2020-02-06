package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}

// Usage xxx.xxx/api?u=https://example.com
func api(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, r.FormValue("u"), r.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Origin", req.URL.Host)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.Copy(w, resp.Body)
}
