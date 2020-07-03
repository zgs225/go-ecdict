package index

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"sort"
	"strings"
)

var (
	// ErrNil 未匹配给定的索引
	ErrNil = errors.New("未匹配")
	// ErrNotInitialized 未初始化
	ErrNotInitialized = errors.New("未初始化")
)

// ScanLinesEscapeDoubleQuotation 按换行符读取内容，但是忽视在双引号中的换行符
func ScanLinesEscapeDoubleQuotation(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	inDoubleQuotation := false

	for i, b := range data {
		if b == '"' {
			inDoubleQuotation = !inDoubleQuotation
		}

		if b == '\n' {
			if !inDoubleQuotation {
				return i + 1, data[0 : i+1], nil
			}
		}
	}

	if atEOF {
		return len(data), data, nil
	}

	// require more data
	return 0, nil, nil
}

// BuildSimpleIndex 根据给定的数据建立索引
func BuildSimpleIndex(r io.Reader, ignore ...bool) (Simple, error) {
	simple := Simple{}
	pos := 0

	scanner := bufio.NewScanner(r)
	scanner.Split(ScanLinesEscapeDoubleQuotation)

	for scanner.Scan() {
		if len(ignore) > 0 && ignore[0] && pos == 0 {
			txt := scanner.Bytes()
			pos += len(txt)
			continue
		}
		txt := scanner.Bytes()
		key := string(txt[:bytes.IndexByte(txt, ',')])
		simple = append(simple, &Item{
			Key: key,
			Len: int32(len(txt)),
			Pos: int32(pos),
		})
		pos += len(txt)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Sort(simple)

	return simple, nil
}

// Simple 简单的已排序数组索引
type Simple []*Item

// Match 查找指定 key 的索引，必须完全匹配
func (s Simple) Match(k string) (*Item, error) {
	if s == nil {
		return nil, ErrNotInitialized
	}

	return binSearch(s, 0, len(s), k)
}

func binSearch(s []*Item, si, ei int, k string) (*Item, error) {
	if si > ei {
		return nil, ErrNil
	}

	i := (si + ei) / 2

	switch strings.Compare(s[i].Key, k) {
	case 0:
		return s[i], nil
	case 1:
		return binSearch(s, si, i-1, k)
	case -1:
		return binSearch(s, i+1, ei, k)
	}

	return nil, ErrNil
}

// Like 使用最左匹配的原则匹配有 key 的所有索引
func (s Simple) Like(k string) ([]*Item, error) {
	if s == nil {
		return nil, ErrNotInitialized
	}

	return binLike(s, 0, len(s), k)
}

func binLike(s []*Item, si, ei int, k string) ([]*Item, error) {
	if si > ei {
		return nil, ErrNil
	}

	i := (si + ei) / 2

	if strings.HasPrefix(s[i].Key, k) {
		return walkLeftRight(s, k, i), nil
	}

	j := strings.Compare(s[i].Key, k)

	if j == 1 {
		return binLike(s, si, i-1, k)
	}

	if j == -1 {
		return binLike(s, i+1, ei, k)
	}

	return nil, ErrNil
}

func walkLeftRight(s []*Item, k string, i int) []*Item {
	v := make([]*Item, 0)

	if i > 0 {
		for j := i - 1; j >= 0; j-- {
			if strings.HasPrefix(s[j].Key, k) {
				v = append(v, s[j])
				continue
			}
			break
		}
	}

	v = append(v, s[i])

	if i < len(s)-1 {
		for j := i + 1; j < len(s); j++ {
			if strings.HasPrefix(s[j].Key, k) {
				v = append(v, s[j])
				continue
			}
			break
		}
	}

	return v
}

func (s Simple) Len() int {
	return len(s)
}

func (s Simple) Less(i, j int) bool {
	return strings.Compare(s[i].Key, s[j].Key) == -1
}

func (s Simple) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ Interface = Simple(nil)
