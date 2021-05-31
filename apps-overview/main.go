// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

type data struct {
	ReviewSlugWithDomain string
	EnvironmentSubdomain string
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	environmentSubdomain := os.Getenv("ENVIRONMENT_SUBDOMAIN")

	templateFilePath := templateForSubdomain(environmentSubdomain)

	w.Header().Set("Permissions-Policy", "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), layout-animations=(), legacy-image-formats=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), oversized-images=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), sync-xhr=(), usb=(), vr=(), wake-lock=(), screen-wake-lock=(), web-share=(), xr-spatial-tracking=()")

	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, &data{
		ReviewSlugWithDomain: os.Getenv("ENVIRONMENT_SLUG_WITH_DOMAIN"),
		EnvironmentSubdomain: environmentSubdomain,
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

func templateForSubdomain(subdomain string) string {
	var templateFilePath string

	switch subdomain {
	case "review":
		templateFilePath = path.Join("templates", "sites-review.html")
	case "acc":
		templateFilePath = path.Join("templates", "sites-acc.html")
	default:
		templateFilePath = path.Join("templates", "sites.html")
	}

	return templateFilePath
}
