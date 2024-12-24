package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %v", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Close error: %s", err)
		}
	}(resp.Body)

	filename := getFilename(url)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v", filename, err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("File error: %s", err)
		}
	}(file)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v", filename, err)
		return
	}

	fmt.Printf("Downloaded %s\n", filename)
}

func getFilename(url string) string {
	// Find the last "/" in the URL
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		// No slash, just return the full URL as the filename
		return url
	}
	// Return the substring after the last slash as the filename
	return url[lastSlashIndex+1:]
}