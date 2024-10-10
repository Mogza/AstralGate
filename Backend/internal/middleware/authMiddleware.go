package middleware

import (
	"context"
	"github.com/Mogza/AstralGate/internal/utils"
	"net/http"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenString, err := utils.ExtractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userId, err := utils.GetUserIdFromToken(token)
		if err != nil {
			http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
			return
		}

		newCtx := context.WithValue(ctx, "user_id", userId)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenString, err := utils.ExtractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userRole, err := utils.GetRoleFromToken(token)
		if err != nil {
			http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
			return
		}

		if userRole != "admin" {
			http.Error(w, "Not an admin", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
