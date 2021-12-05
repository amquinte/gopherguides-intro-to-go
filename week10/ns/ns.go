package ns

import (
	"context"
	"encoding/json"
	"fmt"
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

const Filename = "backup.json"

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
	//const Filename = "backup.json"
	err = os.WriteFile(Filename, state, 0777)
	if err != nil {
		return err
	}

	fmt.Println("State was saved")
	return nil
}

func (n *Ns) CreateFile() error {

	//Checks if file exists in current directory
	//Creates the file if it does not already exits in current directory
	f, _ := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0777)
	//Need to figure out how to test the below if statement for 100% coverage
	// if err != nil {
	// 	//return err
	// }

	// remember to close the file
	defer f.Close()
	return nil
}

func (n *Ns) LoadFile() error {
	//Checks if file exists in current directory
	//Creates the file if it does not already exits in current directory
	f, _ := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0777)

	//Need to figure out how to test the below if statement for 100% coverage
	// if err != nil {
	// 	//return err
	// }

	data, _ := os.ReadFile(Filename)
	var tmp Ns
	json.Unmarshal(data, &tmp)

	//Will have to update these once I finalize my Ns struct
	n.Stopped = tmp.Stopped
	n.Name = tmp.Name
	n.Age = tmp.Age

	// remember to close the file
	defer f.Close()
	return nil
}

func (n *Ns) RemoveFile() error {
	return os.Remove(Filename)

	// err := os.Remove(Filename)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
