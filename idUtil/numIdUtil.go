package idUtil

import (
	"sync/atomic"
)

var (
	defaultNumFactory *NumIdFactory
)

func init() {
	defaultNumFactory = NewNumIdFactory(0, 1)
}

func InitDefaultNumFactory(seed, add int64) {
	defaultNumFactory = NewNumIdFactory(seed, add)
}

type NumIdFactory struct {
	seed int64
	add  int64
}

func NewNumIdFactory(seed, add int64) *NumIdFactory {
	return &NumIdFactory{
		seed: seed,
		add:  add,
	}
}

func (n *NumIdFactory) CreateId() int64 {
	return atomic.AddInt64(&n.seed, n.add)
}
