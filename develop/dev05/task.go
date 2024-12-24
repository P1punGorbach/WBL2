package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Парсинг аргументов командной строки
	before := flag.Int("B", 0, "Print +N lines before each match")
	after := flag.Int("A", 0, "Print +N lines after each match")
	context := flag.Int("C", 0, "Print ±N lines around each match")
	count := flag.Bool("c", false, "Print only the count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Match exact string, not pattern")
	withLineNum := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	// Чтение паттерна поиска
	pattern := flag.Arg(0)

	// Открытие файла для чтения или использование стандартного ввода
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Создание счетчика количества совпадений
	matchCount := 0

	// Чтение файла построчно
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()

		// Применение опций фильтрации
		matches := false

		if *fixed {
			if strings.Contains(line, pattern) {
				matches = true
			}
		} else {
			if *ignoreCase {
				line = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
			}

			if strings.Contains(line, pattern) {
				matches = true
			}
		}

		if (*invert && !matches) || (!*invert && matches) {
			// Применение опций вывода
			if *withLineNum {
				fmt.Print(lineNum, ": ")
			}

			fmt.Println(line)

			if *count {
				matchCount++
				continue
			}

			if *before > 0 {
				for i := lineNum - *before; i < lineNum; i++ {
					if i > 0 {
						scanner.Scan()
						fmt.Println(scanner.Text())
					}
				}
			}

			if *after > 0 {
				for i := lineNum + 1; i <= lineNum+*after; i++ {
					scanner.Scan()
					fmt.Println(scanner.Text())
				}
			}

			if *context > 0 {
				for i := lineNum - *context; i < lineNum+*context; i++ {
					if i > 0 {
						scanner.Scan()
						fmt.Println(scanner.Text())
					}
				}
			}
		}

		lineNum++
	}

	if *count {
		fmt.Println("Matching lines count:", matchCount)
	}
}