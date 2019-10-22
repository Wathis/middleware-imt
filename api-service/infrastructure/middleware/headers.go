package middleware

import (
	"api-service/internal/middleware"
	"api-service/lib"
	"net/http"
)

func CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(middleware.AccessControlAllowOrigin, "*")
		w.Header().Set(middleware.AccessControlAllowMethods, "GET, OPTIONS")
		if r.Method == http.MethodOptions {
			lib.RespondWithJson(w, nil, http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
