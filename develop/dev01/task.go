package main

import (
	"fmt"
	"io"
	"os"

	ntp "github.com/beevik/ntp"
)

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

func PrintCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_ = fmt.Errorf(err.Error(), stderr)
		os.Exit(1)
		return
	}

	fmt.Fprintln(stdout, time)
}
/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
*/
