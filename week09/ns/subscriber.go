package ns

import (
	"context"
	"sync"
)

type Subscriber struct {
	cancel     context.CancelFunc
	categories []string
	suscribed  bool
	stories    []string
	sync.RWMutex
}
