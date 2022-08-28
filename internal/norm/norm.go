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
	// normalize address
	return &addressNormalizer{
		transformer: transform.Chain(
			// 1. full-width
			width.Narrow,
			// 2. chinese number
			// TODO
			// 3. 之
			// TODO
		),
	}
}

func (a *addressNormalizer) Normalize(address string) (string, error) {
	result, _, err := transform.String(a.transformer, address)
	return result, err
}
