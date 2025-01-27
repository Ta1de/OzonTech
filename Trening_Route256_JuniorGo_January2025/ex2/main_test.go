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
	// Папка с тестовыми файлами
	testFolder := "high-load-system"

	// Получаем список файлов из папки
	files, err := os.ReadDir(testFolder)
	if err != nil {
		t.Fatalf("Ошибка при чтении папки с тестами: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		// Получаем имена файлов
		nameI := files[i].Name()
		nameJ := files[j].Name()

		// Извлекаем числовую часть из имени файла
		numberI := extractNumber(nameI)
		numberJ := extractNumber(nameJ)

		// Сравниваем числовые значения
		return numberI < numberJ
	})

	// Проходимся по всем файлам
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputFile := file.Name()
		if strings.HasSuffix(inputFile, ".a") {
			continue
		}

		expectedOutputFile := inputFile + ".a"

		// Читаем входные данные
		inputData, err := os.ReadFile(testFolder + "/" + inputFile)
		if err != nil {
			t.Errorf("Ошибка при чтении файла %s: %v", inputFile, err)
			continue
		}

		// Читаем ожидаемый результат
		expectedOutputData, err := os.ReadFile(testFolder + "/" + expectedOutputFile)
		if err != nil {
			t.Errorf("Ошибка при чтении файла %s: %v", expectedOutputFile, err)
			continue
		}

		// Создаём контекст с тайм-аутом
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		cmd.Stdin = bytes.NewReader(inputData)

		startTime := time.Now()

		// Выполняем команду
		output, err := cmd.Output()
		duration := time.Since(startTime)

		// Проверяем тайм-аут
		if ctx.Err() == context.DeadlineExceeded {
			t.Errorf("Тест %s превысил лимит времени (3000 мс). Выполнение заняло %d мс", inputFile, duration.Milliseconds())
			continue
		}

		// Проверяем ошибки выполнения
		if err != nil {
			t.Errorf("Ошибка при выполнении main.go с входным файлом %s: %v", inputFile, err)
			continue
		}

		// Сравниваем результат
		actualOutput := strings.TrimSpace(string(output))
		expectedOutput := strings.TrimSpace(string(expectedOutputData))
		if actualOutput != expectedOutput {
			t.Errorf("Несоответствие для теста %s (время выполнения: %v):\nОжидалось:\n%s\nПолучено:\n%s",
				inputFile, duration, expectedOutput, actualOutput)
		} else {
			t.Logf("Тест %s успешно пройден. Ожидание совпало с результатом. (время выполнения: %v)", inputFile, duration)
		}
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
