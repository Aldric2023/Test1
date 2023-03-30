//tasks

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	filename := "index.html"
	body, _ := os.ReadFile(filename)
	fmt.Fprintf(w, string(body), r.URL.Path[1:])

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
