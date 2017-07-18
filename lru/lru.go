package lru

import (
	"gitlab.1dmy.com/go-lru/doublelink"
)

type lruCache struct {
	size   int                              // lru的大小
	dlink  doublelink.DoubleNode            //lru的数据结构
	lruMap map[interface{}]doublelink.Value // 映射关系
}

func NewlruCache(size int) *lruCache {
	return &lruCache{
		size:   size,
		dlink:  doublelink.NewDlist(),
		lruMap: make(map[interface{}]doublelink.Value),
	}
}
