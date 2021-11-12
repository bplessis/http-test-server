package main

import (
	"fmt"
	"log"
	"net/http"
)

type testHttpHandler struct {
}

func (h *testHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Server", "Go-http-test-server/1.0")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<!DOCTYPE html5>\n<html>\n<head><title>Requested page: %f</title></head>\n<body>\n", r.URL)
    fmt.Fprintf(w, "<article>\n")
	fmt.Fprintf(w, "<section><h1>Request</h1><p><table id=\"request\">")
	fmt.Fprintf(w, "<tr><th>Method</th><td>%f</td>", r.Method)
	fmt.Fprintf(w, "<tr><th>Host</th><td>%f</td>", r.Host)
	fmt.Fprintf(w, "<tr><th>URL</th><td>%f</td>", r.URL)
	fmt.Fprintf(w, "<tr><th>RemoteAddr</th><td>%f</td>", r.RemoteAddr)
	fmt.Fprintf(w, "<tr><th>RequestURI</th><td>%f</td>", r.RequestURI)
	fmt.Fprintf(w, "</table></p></section>\n")
	
    fmt.Fprintf(w, "<section><h1>Headers</h1><p><table id=\"headers\"><tr><th>Header</th><th>Value</th></tr>")
    for name, headers := range r.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "<tr><td>%v</td><td>%v</td></tr>", name, h)
        }
    }
	fmt.Fprintf(w, "</table></p></section>\n</article>\n</body>\n</html>\n")
    
}

func main() {
	http.Handle("/", new(testHttpHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}