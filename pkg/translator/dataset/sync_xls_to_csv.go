//go:build tools
// +build tools

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/huangtao-sh/xls"
)

func main() {
	url := os.Args[1]
	outputFile := os.Args[2]

	xlsFile, err := downloadXLSFile(url)
	if err != nil {
		log.Fatalln(err)
	}

	if err := writeFile(xlsFile, outputFile); err != nil {
		log.Fatalln(err)
	}
}

func downloadXLSFile(url string) (*xls.WorkBook, error) {
	var c http.Client

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download xlsx file: %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	f, err := xls.OpenReader(bytes.NewReader(b), "")
	if err != nil {
		return nil, fmt.Errorf("failed to open xlsx binary: %w", err)
	}

	return f, nil
}

func writeFile(xlsFile *xls.WorkBook, filename string) error {
	csvFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create csv file: %w", err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	rows, err := xlsFile.GetRowsIndex(0)
	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	return csvWriter.WriteAll(rows)
}
