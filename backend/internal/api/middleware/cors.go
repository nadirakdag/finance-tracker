package middleware

import (
	"github.com/nadirakdag/finance-tracker/internal/config"
	"net/http"
	"strings"
)

func Cors(cfg *config.CorsConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Check if origin is allowed
			if len(cfg.AllowedOrigins) > 0 {
				allowed := false
				for _, allowedOrigin := range cfg.AllowedOrigins {
					if allowedOrigin == "*" || allowedOrigin == origin {
						allowed = true
						break
					}
				}
				if allowed {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				}
			}

			// Set allowed methods
			w.Header().Set("Access-Control-Allow-Methods",
				strings.Join(cfg.AllowedMethods, ", "))

			// Set allowed headers
			w.Header().Set("Access-Control-Allow-Headers",
				strings.Join(cfg.AllowedHeaders, ", "))

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
