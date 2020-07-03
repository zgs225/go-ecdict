package dict

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/zgs225/go-ecdict/index"
)

var (
	// ErrShortRead 记录未完全读取
	ErrShortRead = errors.New("记录未读取完全")
)

// SimpleDict 使用 Simple 索引的词典查询
type SimpleDict struct {
	file  io.ReadSeeker
	index index.Interface
}

// NewSimpleDict 生成查询词典并构建索引
func NewSimpleDict(f string) (*SimpleDict, error) {
	dictFile, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	is, err := index.BuildSimpleIndex(dictFile, true)
	if err != nil {
		return nil, err
	}
	if _, err = dictFile.Seek(0, 0); err != nil {
		return nil, err
	}
	return &SimpleDict{
		file:  dictFile,
		index: is,
	}, nil
}

// Match 查询词典
func (d *SimpleDict) Match(k string) (*Record, error) {
	i, err := d.index.Match(k)
	if err != nil {
		return nil, err
	}
	r := Record{}

	if err = d.readRecordByIndex(&r, i); err != nil {
		return nil, err
	}
	return &r, nil
}

// Like 根据最左匹配原则获取满足条件的所有记录
func (d *SimpleDict) Like(k string) ([]*Record, error) {
	is, err := d.index.Like(k)
	if err != nil {
		return nil, err
	}

	v := make([]*Record, len(is))

	// FIXME: 将连续的索引记录合并成一次文件读取
	for i, idx := range is {
		r := Record{}
		if err = d.readRecordByIndex(&r, idx); err != nil {
			return nil, err
		}
		v[i] = &r
	}

	return v, nil
}

func (d *SimpleDict) readRecordByIndex(r *Record, i *index.Item) error {
	b := make([]byte, i.Len)
	_, err := d.file.Seek(int64(i.Pos), io.SeekStart)
	if err != nil {
		return err
	}
	n, err := d.file.Read(b)
	if int32(n) < i.Len {
		return ErrShortRead
	}
	return d.readRecordFromBytes(r, b[:len(b)-1])
}

func (d *SimpleDict) readRecordFromBytes(r *Record, b []byte) error {
	cr := csv.NewReader(bytes.NewBuffer(b))
	ss, err := cr.Read()
	if err != nil {
		return err
	}
	r.Word = ss[0]
	r.Phonetic = ss[1]
	r.Definition = ss[2]
	r.Translation = ss[3]
	r.Pos = ss[4]
	r.Collins = ss[5]
	r.Oxford = ss[6]
	r.Tag = ss[7]

	bnc := ss[8]
	if len(bnc) > 0 {
		n, _ := strconv.ParseInt(bnc, 10, 64)
		r.Bnc = int(n)
	}

	frq := ss[9]
	if len(frq) > 0 {
		n, _ := strconv.ParseInt(frq, 10, 64)
		r.Frq = int(n)
	}
	r.Exchange = ss[10]
	return nil
}

var _ Interface = (*SimpleDict)(nil)
