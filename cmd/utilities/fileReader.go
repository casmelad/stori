package utilities

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type FileReader struct {
}

func (fr FileReader) ReadTransationLinesFromCsvFile(fullURLFile string) ([]string, error) {
	// Get file from git
	resp, err := http.Get(fullURLFile)
	if err != nil {
		fmt.Printf("file not found")
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respString := buf.String()
	scanner := bufio.NewScanner(strings.NewReader(respString))

	fileRecords := []string{}
	for scanner.Scan() {
		fileRecords = append(fileRecords, scanner.Text())
	}

	defer resp.Body.Close()

	return fileRecords[1:], nil
}
