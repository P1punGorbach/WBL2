package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	// Создание временного файла и запись туда тестовых данных
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(file.Name())

	text := "Hello\nworld\n"
	if _, err := file.WriteString(text); err != nil {
		t.Fatal(err)
	}

	// Предотвращение чтения файла из стандартного ввода
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = nil

	// Чтение вывода программы в буфер
	buf := new(bytes.Buffer)
	//todo: fix tests
	os.Stdout = buf

	// Запуск программы с различными опциями
	_ = []string{"-n", "-F", "-A", "1", "-B", "1", "-C", "1", "-c", "-i", "-v", "world", file.Name()}
	main()

	actual := buf.String()

	// Проверка основного вывода программы
	expected := "2: world\nworld\n"
	if actual != expected {
		t.Errorf("Expected output: %s, got: %s", expected, actual)
	}

	// Проверка количества совпадений
	expectedCount := "Matching lines count: 2\n"
	if !strings.Contains(actual, expectedCount) {
		t.Errorf("Expected count output: %s not found in actual output: %s", expectedCount, actual)
	}
}