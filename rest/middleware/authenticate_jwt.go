package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/turjoc120/ecom/config"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		accessTokenArr := strings.Split(accessToken, ".")
		if len(accessTokenArr) != 3 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := accessTokenArr[0]
		jwtPayload := accessTokenArr[1]
		jwtSignature := accessTokenArr[2]

		message := jwtHeader + "." + jwtPayload

		byteArrMessage := []byte(message)
		byteArrSecret := []byte(config.GetConfig().JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if newSignature != jwtSignature {
			http.Error(w, "tui hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
