package main

import (
	"encoding/json"
	"github.com/ascottqqq/rfc7539"
	"net/http"
)

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "GET Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		decoder := json.NewDecoder(r.Body)
		var t rfc7539.ChaCha20
		err := decoder.Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			ciphertext := rfc7539.Encrypt(&t)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			a, err := json.Marshal(ciphertext)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.Write(a)
			}
		}
	}
	defer r.Body.Close()
}

func main() {
	http.HandleFunc("/encrypt/", EncryptHandler)
	// ChaCha20 uses same algorithm to encrypt and decrypt
	http.HandleFunc("/decrypt/", EncryptHandler)
	http.ListenAndServe(":8080", nil)
}
