package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 3000, "the port for webserver")
	flag.Parse()

	tmpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handle(tmpl))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handle(tmpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &viewData{
			Time: time.Now(),
		}

		err := tmpl.Execute(w, data)

		if err != nil {
			http.Error(w, "Failed on process the template", http.StatusInternalServerError)
			log.Println(err)
		}
	})
}

type viewData struct {
	Time time.Time
}
