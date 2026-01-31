package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecoommerce/util"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecrect := []byte(m.cnf.JwtSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecrect)
		h.Write(byteArrMessage)

		signature := h.Sum(nil)
		oldSignature := base64UrlEncode(signature)

		if jwtSignature != oldSignature {
			util.SendData(w, http.StatusUnauthorized, "jah tui hacker")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
