package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	testFolder := "countdown"

	files, err := os.ReadDir(testFolder)
	if err != nil {
		t.Fatalf("Ошибка при чтении папки с тестами: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return extractNumber(files[i].Name()) < extractNumber(files[j].Name())
	})

	for _, file := range files {
		if file.IsDir() || strings.HasSuffix(file.Name(), ".a") {
			continue
		}

		inputFile := file.Name()
		expectedOutputFile := inputFile + ".a"

		inputData, err := os.ReadFile(testFolder + "/" + inputFile)
		if err != nil {
			t.Fatalf("Ошибка при чтении файла %s: %v", inputFile, err)
		}

		expectedOutputData, err := os.ReadFile(testFolder + "/" + expectedOutputFile)
		if err != nil {
			t.Fatalf("Ошибка при чтении файла %s: %v", expectedOutputFile, err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		cmd.Stdin = bytes.NewReader(inputData)

		startTime := time.Now()
		output, err := cmd.Output()
		duration := time.Since(startTime)

		if ctx.Err() == context.DeadlineExceeded {
			t.Fatalf("Тест %s превысил лимит времени (3000 мс). Выполнение заняло %d мс", inputFile, duration.Milliseconds())
		}

		if err != nil {
			t.Fatalf("Ошибка при выполнении main.go с входным файлом %s: %v", inputFile, err)
		}

		// Разбиваем входные данные, ожидаемый и фактический вывод на строки
		inputLines := strings.Split(strings.TrimSpace(string(inputData)), "\n")
		expectedLines := strings.Split(strings.TrimSpace(string(expectedOutputData)), "\n")
		actualLines := strings.Split(strings.TrimSpace(string(output)), "\n")

		// Проверяем количество строк
		if len(actualLines) != len(expectedLines) {
			t.Fatalf("Несовпадение количества строк в тесте %s:\nОжидалось: %d строк, Получено: %d строк",
				inputFile, len(expectedLines), len(actualLines))
		}

		// Проверяем каждую строку
		for i := 0; i < len(expectedLines); i++ {
			if expectedLines[i] != actualLines[i] {
				t.Fatalf("Ошибка в тесте %s (время выполнения: %v):\n"+
					"Входные данные: %s\nОжидалось: %s\nПолучено: %s",
					inputFile, duration, inputLines[i], expectedLines[i], actualLines[i])
			}
		}

		t.Logf("Тест %s успешно пройден. (время выполнения: %v)", inputFile, duration)
	}
}
func extractNumber(fileName string) int {
	for i := 0; i < len(fileName); i++ {
		if fileName[i] >= '0' && fileName[i] <= '9' {
			var num int
			fmt.Sscanf(fileName[i:], "%d", &num)
			return num
		}
	}
	return 0 // Если числа в имени файла нет
}
