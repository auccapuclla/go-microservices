package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(`{"proto: "%s", "method": "%s", "path": "%s", "duration": "%s"}`,
			r.Proto, r.Method, r.URL.Path, time.Since(t))
	}
}
