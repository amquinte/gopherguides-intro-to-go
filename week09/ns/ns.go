package ns

import (
	"context"
	"sync"
)

type Ns struct {
	cancel  context.CancelFunc
	errs    chan error
	subs    chan *Subscriber
	stopped bool
	sync.RWMutex
	sync.Once
}

func (n *Ns) Start(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
