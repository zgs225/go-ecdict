package index

import (
	"bytes"
	"os"
	"testing"
)

var text = `bridgeless,'brɪdʒlɪs,a. Having no bridge; not bridged.,a. 无桥的,,,,,0,0,,,
Bridgeman,'bridʒmən,,n. 布里奇曼（物理学家）,,,,,41397,0,,,
Bridgend,,,n. 布里真德 (英国国会选区),,,,,17651,0,,,
bridgeport,,n. a port in southwestern Connecticut on Long Island Sound,n. 布里奇波特（地名）,,,,,0,0,,,
bridger,,,布里杰（男子名）,,,,,0,0,,,
bridgers,,, [人名] 布里杰斯,,,,,0,0,,,
bridges,'bridʒiz,n. United States labor leader who organized the longshoremen (1901-1990),n. 桥梁；纽带（bridge的复数）,,,,,16355,0,0:bridge/1:s3,,
bridgestone,,, 普利斯通公司总部所在地：日本主要业务：轮胎橡胶,,,,,0,0,,,
bridget,'bridʒit,n. Irish abbess; a patron saint of Ireland (453-523),"n. 布丽奇特（女子名）；圣布里奇特（瑞典修女, 布里奇特勋章的创立者）",,,,,9234,0,,,
Bridgeton,,,布里奇顿（美国城市）,,,,,0,0,,,
bridgetown,'bridʒtaun,n. capital of Barbados; a port city on the southwestern coast of Barbados,n. 布里奇顿（巴巴多斯首都）,,,,,45668,0,,,
Bridgett's line,,,[医] 布里杰特氏线(指示面神经管的途径),,,,,0,0,,,
bridgette,,,布里奇特\n布丽奇特（人名）,,,,,0,0,,,
bridgeview,,, [地名] [美国] 布里奇维尤,,,,,0,0,,,
bridgeville,,, [地名] [加拿大、美国] 布里奇维尔,,,,,0,0,,,
bridgewall,,,火墙,,,,,0,0,,,
bridgeware,,,"[计] 桥接件, 转换件",,,,,0,0,,,`

func TestSimple_Match(t *testing.T) {
	r := bytes.NewBuffer([]byte(text))
	is, err := BuildSimpleIndex(r)
	if err != nil {
		t.Fatal(err)
	}
	assertMatchKey(t, is, "Bridgeman")
	assertMatchKey(t, is, "Bridgend")
	assertMatchKey(t, is, "bridgeport")
	assertMatchKey(t, is, "bridger")
	assertMatchKey(t, is, "bridgers")
	assertMatchKey(t, is, "bridges")
	assertMatchKey(t, is, "bridgestone")
	assertMatchKey(t, is, "bridget")
	assertMatchKey(t, is, "Bridgeton")
	assertMatchKey(t, is, "bridgetown")
}

func assertMatchKey(t *testing.T, is Simple, key string) {
	i, err := is.Match(key)
	if err != nil {
		t.Error("Simple.Match error: ", err)
	}
	if i.Key != key {
		t.Errorf("Simple.Match error: key not match, expected %s but got %s\n", key, i.Key)
	}
}

func getDictFilePath() string {
	dict := os.Getenv("ECDICT")
	if len(dict) == 0 {
		dict = "/Users/lucky/Development/Other/ECDICT/ecdict.csv"
	}
	return dict
}

func Benchmark_Build_SimpleIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(getDictFilePath())
		if err != nil {
			b.Fatal(err)
		}
		if _, err := BuildSimpleIndex(f); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Simple_Match(b *testing.B) {
	f, err := os.Open(getDictFilePath())
	if err != nil {
		b.Fatal(err)
	}

	is, err := BuildSimpleIndex(f)
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

	for i := 0; i < b.N; i++ {
		k := keys[i%len(keys)]
		i, err := is.Match(k)
		if err != nil {
			b.Fatal(err)
		}
		if i.Key != k {
			b.Fatal(err)
		}
	}
}
