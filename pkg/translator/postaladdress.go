package translator

import (
	"fmt"
	"strings"

	"github.com/mkfsn/tw-address-translator/internal/trie"
)

type postalAddressTranslator struct {
	countyTrie              *trie.Trie
	villageTrie             *trie.Trie
	roadTrie                *trie.Trie
	maxCountyLen            int
	maxRoadLen              int
	streetAddressTranslator Translator
}

func newPostalAddressTranslator(countyTrie, villageTrie, roadTrie *trie.Trie) *postalAddressTranslator {
	w := &postalAddressTranslator{
		countyTrie:              countyTrie,
		villageTrie:             villageTrie,
		roadTrie:                roadTrie,
		maxRoadLen:              roadTrie.MaxDepth(),
		maxCountyLen:            countyTrie.MaxDepth(),
		streetAddressTranslator: &streetAddressTranslator{},
	}

	return w
}

func (p *postalAddressTranslator) Translate(address string) (result string, err error) {
	val, ok := p.countyTrie.Search(address)
	if !ok {
		return "", ErrNoCounty
	}
	county := val.(*County)
	address = address[len(county.ChineseName):]

	val, ok = p.roadTrie.Search(address)
	if !ok {
		return "", ErrNoRoad
	}
	road := val.(*Road)
	address = address[len(road.ChineseName):]

	remaining, _ := p.streetAddressTranslator.Translate(address)

	return strings.Join(
		[]string{
			remaining,
			road.EnglishName,
			fmt.Sprintf("%s %s", county.EnglishName, county.Code),
			"Taiwan",
		},
		", ",
	), nil
}
