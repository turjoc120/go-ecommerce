package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byeArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byeArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	headerB64 := base64UrlEncode(byeArrHeader)
	payloadB64 := base64UrlEncode(byeArrPayload)
	message := headerB64 + "." + payloadB64

	byteArrSecrect := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecrect)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)
	jwt := message + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
