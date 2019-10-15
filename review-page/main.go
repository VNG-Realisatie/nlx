package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

type data struct {
	EnvironmentSlugWithDomain string
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "sites.html")

	t, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, &data{
		EnvironmentSlugWithDomain: os.Getenv("ENVIRONMENT_SLUG_WITH_DOMAIN"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", serveHTML)
	log.Println("starting reviewPage")

	err := http.ListenAndServe(":8080", nil)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
