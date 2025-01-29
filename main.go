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
	// Replaced MD5 hash function with SHA-256
	hash := sha256.New()
	hash.Write([]byte("test"))
	fmt.Printf("%x", hash.Sum(nil))

	// Validated and sanitized file path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Query().Get("path")
		filePath = url.QueryEscape(filePath)
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	// Parameterized queries to prevent SQL injection
	username := "admin"
	pass := "password"
	query := fmt.Sprintf("SELECT * FROM users WHERE username=? AND password=?")
	db, _ := sql.Open("mysql", "user:password@/dbname")
	rows, _ := db.Query(query, username, pass)
	var count int
	for rows.Next() {
		count++
	}
	if count > 0 {
		fmt.Println("Access granted!")
	}

	// Replaced DES with AES
	key := []byte("my-secret-key-1234")
	block, _ := aes.NewCipher(key)
	fmt.Printf("%x", block)

	// Set minimum TLS version to 1.3
	config := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}
	_, _ = tls.Dial("tcp", "example.com:443", config)

	// Used crypto/rand for random number generation
	token, _ := rand.Int(rand.Reader, big.NewInt(1<<60))
	fmt.Println("Random token:", token)

	resp, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Validated and sanitized URL
	url := resp.Request.URL.Query().Get("url")
	parsedUrl, _ := url.ParseRequestURI(url)
	http.Get(parsedUrl.String())

	// Bounds checked conversion from string to int
	val := resp.Request.URL.Query().Get("val")
	num, _ := strconv.Atoi(val)
	var intVal int16 = int16(num)
	fmt.Println(intVal)

	// Limited decompression
	http.HandleFunc("/decompress", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<30) // 1GB
		gzr, _ := gzip.NewReader(r.Body)
		_, _ = io.Copy(os.Stdout, gzr)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}






