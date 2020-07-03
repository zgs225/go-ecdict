package dict

// Interface 词典的接口
type Interface interface {
	Match(string) (*Record, error)
}
