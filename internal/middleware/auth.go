package middleware

import (
	"context"
	"net/http"
	"time"

	"fitness-tracker/internal/login"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve cookie
		cookie, err := r.Cookie("authToken")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate JWT
		userID, err := login.ValidateJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add userID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve current token from cookie
	cookie, err := r.Cookie("authToken")
	if err != nil || cookie.Value == "" {
		http.Error(w, "No valid token", http.StatusUnauthorized)
		return
	}

	// Validate & Generate new token
	userID, err := login.ValidateJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	newToken, err := login.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Update cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    newToken,
		Expires:  time.Now().Add(36 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	w.Write([]byte("Token refreshed"))
}