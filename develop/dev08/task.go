package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Выводим приглашение на экран
		fmt.Print("> ")

		// Читаем команду из ввода пользователя
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения команды:", err)
			continue
		}

		// Удаляем символы новой строки из команды
		command = strings.TrimSuffix(command, "\n")

		// Обработка введенной команды
		switch {
		case command == "quit":
			// Завершаем выполнение программы
			return
		case strings.HasPrefix(command, "cd"):
			// Изменяем текущую директорию
			args := strings.TrimSpace(strings.TrimPrefix(command, "cd"))
			err := os.Chdir(args)
			if err != nil {
				fmt.Println("Ошибка смены директории:", err)
			}
		case command == "pwd":
			// Выводим текущую директорию
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка получения текущей директории:", err)
			}
			fmt.Println(wd)
		case strings.HasPrefix(command, "echo"):
			// Выводим аргументы на экран
			args := strings.TrimSpace(strings.TrimPrefix(command, "echo"))
			fmt.Println(args)
		case strings.HasPrefix(command, "kill"):
			// "Убиваем" процесс
			args := strings.TrimSpace(strings.TrimPrefix(command, "kill"))
			cmd := exec.Command("kill", args)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды kill:", err)
			}
		case command == "ps":
			// Выводим информацию о текущих процессах
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды ps:", err)
			}
		default:
			// Выполняем команду через shell
			cmd := exec.Command("bash", "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды:", err)
			}
		}
	}
}