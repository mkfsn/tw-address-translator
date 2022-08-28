package trie_test

import (
	"testing"

	"github.com/mkfsn/tw-address-translator/internal/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie_Search(t *testing.T) {
	tests := []struct {
		pairs [][2]string
		text  string
		want  string
		found bool
	}{
		{
			pairs: [][2]string{
				{"aaa", "b"},
				{"aa", "c"},
			},
			text:  "aaa",
			want:  "b",
			found: true,
		},
		{
			pairs: [][2]string{
				{"aaa", "b"},
				{"aa", "c"},
			},
			text:  "aa",
			want:  "c",
			found: true,
		},
		{
			pairs: [][2]string{
				{"aaa", "b"},
				{"aa", "c"},
			},
			text:  "aaaa",
			want:  "b",
			found: true,
		},
		{
			pairs: [][2]string{
				{"aaa", "b"},
				{"aa", "c"},
			},
			text:  "d",
			want:  "",
			found: false,
		},
		{
			pairs: [][2]string{
				{"測試", "a"},
				{"測測試", "b"},
			},
			text:  "測試",
			want:  "a",
			found: true,
		},
		{
			pairs: [][2]string{
				{"測試", "a"},
				{"測測試", "b"},
			},
			text:  "測測試",
			want:  "b",
			found: true,
		},
		{
			pairs: [][2]string{
				{"測試", "a"},
				{"測測試", "b"},
			},
			text:  "測測試試",
			want:  "b",
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			newTrie := trie.NewTrie()

			for _, pair := range tt.pairs {
				newTrie.Insert(pair[0], pair[1])
			}

			got, ok := newTrie.Search(tt.text)
			if !tt.found {
				assert.False(t, ok)
				assert.Nil(t, got)
				return
			}

			assert.True(t, ok)
			assert.Equal(t, tt.want, got.(string))
		})
	}
}

func TestTrie_MaxDepth(t1 *testing.T) {
	tests := []struct {
		pairs [][2]string
		want  int
	}{
		{
			pairs: [][2]string{
				{"aaa", "b"},
				{"aa", "b"},
			},
			want: 3,
		},
		{
			pairs: [][2]string{},
			want:  0,
		},
		{
			pairs: [][2]string{
				{"aa", "b"},
				{"ab", "c"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t1.Run("", func(t1 *testing.T) {
			newTrie := trie.NewTrie()

			for _, pair := range tt.pairs {
				newTrie.Insert(pair[0], pair[1])
			}

			if got := newTrie.MaxDepth(); got != tt.want {
				t1.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
