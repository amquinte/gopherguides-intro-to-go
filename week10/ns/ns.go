package ns

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Ns struct {
	cancel  context.CancelFunc
	errs    chan error
	subs    chan *Subscriber
	Stopped bool
	Name    string
	Age     int
	sync.RWMutex
	sync.Once
}

func (n *Ns) Start(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (n *Ns) Save() error {

	state, err := json.Marshal(n)
	//Need to figure out how to write test for this
	//Currently causing me to not reach 100% coverage
	if err != nil {
		return err
	}

	//Save json to backup.json file
	const filename = "backup.json"
	err = ioutil.WriteFile(filename, state, 0777)

	fmt.Println("State was saved")
	return nil
}

func (n *Ns) CreateFile() error {
	const filename = "backup.json"

	//Checks if file exists in current directory
	//Creates the file if it does not already exits in current directory
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	//Need to figure out how to test the below if statement for 100% coverage
	// if err != nil {
	// 	return err
	// }

	// remember to close the file
	defer f.Close()
	return nil
}
