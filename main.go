//tasks

package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//this is executed on the way down to the handeler
		log.Println("Executing middleware")
		log.Println(r.URL.Path)
		
		if r.URL.Path != "/" && r.URL.Path !="/random" && r.URL.Path != "/greeting" {
			//return and dont go any further
			fmt.Println("invalid link")
			return
		}

		next.ServeHTTP(w, r)
		//this is executed on the way up
		log.Println("Executing middlewareB again")
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	info := "Using template"

	ts, _ := template.ParseFiles("public/index.html.tmpl")

	ts.Execute(w,info)

	log.Println("working")
	log.Println(r.RemoteAddr)

}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	//get random quote using API
	//then save to a file
	//used in index.html

	///SystemsProgrammingTest-AldricRivero/
	http.ServeFile(w, r, "public/index.html")

	log.Println("working")
	log.Println(r.RemoteAddr)

}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	
	http.ServeFile(w, r, "public/index.html")

	log.Println("working")
	log.Println(r.RemoteAddr)

}

func main() {
	// serve static files from the "static" directory
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", middleware(http.HandlerFunc(homeHandler)))
	http.Handle("/random", middleware(http.HandlerFunc(randomHandler)))
	http.Handle("/greeting", middleware(http.HandlerFunc(greetingHandler)))

	log.Fatal(http.ListenAndServe("192.168.18.24:80", nil))
}
