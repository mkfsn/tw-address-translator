package translator_test

import (
	"testing"

	"github.com/mkfsn/tw-address-translator/pkg/translator"
	"github.com/stretchr/testify/assert"
)

func Test_Translate(t *testing.T) {
	tests := []struct {
		address string
		want    string
		err     error
	}{
		{address: "臺南市南區國民路16巷11弄11號", want: "No. 11, Aly. 11, Ln. 16, Guomin Rd., South Dist., Tainan City 702, Taiwan"},
		{address: "台南市南區國民路16巷11弄11號", want: "No. 11, Aly. 11, Ln. 16, Guomin Rd., South Dist., Tainan City 702, Taiwan"},
		{address: "新竹市東區金山十一街41號4樓F室", want: "Rm. F, 4F., No. 41, Jinshan 11th St., East Dist., Hsinchu City 300, Taiwan"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := translator.New().Translate(tt.address)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
