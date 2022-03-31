package tf

import (
	"fmt"
	"sync"
	"time"
)

// simple uuid 25 length

type UUID struct {
	idx int
	m   sync.Mutex
}

func NewUUID() UUID {
	return UUID{}
}

// 1585539806824507500000002
func (u *UUID) New() string {
	u.m.Lock()
	u.idx++
	idx := u.idx
	if u.idx >= 10^7 {
		u.idx = 0
	}
	u.m.Unlock()
	return fmt.Sprintf("%d%06d", time.Now().UnixNano(), idx)
}
