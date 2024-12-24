package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)

// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

func Unpack(str string) (string, error) {
	if len(str) == 0 { // Если ввели пустую строку - сразу возвращаем
		return str, nil
	}

	input := []rune(str)

	if unicode.IsDigit(input[0]) { // Если первый символ цифра - ошибка
		return "", errors.New("некорректная строка")
	}

	result := strings.Builder{} // Для наиболее эффективной конкатенации

	for i := 0; i < len(input)-1; i++ {
		switch {
		case string(input[i]) == `\`:
			escape(&result, input, &i)

		default:
			if len(input) >= i+1 && unicode.IsDigit(input[i+1]) {
				repeat, err := strconv.Atoi(string(input[i+1]))
				if err != nil {
					return "", errors.New("ошибка при переводе символа в цифру")
				}
				write(&result, input[i], &i, repeat)

			} else if len(input) >= i {
				write(&result, input[i], &i, 1)

			} else {
				return result.String(), nil
			}
		}
	}

	if !unicode.IsDigit(input[len(input)-1]) || (string(input[len(input)-1]) == `\`) {
		result.WriteRune(input[len(input)-1])

	}

	return result.String(), nil
}

func escape(builder *strings.Builder, input []rune, position *int) {
	if *position+2 < len(input) && unicode.IsDigit(input[*position+2]) {
		r, _ := strconv.Atoi(string(input[*position+2]))
		*position++
		write(builder, input[*position], position, r)

	} else {
		*position++
		write(builder, input[*position], position, 1)
	}
}

func write(builder *strings.Builder, symbol rune, position *int, repetition int) {
	for i := 0; i < repetition; i++ {
		builder.WriteRune(symbol)

	}

	if repetition > 1 {
		*position++
	}
}

func main() {
	Unpack("abcd")
}
