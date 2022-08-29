package translator

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

//go:embed dataset/village.csv
var villageCSVData []byte

type Village struct {
	ChineseName string
	EnglishName string
}

func buildVillageTrie(rawData []byte) *trie.Trie {
	newTrie := trie.NewTrie()

	csvReader := csv.NewReader(bytes.NewReader(rawData))

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(fmt.Errorf("failed to read csv records: %s", err))
	}

	for _, record := range records {
		road := Village{
			ChineseName: record[0],
			EnglishName: record[1],
		}

		newTrie.Insert(road.ChineseName, &road)
	}

	return newTrie
}
