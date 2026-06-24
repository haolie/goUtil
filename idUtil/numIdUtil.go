package idutil

import (
	"sync/atomic"
)

var (
	defaultNumFactory *NumIdFactory
)

func init() {
	defaultNumFactory = NewNumIdFactory(0, 1)
}

// InitDefaultNumFactory 重置包级默认工厂。seed 为起始值，add 为每次递增步长。
func InitDefaultNumFactory(seed, add int64) {
	defaultNumFactory = NewNumIdFactory(seed, add)
}

// GetDefault 返回包级默认 NumIdFactory 实例。
func GetDefault() *NumIdFactory {
	return defaultNumFactory
}

// NumIdFactory 基于原子操作的自增 ID 工厂，并发安全。
type NumIdFactory struct {
	seed int64
	add  int64
}

// NewNumIdFactory 创建一个新的 ID 工厂。seed 为初始值，add 为步长，首次 CreateId 返回 seed+add。
func NewNumIdFactory(seed, add int64) *NumIdFactory {
	return &NumIdFactory{
		seed: seed,
		add:  add,
	}
}

// CreateId 原子递增并返回下一个 ID。
func (n *NumIdFactory) CreateId() int64 {
	return atomic.AddInt64(&n.seed, n.add)
}
