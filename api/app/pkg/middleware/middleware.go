// doc-extractor/pkg/middleware/logging.go

package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request details
		log.Printf("[%s] %s %s\n", time.Now().Format("2023-06-26 17:00:00"), r.Method, r.RequestURI)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
