package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestGetTime(t *testing.T) {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		t.Fatalf("Ошибка при получении времени: %v", err)
	}

	if ntpTime.IsZero() {
		t.Error("Получено пустое время")
	}

	localTime := time.Now()
	if localTime.Sub(ntpTime) > time.Second || localTime.Sub(ntpTime) < -time.Second {
		fmt.Println("Отклонение между локальным и временем ntp слишком велико")
	}
}

func TestMain(m *testing.M) {
	_, err := ntp.Query("pool.ntp.org")
	if err != nil {
		log.Fatalf("Ошибка при проверке доступности NTP сервера: %v", err)
	}

	fmt.Println("Запуск тестов")
	exitCode := m.Run()
	fmt.Println("Тесты завершены")

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
