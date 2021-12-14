package ns

import (
	"context"
	"sync"
)

type Subscriber struct {
	cancel    context.CancelFunc
	suscribed bool
	//stories   []Story
	stories chan *Story
	sync.RWMutex
}

// func (n *Ns) Subscribe(category string, subName string, subsChan chan *Subscriber) error{

// 	n.subscibers[subName] = subsChan
// 	//var tmp string
// 	for s := range n.stories[category]{
// 		 := <- s
// 		stories = append(stories, msg)
// 	}
// 	return nil
// }
