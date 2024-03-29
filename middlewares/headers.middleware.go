package middlewares

import (
	"net/http"
)

func SetHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Created-By", "Ophelia Pavlova")
		next.ServeHTTP(res, req)
	})
}