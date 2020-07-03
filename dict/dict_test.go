package dict

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dictFile() string {
	p := os.Getenv("EDICT")
	if len(p) == 0 {
		p = "/Users/lucky/Development/Other/ECDICT/ecdict.csv"
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

	t.Log(r)
}
