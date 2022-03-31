package tf

import (
	"sync"
)

// SafeMapStringString 并发安全的 map[string]string
type SafeMapStringString struct {
	m map[string]string
	l *sync.RWMutex
}

// SafeMapStringBool 并发安全的map[string]bool
type SafeMapStringBool struct {
	m map[string]bool
	l *sync.RWMutex
}

// NewSafeMapStringBool 初始化 SafeMapStringBool
func NewSafeMapStringBool() *SafeMapStringBool {
	sm := new(SafeMapStringBool)
	sm.m = make(map[string]bool)
	sm.l = new(sync.RWMutex)
	return sm
}
