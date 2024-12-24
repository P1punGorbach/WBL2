package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetFilename(t *testing.T) {
	testCases := []struct {
		url      string
		expected string
	}{
		{"https://www.example.com", "www.example.com"},
		{"https://www.example.com/index.html", "index.html"},
		{"https://www.example.com/dir/file.html", "file.html"},
	}

	for _, testCase := range testCases {
		actual := getFilename(testCase.url)
		if actual != testCase.expected {
			t.Errorf("Expected %s, but got %s for URL %s", testCase.expected, actual, testCase.url)
		}
	}
}

func TestDownloadFile(t *testing.T) {
	testCases := []struct {
		url      string
		expected string
	}{
		{"https://www.example.com", "www.example.com"},
		{"https://www.example.com/index.html", "index.html"},
		{"https://www.example.com/dir/file.html", "file.html"},
	}

	for _, testCase := range testCases {
		actual := getFilename(testCase.url)
		if actual != testCase.expected {
			t.Errorf("Expected %s, but got %s for URL %s", testCase.expected, actual, testCase.url)
		}
	}
}
func TestMainFunction(t *testing.T) {
	content := "Test file content"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(content))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	// Create a temporary file
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(file.Name())

	// Call the main function
	os.Args = []string{"cmd", server.URL}
	main()

	// Read the file
	fileContent, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	// Verify the file content
	if string(fileContent) != content {
		t.Errorf("Expected %q, but got %q for file content", content, string(fileContent))
	}
}