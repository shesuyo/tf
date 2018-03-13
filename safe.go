package wtf

import (
	"sync"
)

// SafeMapStringString 并发安全的 map[string]string
type SafeMapStringString struct {
	m map[string]string
	l *sync.RWMutex
}
