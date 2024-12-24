package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Тест функции main с командой "echo"
func TestMainEcho(t *testing.T) {
	// Перехватываем вывод в STDOUT
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Подготавливаем команду
	input := "echo Hello, World!"
	expectedOutput := "Hello, World!"
	inR, inW, _ := os.Pipe()
	_, err := inW.Write([]byte(input))
	if err != nil {
		return
	}
	err = inW.Close()
	if err != nil {
		return
	}
	oldInput := os.Stdin
	os.Stdin = inR

	// Запускаем тестирование
	go func() {
		defer func(w *os.File) {
			err := w.Close()
			if err != nil {

			}
		}(w)
		main()
	}()

	// Считываем вывод из STDOUT
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}

	// Восстанавливаем STDOUT и STDIN
	os.Stdout = oldOutput
	os.Stdin = oldInput

	// Проверяем результат
	actualOutput := strings.TrimSpace(buf.String())
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод: %s, Фактический вывод: %s", expectedOutput, actualOutput)
	}
}

// Тест функции main с командой "pwd"
func TestMainPwd(t *testing.T) {
	// Перехватываем вывод в STDOUT
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Подготавливаем команду
	expectedOutput, _ := os.Getwd()
	inR, inW, _ := os.Pipe()
	_, err := inW.Write([]byte("pwd"))
	if err != nil {
		return
	}
	err = inW.Close()
	if err != nil {
		return
	}
	oldInput := os.Stdin
	os.Stdin = inR

	// Запускаем тестирование
	go func() {
		defer func(w *os.File) {
			err := w.Close()
			if err != nil {
				return
			}
		}(w)
		main()
	}()

	// Считываем вывод из STDOUT
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}

	// Восстанавливаем STDOUT и STDIN
	os.Stdout = oldOutput
	os.Stdin = oldInput

	// Проверяем результат
	actualOutput := strings.TrimSpace(buf.String())
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод: %s, Фактический вывод: %s", expectedOutput, actualOutput)
	}
}

// Тест функции main с командой "ls"
func TestMainLs(t *testing.T) {
	// Перехватываем вывод в STDOUT
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Подготавливаем команду
	expectedOutput := executeCommand("ls")
	inR, inW, _ := os.Pipe()
	_, err := inW.Write([]byte("ls"))
	if err != nil {
		return
	}
	err = inW.Close()
	if err != nil {
		return
	}
	oldInput := os.Stdin
	os.Stdin = inR

	// Запускаем тестирование
	go func() {
		defer func(w *os.File) {
			err := w.Close()
			if err != nil {
				return
			}
		}(w)
		main()
	}()

	// Считываем вывод из STDOUT
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}

	// Восстанавливаем STDOUT и STDIN
	os.Stdout = oldOutput
	os.Stdin = oldInput

	// Проверяем результат
	actualOutput := strings.TrimSpace(buf.String())
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод: %s, Фактический вывод: %s", expectedOutput, actualOutput)
	}
}

// Функция для выполнения внешних команд и возврата вывода
func executeCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	output, _ := cmd.Output()
	return string(output)
}