package translator

import (
	_ "embed"
	"strings"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

//go:embed dataset/6.5_CEROAD11104.txt
var ceRoadTXT string

type Road struct {
	ChineseName string
	EnglishName string
}

func buildRoadTrie(roadData string) *trie.Trie {
	newTrie := trie.NewTrie()

	for _, line := range strings.Split(roadData, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		road := Road{
			ChineseName: parts[0][1 : len(parts[0])-1],
			EnglishName: parts[1][1 : len(parts[1])-2],
		}

		newTrie.Insert(road.ChineseName, &road)
	}

	return newTrie
}
