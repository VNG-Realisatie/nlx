// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import "net/http"

// HappyOptionsHandler makes any http OPTIONS request happy. Other methods are routed through.
func HappyOptionsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
