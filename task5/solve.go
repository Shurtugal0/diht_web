package main

import (
	"encoding/json"
	"net/http"
)

var links map[string]string

type Request struct {
	Url string `json:"url"`
}

func shortener(link string) string {
	shortlink := ""
	len := len(links) + 1
	for ; len > 0; len /= 26 {
		shortlink += string('a' + len % 26)
	}
	links[shortlink] = link
	return shortlink
}

func shortenerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if req.Url == "" || err != nil {
			http.Error(w, "", 400)
			return
		}
		ans := make(map[string]string)
		ans["key"] = shortener(req.Url)
		out_ans, _ := json.Marshal(ans)
		w.Write(out_ans)
	} else if r.Method == "GET" {
		link, ok := links[r.RequestURI[1:]]
		if !ok {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, link, 301)
	}
}

func main() {
	links = make(map[string]string)
	http.HandleFunc("/", shortenerHandler)
	http.ListenAndServe(":8082", nil)
}