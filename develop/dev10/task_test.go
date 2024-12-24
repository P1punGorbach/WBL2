package main

import (
	"io"
	"log"
	"net"
	"os/exec"
	"testing"
	"time"
)

// todo: fix tests
func TestTelnetClient(t *testing.T) {
	// Запускаем тестовый TCP-сервер на случайном порту
	server, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("Failed to start test server:", err)
	}
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(server)

	// Запускаем telnet-клиент в отдельном процессе
	go runTelnetClient(server.Addr().String())

	// Ждем подключения клиента к тестовому серверу
	conn, err := server.Accept()
	if err != nil {
		t.Fatal("Failed to accept connection:", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(conn)

	// Отправляем данные в сокет клиента
	expected := "Hello, world!\n"
	_, err = conn.Write([]byte(expected))
	if err != nil {
		t.Fatalf("Send error: %s", err)
	}

	// Читаем данные, которые клиент должен принять
	actual, err := io.ReadAll(conn)
	if err != nil {
		t.Fatal("Failed to read from client:", err)
	}

	// Проверяем, что принятые данные совпадают с ожидаемыми
	if string(actual) != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, string(actual))
	}

	// Отправляем сигнал завершения клиентской программы
	cmd := exec.Command("pkill", "-SIGINT", "go-telnet")
	err = cmd.Run()
	if err != nil {
		t.Fatal("Failed to send interrupt signal:", err)
	}

	// Ждем пока клиент закроет соединение
	time.Sleep(100 * time.Millisecond)
}

// Функция для запуска telnet-клиента в отдельном процессе
func runTelnetClient(addr string) {
	// go build -o task.exe
	cmd := exec.Command("./task.exe", addr) // Используйте путь к исполняемому файлу вашего клиента
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}