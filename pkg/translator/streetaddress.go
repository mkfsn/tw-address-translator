package translator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// street address extractor
// https://en.wikipedia.org/wiki/Postal_addresses_in_Taiwan
type streetAddressTranslator struct{}

func (s *streetAddressTranslator) Translate(address string) (string, error) {
	info, err := extractStreetAddressInfo(address)
	if err != nil {
		return "", err
	}

	return info.String(), nil
}

type streetAddressInfo struct {
	Lane        string
	Alley       string
	HouseNumber string
	Floor       string
	Room        string
}

func (s streetAddressInfo) String() string {
	parts := make([]string, 0, 5)

	if s.Room != "" {
		parts = append(parts, fmt.Sprintf("Rm. %s", s.Room))
	}

	if s.Floor != "" {
		parts = append(parts, fmt.Sprintf("%sF.", s.Floor))
	}

	if s.HouseNumber != "" {
		parts = append(parts, fmt.Sprintf("No. %s", s.HouseNumber))
	}

	if s.Alley != "" {
		parts = append(parts, fmt.Sprintf("Aly. %s", s.Alley))
	}

	if s.Lane != "" {
		parts = append(parts, fmt.Sprintf("Ln. %s", s.Lane))
	}

	return strings.Join(parts, ", ")
}

func extractStreetAddressInfo(address string) (streetAddressInfo, error) {
	info := streetAddressInfo{}

	i, j := 0, 0

charIterator:
	for i < len(address) && j < len(address) {
		ch, _ := utf8.DecodeRuneInString(address[j:])

		switch ch {
		case '巷':
			info.Lane = address[i:j]
		case '弄':
			info.Alley = address[i:j]
		case '樓':
			info.Floor = address[i:j]
		case '室':
			info.Room = address[i:j]
		case '號':
			info.HouseNumber = address[i:j]
		default:
			j++
			continue charIterator
		}

		i = j + utf8.RuneLen(ch)
		j = i
	}

	return info, nil
}
