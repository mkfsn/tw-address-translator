package translator

import (
	_ "embed"
	"encoding/xml"
	"strings"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

//go:embed dataset/County_h_10906.xml
var countyXML []byte

type Counties struct {
	Counties []County `xml:"County_h_10906"`
}

type County struct {
	Code        int    `xml:"欄位1"`
	ChineseName string `xml:"欄位2"`
	EnglishName string `xml:"欄位3"`
}

func buildCountyTrie(countyXML []byte) *trie.Trie {
	var counties Counties

	if err := xml.Unmarshal(countyXML, &counties); err != nil {
		panic("failed to parse county data: invalid dataset")
	}

	newTrie := trie.NewTrie()

	for _, county := range counties.Counties {
		county := county

		newTrie.Insert(county.ChineseName, &county)

		if strings.Contains(county.ChineseName, "臺") {
			newName := strings.Replace(county.ChineseName, "臺", "台", -1)
			newTrie.Insert(newName, &county)
		}
	}

	return newTrie
}
