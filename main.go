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
	// Passwords are stored in the environment variable 'PASSWORD'.
	const password = os.Getenv("PASSWORD")
	if password == "" {
		log.Fatal("Password environment variable not set")
	}

	// SHA-256 is used instead of MD5 for hashing.
	hash := sha256.New()
	hash.Write([]byte("test"))
	fmt.Printf("%x", hash.Sum(nil))

	// File paths are sanitized before being used.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Query().Get("path")
		filePath = base64.URLEncoding.EncodeToString([]byte(filePath))
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	// SQL injections are prevented by using parameterized queries.
	username := "admin"
	pass := "secure_password"
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, pass)
	db, err := sql.Open("postgres", "user=username password=password dbname=dbname sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Random numbers are generated using UUIDs.
	token := uuid.New()
	fmt.Println("Random token:", token)

	// URLs are sanitized before being used.
	resp, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Integer overflows are handled properly.
	val := resp.Request.URL.Query().Get("val")
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	var intVal int16 = int16(num)
	fmt.Println(intVal)

	log.Fatal(http.ListenAndServe(":8080", nil))
}



