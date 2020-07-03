package index

import "fmt"

// Item 单条索引记录
type Item struct {
	Key string
	Pos int32
	Len int32
}

// Interface 词典索引
type Interface interface {
	Match(k string) (*Item, error)
	Like(k string) ([]*Item, error)
}

func (i Item) String() string {
	return fmt.Sprintf(`Item = {
		Key = %s,
		Pos = %d,
		Len = %d
}`, i.Key, i.Pos, i.Len)
}
