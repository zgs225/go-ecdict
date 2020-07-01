package index

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
