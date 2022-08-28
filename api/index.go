package api

import (
	"fmt"
	"net/http"

	"github.com/mkfsn/tw-address-translator/pkg/translator"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := translator.New()

	address := r.URL.Query().Get("address")

	translatedAddress, err := t.Translate(address)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("error: %s", err))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%s", translatedAddress))
}
