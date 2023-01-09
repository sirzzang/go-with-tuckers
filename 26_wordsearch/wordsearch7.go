package main

import (
	"bufio"
	"fmt"
	"io/fs"
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

func GetFileList(pattern string) ([]string, error) {
	fileList := []string{}

	// 현재 폴더의 하위 모든 폴더를 재귀적으로 탐색
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				fileList = append(fileList, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return fileList, nil
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Printf("No file list found. Error: %s", err)
		return findInfos
	}

	// make channel
	ch := make(chan FindInfo)

	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
		// findInfos = append(findInfos, FindWordInFile(word, filename))
	}

	// retrieve from channel and append to result slice
	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt { // all received
			break
		}
	}

	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("No file found. Error: %s", err)
		ch <- findInfo
		return
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
	ch <- findInfo // send findInfo into channel
}
