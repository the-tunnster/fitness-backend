package middleware

import (
	"context"
	"net/http"

	"fitness-tracker/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ctxKey string

const UserIDKey ctxKey = "user_id"

// RequireUser ensures requests include a valid user_id (header or query) and
// attaches it to the request context. If missing/invalid, responds 401.
func RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			userID = r.URL.Query().Get("user_id")
		}
		if userID == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized: missing user_id")
			return
		}
		if _, err := primitive.ObjectIDFromHex(userID); err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized: invalid user_id")
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
