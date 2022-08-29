package norm

import (
	"bytes"
	"unicode/utf8"

	"golang.org/x/text/transform"
)

var (
	chineseHyphenTransformer transform.Transformer = &textTransformer{
		translations: map[rune]string{
			'之': "-",
		},
	}

	chineseNumberTransformer transform.Transformer = &textTransformer{
		translations: map[rune]string{
			'一': "1",
			'二': "2",
			'三': "3",
			'四': "4",
			'五': "5",
			'六': "6",
			'七': "7",
			'八': "8",
			'九': "9",
		},
	}
)

type textTransformer struct {
	translations map[rune]string
}

func (n *textTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	result, err := n.replace(src)
	if err != nil {
		return 0, 0, nil
	}

	size := copy(dst, result)
	if size < len(src) {
		err = transform.ErrShortDst
	}

	return size, len(src), err
}

func (n *textTransformer) Reset() {
	// Intended to be empty
}

func (n *textTransformer) replace(src []byte) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, len(src)*2))

	for i := 0; i < len(src); {
		r, width := utf8.DecodeRune(src[i:])

		if d, ok := n.translations[r]; ok {
			buf.WriteString(d)
		} else {
			buf.WriteRune(r)
		}

		i += width
	}

	return buf.Bytes(), nil
}
