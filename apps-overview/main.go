// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type data struct {
	ReviewSlugWithDomain string
	EnvironmentSubdomain string
}

const readHeaderTimeout = time.Second * 60

func serveHTML(w http.ResponseWriter, r *http.Request) {
	environmentSubdomain := os.Getenv("ENVIRONMENT_SUBDOMAIN")
	reviewSlugWithDomain := os.Getenv("ENVIRONMENT_SLUG_WITH_DOMAIN")

	templateFilePath := templateForSubdomain(environmentSubdomain)

	w.Header().Set("Permissions-Policy", "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), layout-animations=(), legacy-image-formats=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), oversized-images=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), sync-xhr=(), usb=(), vr=(), wake-lock=(), screen-wake-lock=(), web-share=(), xr-spatial-tracking=()")

	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, &data{
		ReviewSlugWithDomain: reviewSlugWithDomain,
		EnvironmentSubdomain: environmentSubdomain,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", serveHTML)
	log.Println("starting reviewPage")

	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: readHeaderTimeout,
	}

	err := s.ListenAndServe()
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
