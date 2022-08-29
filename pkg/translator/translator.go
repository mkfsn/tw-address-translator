package translator

type Translator interface {
	Translate(address string) (string, error)
}

var (
	translator *postalAddressTranslator
)

func init() {
	translator = newPostalAddressTranslator(
		buildRoadTrie(roadCSVData),
		buildCountyTrie(countyCSVData),
	)
}

func New() Translator {
	return translator
}
