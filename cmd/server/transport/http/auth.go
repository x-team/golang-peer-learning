package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWtAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(accessToken string) bool {
	var signingKey = []byte("THE_KEY")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid Auth Token")
		}
		return signingKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
