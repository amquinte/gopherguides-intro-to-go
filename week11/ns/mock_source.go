package ns

import (
	"fmt"
)

func (n *Ns) PublishStory(category string, stories ...*Story) error {

	var storyChan = make(chan *Story)
	defer close(storyChan)
	for _, s := range stories {

		fmt.Println("pushing story")
		storyChan <- s
		//n.Story(category) <- s

	}
	fmt.Println("finished pushing stories to story chan")
	n.GetStory(category, storyChan)
	//return storyChan, nil
	return nil
}

// func (n *Ns) Story(category string) chan *Story {
// 	n.Lock()
// 	defer n.Unlock()
// 	if n.stories == nil {
// 		n.stories[category] = make(chan *Story)
// 	}
// 	return n.stories[category]
// }
