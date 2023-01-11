package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("At least two arguments required. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	// 출력
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()
	}

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Printf("No file list found. Error: %s", err)
		return findInfos
	}

	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("No file found. Error: %s", err)
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}