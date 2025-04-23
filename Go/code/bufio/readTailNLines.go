package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// from file end to head, count numLines '\n' then get the lines
func tail(file *os.File, numLines int) ([]string, error) {
	var lines []string
	var lineCount int
	var readPosition int64

	bufferSize := 1024
	buffer := make([]byte, bufferSize)

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	for {
		readPosition = fileSize - int64(bufferSize)

		if readPosition < 0 {
			readPosition = 0
			break
		}

		_, err := file.ReadAt(buffer, readPosition)
		if err != nil {
			return nil, err
		}

		lineCount += bytes.Count(buffer, []byte{'\n'})
		if lineCount >= numLines {
			break
		}

		fileSize = readPosition
	}

	file.Seek(readPosition, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if readPosition == 0 {
		return lines, nil
	} else {
		return lines[1:], nil
	}
}

func main() {
	filePath := "file.txt"
	numLines := 20 // 你想要读取的行数

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines, err := tail(file, numLines)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 打印读取到的行数
	for _, line := range lines {
		fmt.Println(line)
	}
}
