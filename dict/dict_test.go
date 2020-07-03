package dict

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dictFile() string {
	p := os.Getenv("EDICT")
	if len(p) == 0 {
		p = os.ExpandEnv("$HOME/Development/Other/ECDICT/ecdict.csv")
	}
	return p
}

func TestDictMatch(t *testing.T) {
	dict, err := NewSimpleDict(dictFile())
	if err != nil {
		t.Fatal(err)
	}
	r, err := dict.Match("aesthete")
	if err != nil {
		t.Fatal("SimpleDict match error: ", err)
	}

	assert.Equal(t, "aesthete", r.Word)
	assert.Equal(t, "'i:sθi:t", r.Phonetic)
	assert.Equal(t, "n one who professes great sensitivity to the beauty of art and nature", r.Definition)
	assert.Equal(t, "n. 审美家, 唯美主义者", r.Translation)
	assert.Equal(t, "gre", r.Tag)
	assert.Equal(t, 34101, r.Bnc)
	assert.Equal(t, 29682, r.Frq)
	assert.Equal(t, "s:aesthetes", r.Exchange)

	keys := []string{
		"aburamycin", "aburamycin", "zymophosphate", "zymophyte", "zymoplasm", "zymoplastic", "wilfully", "wilfulness",
		"wilga", "wilgus", "wilhelm", "vertin", "vertiplane", "vertiport", "vertisol", "vertisols", "vertisporin",
		"unspotted", "unsprung", "unsqueeze", "unsqueezing", "two-bath chrome tannage", "two-beam", "two-bedroom", "two-bin system",
		"two-bit", "two-blade propeller", "nyack", "nyad", "nyaff", "nyah", "nyah-nyah", "nyala", "nyam", "nyama", "nyamps",
		"Nyamuragira", "Nyamwezi", "nyang",
	}

	for _, k := range keys {
		r, err := dict.Match(k)
		assert.NoError(t, err)
		assert.Equal(t, k, r.Word)
	}
}

func BenchmarkSimpleDict_Match(b *testing.B) {
	dict, err := NewSimpleDict(dictFile())
	if err != nil {
		b.Fatal(err)
	}

	keys := []string{
		"aburamycin", "aburamycin", "zymophosphate", "zymophyte", "zymoplasm", "zymoplastic", "wilfully", "wilfulness",
		"wilga", "wilgus", "wilhelm", "vertin", "vertiplane", "vertiport", "vertisol", "vertisols", "vertisporin",
		"unspotted", "unsprung", "unsqueeze", "unsqueezing", "two-bath chrome tannage", "two-beam", "two-bedroom", "two-bin system",
		"two-bit", "two-blade propeller", "nyack", "nyad", "nyaff", "nyah", "nyah-nyah", "nyala", "nyam", "nyama", "nyamps",
		"Nyamuragira", "Nyamwezi", "nyang",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := keys[i%len(keys)]
		dict.Match(k)
	}
}
