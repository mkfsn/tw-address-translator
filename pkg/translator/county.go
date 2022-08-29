package translator

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

//go:embed dataset/county.csv
var countyCSVData []byte

type County struct {
	Code        string
	ChineseName string
	EnglishName string
}

func buildCountyTrie(rawData []byte) *trie.Trie {
	csvReader := csv.NewReader(bytes.NewReader(rawData))

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(fmt.Errorf("failed to read csv records: %s", err))
	}

	newTrie := trie.NewTrie()

	for _, record := range records {
		county := County{
			Code:        record[0],
			ChineseName: record[1],
			EnglishName: record[2],
		}

		newTrie.Insert(county.ChineseName, &county)

		if strings.Contains(county.ChineseName, "臺") {
			newName := strings.Replace(county.ChineseName, "臺", "台", -1)
			newTrie.Insert(newName, &county)
		}
	}

	return newTrie
}
