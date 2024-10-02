package middlewares

import (
	"knightstar/pkg/util"
	"net/http"
	"strings"
)

// AuthMiddleware middleware structure
type AuthMiddleware struct {
	avoidAuthPaths map[string][]string
}

// NewAuthMiddleware creates a new instance of the AuthMiddleware
func NewAuthMiddleware(avoidAuthPaths map[string][]string) *AuthMiddleware {
	return &AuthMiddleware{
		avoidAuthPaths: avoidAuthPaths,
	}
}

// Helper function to check if a string exists in a slice
func (am *AuthMiddleware) contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// AuthMiddleware validates the JWT token and user ID before processing the request.
// It checks the Authorization header for the token and x-key header for the user ID.
func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow specific paths without authentication
		// Check if the path is allowed for the request method
		// TODO: Add support for wildcards
		allowedMethods, pathExists := am.avoidAuthPaths[r.URL.Path]
		if pathExists {
			if am.contains(allowedMethods, r.Method) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Extract token from Authorization header or x-access-token header
		authorizationHeader := r.Header.Get("Authorization")

		var tokenString string
		if authorizationHeader != "" {
			parts := strings.Split(authorizationHeader, " ")
			if len(parts) == 2 {
				tokenString = parts[1]
			}
		} else {
			tokenString = r.Header.Get("x-access-token")
		}

		// Extract user ID from x-key header
		userId := r.Header.Get("x-key")

		// Validate the token and user ID
		if tokenString != "" && userId != "" {
			result, err := util.ValidateToken(tokenString, userId)
			if err != nil {
				util.WriteJSONResponse(w, http.StatusUnauthorized, util.JSON{"message": err.Error()})
				return
			}
			if result {
				next.ServeHTTP(w, r)
			} else {
				util.WriteJSONResponse(w, http.StatusUnauthorized, util.JSON{"message": "Invalid Token"})
			}
		} else {
			util.WriteJSONResponse(w, http.StatusBadRequest, util.JSON{"message": "Token or Key is missing"})
		}
	})
}
