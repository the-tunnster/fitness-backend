package middleware

import (
	"net/http"
	"strings"

	"fitness-tracker/internal/utils"
)

// AllowMethods enforces a set of allowed HTTP methods for a handler.
// If the request method is not allowed, it responds with 405 and sets
// the Allow header to the permitted methods.
// If GET is allowed, HEAD is implicitly allowed.
func AllowMethods(allowed []string, next http.Handler) http.Handler {
	// Normalize and ensure HEAD if GET is present
	set := make(map[string]struct{}, len(allowed)+1)
	for _, m := range allowed {
		mm := strings.ToUpper(m)
		set[mm] = struct{}{}
		if mm == http.MethodGet {
			set[http.MethodHead] = struct{}{}
		}
	}

	// Precompute Allow header value
	allowList := make([]string, 0, len(set))
	for m := range set {
		allowList = append(allowList, m)
	}
	allowHeader := strings.Join(allowList, ", ")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := set[r.Method]; !ok {
			w.Header().Set("Allow", allowHeader)
			utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
			return
		}
		next.ServeHTTP(w, r)
	})
}
