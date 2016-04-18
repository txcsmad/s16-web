package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// A simple struct with an embedded mutex
type counter struct {
	sync.Mutex
	value int
}

func (c *counter) reset() {
	c.Lock()
	defer c.Unlock()
	c.value = 0
}

var (
	c       counter
	htmlStr = `
<html>
	<head>
		<title>Hits Counter</title>
	</head>
	<body>
		Woah, you're visitor {{ .count }}
	</body>
</html>`
	htmlTmpl *template.Template
)

func main() {
	htmlTmpl = template.Must(template.New("index").Parse(htmlStr))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/reset", resetHandler)

	fmt.Println("serving on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	c.Lock()
	defer c.Unlock()
	c.value += 1

	htmlTmpl.Execute(w, map[string]int{
		"count": c.value,
	})
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	c.Lock()
	defer c.Unlock()
	fmt.Fprintf(w, fmt.Sprintf("%d", c.value))
}

func resetHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	c.reset()
	w.WriteHeader(http.StatusOK)
}
