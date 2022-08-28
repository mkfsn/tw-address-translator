//go:build tools
// +build tools

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tealeg/xlsx/v3"
)

func main() {
	url := "https://www.post.gov.tw/post/download/6.5_CEROAD11107.xlsx"
	outputFile := "./road.csv"

	xlsxFile, err := downloadXLSXFile(url)
	if err != nil {
		log.Fatalln(err)
	}

	if err := writeFile(xlsxFile, outputFile); err != nil {
		log.Fatalln(err)
	}
}

func downloadXLSXFile(url string) (*xlsx.File, error) {
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
		return nil, fmt.Errorf("failed to read http body: %w", err)
	}

	f, err := xlsx.OpenBinary(b)
	if err != nil {
		return nil, fmt.Errorf("failed to open xlsx binary: %w", err)
	}

	return f, nil
}

func writeFile(xlsxFile *xlsx.File, filename string) error {
	csvFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to craete csv file: %w", err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	sheets, err := xlsxFile.ToSlice()
	if err != nil {
		log.Fatalln(err)
	}

	return csvWriter.WriteAll(sheets[0])
}
