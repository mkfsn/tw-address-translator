package translator

// Translator translates the provided address from Chinese to English
type Translator interface {
	Translate(address string) (string, error)
}

var (
	translator *postalAddressTranslator
)

func init() {
	translator = newPostalAddressTranslator(
		buildCountyTrie(countyCSVData),
		buildVillageTrie(villageCSVData),
		buildRoadTrie(roadCSVData),
	)
}

// New returns a chinese-to-english address translator.
func New() Translator {
	return translator
}
