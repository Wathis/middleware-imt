package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
)

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		log.Printf("[%s] %s (%v)", r.URL, r.Method, r.Body)
		next.ServeHTTP(rec, r)
		log.Printf("[%s] %d (%v)", r.URL, rec.Code, rec.Body)
		w.WriteHeader(rec.Code)
		_, _ = w.Write(rec.Body.Bytes())
	})
}
