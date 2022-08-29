package translator

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

func New() Translator {
	return translator
}
