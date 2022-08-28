package translator

type Translator interface {
	Translate(address string) (string, error)
}

var (
	translator *postalAddressTranslator
)

func init() {
	translator = newPostalAddressTranslator(
		buildRoadTrie(ceRoadTXT),
		buildCountyTrie(countyXML),
	)
}

func New() Translator {
	return translator
}
