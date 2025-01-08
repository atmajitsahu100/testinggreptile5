package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SecretHandler struct {
	SecretKey   string
	AccessToken string
	DatabaseURL string
	PrivateKey  string
}

func NewSecretHandler(key, token, dbURL, privKey string) *SecretHandler {
	return &SecretHandler{
		SecretKey:   key,
		AccessToken: token,
		DatabaseURL: dbURL,
		PrivateKey:  privKey,
	}
}

func (h *SecretHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/get-secret" {
		h.handleSecretRequest(w, r)
		return
	}
	if r.URL.Path == "/get-token" {
		h.handleTokenRequest(w, r)
		return
	}
	http.NotFound(w, r)
}

func (h *SecretHandler) handleSecretRequest(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"secret": h.SecretKey, "db_url": h.DatabaseURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SecretHandler) handleTokenRequest(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"access_token": h.AccessToken, "private_key": h.PrivateKey}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main1() {
	secretKey := "adnsjfsdfffdsssfsfdsfdfdfsaddada12345"
	accessToken := "ya29.a0ARrdaM8wSkX9F6hje43gd5K7V8g7Bd3bbdkv872jk"
	databaseURL := "jdbc:mysql://username:password@127.0.0.1:3306/testdb"
	privateKey := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCz2p6D"

	mux := http.NewServeMux()
	secretHandler := NewSecretHandler(secretKey, accessToken, databaseURL, privateKey)
	mux.Handle("/get-secret", secretHandler)
	mux.Handle("/get-token", secretHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
