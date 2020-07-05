package middleware

import (
	"log"
	"net/http"
)

func ActivityLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		log.Printf("Accessing path %v with application %v\n", r.RequestURI, userAgent)

		// Untuk melanjutkan ke middlware selanjutnya
		next.ServeHTTP(w, r)
	})
}
