package norm

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

type Normalizer interface {
}

var _ Normalizer = (*addressNormalizer)(nil)

type addressNormalizer struct {
	transformer transform.Transformer
}

func NewAddressNormalizer() *addressNormalizer {
	return &addressNormalizer{
		transformer: transform.Chain(
			width.Narrow,             // full-width numbers to half-width numbers
			chineseNumberTransformer, // chinese number to alpha-numbers
			chineseHyphenTransformer, // chinese 之 to -
		),
	}
}

func (a *addressNormalizer) Normalize(address string) (string, error) {
	result, _, err := transform.String(a.transformer, address)
	return result, err
}
