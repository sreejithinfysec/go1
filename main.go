package main

import (
	"compress/gzip"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// Gosec G501: Blacklisted import crypto/md5
	hash := sha256.New() // Use SHA-256 instead of MD5
	hash.Write([]byte("test"))
	fmt.Printf("%x", hash.Sum(nil))

	// Gosec G304: File path provided as taint input
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Query().Get("path")
		filePath = url.QueryEscape(filePath) // Sanitize the file path
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}




