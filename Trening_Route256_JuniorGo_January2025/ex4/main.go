package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		root := readInput(in)
		countInf := countInfectedFiles(root, false)
		fmt.Fprintln(out, countInf)
	}
}

func readInput(in *bufio.Reader) (root Folder) {
	var n int
	fmt.Fscan(in, &n)
	in.ReadString('\n')

	var lines []string

	for i := 0; i < n; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка при чтении строки %d: %v\n", i+1, err)
			return
		}
		lines = append(lines, line)
	}

	jsonData := []byte(strings.Join(lines, ""))

	err := json.Unmarshal(jsonData, &root)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
	}
	return
}

func countInfectedFiles(folder Folder, parentInfected bool) int {
	isInfected := parentInfected

	for _, file := range folder.Files {
		if strings.HasSuffix(file, ".hack") {
			isInfected = true
			break
		}
	}

	infectedFiles := 0
	if isInfected {
		infectedFiles += len(folder.Files)
	}

	for _, subfolder := range folder.Folders {
		infectedFiles += countInfectedFiles(subfolder, isInfected)
	}

	return infectedFiles
}
