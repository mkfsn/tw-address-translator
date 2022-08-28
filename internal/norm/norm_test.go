package norm_test

import (
	"testing"

	"github.com/mkfsn/tw-address-translator/internal/norm"
	"github.com/stretchr/testify/assert"
)

func TestAddressNormalizer_Normalize(t *testing.T) {
	tests := []struct {
		address string
		want    string
		err     error
	}{
		{
			address: "１０樓",
			want:    "10樓",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			n := norm.NewAddressNormalizer()
			got, err := n.Normalize(tt.address)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				assert.Empty(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}
