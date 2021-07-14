package loggermiddleware

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		fmt.Println("Processing request...")
		next.ServeHTTP(w, r)
	})
}
