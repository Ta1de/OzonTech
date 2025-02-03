package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	testFolder := "even-strings"

	files, err := os.ReadDir(testFolder)
	if err != nil {
		t.Fatalf("Ошибка при чтении папки с тестами: %v", err)
	}

	// Фильтруем только входные файлы (без .a)
	var inputFiles []string
	for _, file := range files {
		if !file.IsDir() && !strings.HasSuffix(file.Name(), ".a") {
			inputFiles = append(inputFiles, file.Name())
		}
	}

	// Сортируем файлы по числовому значению
	sort.Slice(inputFiles, func(i, j int) bool {
		return extractNumber(inputFiles[i]) < extractNumber(inputFiles[j])
	})

	for testIndex, inputFile := range inputFiles {
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
			t.Fatalf("Тест %d (%s) превысил лимит времени (3000 мс). Выполнение заняло %d мс", testIndex+1, inputFile, duration.Milliseconds())
		}

		if err != nil {
			t.Fatalf("Ошибка при выполнении main.go с входным файлом %s: %v", inputFile, err)
		}

		expectedLines := strings.Split(strings.TrimSpace(string(expectedOutputData)), "\n")
		actualLines := strings.Split(strings.TrimSpace(string(output)), "\n")
		inputLines := strings.Split(strings.TrimSpace(string(inputData)), "\n")

		// Количество наборов данных указано в первой строке файла
		if len(inputLines) == 0 {
			t.Fatalf("Файл %s пуст или некорректен", inputFile)
		}

		// Ищем первую ошибку и выводим соответствующий набор данных
		for i := range expectedLines {
			if i >= len(actualLines) {
				t.Fatalf("❌ В ТЕСТЕ %s ОШИБКА В НАБОРЕ %d\nОжидалось: %q\nПолучено: --- (отсутствует строка)\nВходные данные:\n%s\n",
					inputFile, i+1, expectedLines[i], extractDataSet(inputLines, i))
			}
			if expectedLines[i] != actualLines[i] {
				t.Fatalf("❌ В ТЕСТЕ %s ОШИБКА В НАБОРЕ %d\nОжидалось: %q\nПолучено: %q\nВходные данные:\n%s\n",
					inputFile, i+1, expectedLines[i], actualLines[i], extractDataSet(inputLines, i))
			}
		}

		if len(actualLines) > len(expectedLines) {
			t.Fatalf("❌ В ТЕСТЕ %s ЛИШНИЕ ДАННЫЕ НАЧИНАЯ С НАБОРА %d\n", inputFile, len(expectedLines)+1)
		}

		t.Logf("✅ Тест %s успешно пройден. (время выполнения: %v)", inputFile, duration)
	}
}

// Функция извлекает номер из имени файла (например, "test_5.txt" → 5)
func extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(fileName)
	if match == "" {
		return 0
	}
	num, _ := strconv.Atoi(match)
	return num
}

// Функция извлекает соответствующий набор данных по индексу
func extractDataSet(inputLines []string, index int) string {
	// Проверяем, чтобы индекс был в пределах допустимого диапазона
	if index < len(inputLines) {
		return inputLines[index] // Возвращаем конкретный набор данных по индексу
	}
	// Если индекс выходит за пределы, возвращаем значение по умолчанию
	return "(отсутствует)"
}
