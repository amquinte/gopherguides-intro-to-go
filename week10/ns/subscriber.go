package ns

import (
	"context"
	"sync"
)

type Subscriber struct {
	cancel    context.CancelFunc
	suscribed bool
	stories   chan *Story
	sync.RWMutex
}
