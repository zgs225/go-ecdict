package index

import (
	"errors"
	"strings"
)

var (
	// ErrNil 未匹配给定的索引
	ErrNil = errors.New("未匹配")
	// ErrNotInitlized 未初始化
	ErrNotInitlized = errors.New("未初始化")
)

// Simple 简单的已排序数组索引
type Simple []*Item

// Match 查找指定 key 的索引，必须完全匹配
func (s Simple) Match(k string) (*Item, error) {
	if s == nil {
		return nil, ErrNotInitlized
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
		return nil, ErrNotInitlized
	}

	return binLike(s, 0, len(s), k)
}

func binLike(s []*Item, si, ei int, k string) ([]*Item, error) {
	if ei > si {
		return nil, ErrNil
	}

	rt := make([]*Item, 0)

	i := (si + ei) / 2

	if strings.HasPrefix(s[i].Key, k) {
		rt = append(rt, s[i])
	}

	j := strings.Compare(s[i].Key, k)

	if j == 1 {
		v, err := binLike(s, si, i-1, k)
		if err == nil {
			return append(rt, v...), nil
		}
	}

	if j == -1 {
		v, err := binLike(s, i+1, ei, k)
		if err == nil {
			return append(rt, v...), nil
		}
	}

	if len(rt) == 0 {
		return nil, ErrNil
	}

	return rt, nil
}

var _ Interface = Simple(nil)
