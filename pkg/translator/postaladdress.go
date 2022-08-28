package translator

import (
	"fmt"
	"strings"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

type postalAddressTranslator struct {
	countyTrie              *trie.Trie
	roadTrie                *trie.Trie
	maxCountyLen            int
	maxRoadLen              int
	streetAddressTranslator Translator
}

func newPostalAddressTranslator(roadTrie *trie.Trie, countyTrie *trie.Trie) *postalAddressTranslator {
	w := &postalAddressTranslator{
		countyTrie:              countyTrie,
		roadTrie:                roadTrie,
		maxRoadLen:              roadTrie.MaxDepth(),
		maxCountyLen:            countyTrie.MaxDepth(),
		streetAddressTranslator: &streetAddressTranslator{},
	}

	return w
}

func (p *postalAddressTranslator) Translate(address string) (result string, err error) {
	originCounty, translatedCounty, ok := p.countyTrie.Search(address)
	if ok {
		address = address[len(originCounty):]
	}

	originRoad, translatedRoad, ok := p.roadTrie.Search(address)
	if ok {
		address = address[len(originRoad):]
	}

	remaining, _ := p.streetAddressTranslator.Translate(address)

	return strings.Join(
		[]string{
			remaining,
			translatedRoad.(string),
			fmt.Sprintf("%s %d", translatedCounty.(*County).EnglishName, translatedCounty.(*County).Code),
			"Taiwan",
		},
		", ",
	), nil
}
