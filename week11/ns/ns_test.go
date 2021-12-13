package ns

import (
	"context"
	"fmt"
	"testing"
)

func Test_Ns(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	n := &Ns{}

	ctx, err := n.Start(ctx)
	if err != nil {
		t.Fatal("unexpected error")
	}

	fmt.Println("test passed")
}

//Need to figure out how to get this test to pass
//How can me make Marshall throw an error?

// func Test_Ns_Save_Err(t *testing.T) {
// 	t.Parallel()
// 	//ctx := context.Background()
// 	//n := &Ns{}
// 	var n Ns{}
// 	err := n.Save()

// 	if err == nil {
// 		t.Fatalf("expected error but received: %s", err)
// 	}
// 	fmt.Println("Second test passed")
// }

func Test_Ns_File(t *testing.T) {
	t.Parallel()
	n := &Ns{}

	err := n.CreateFile()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	n = &Ns{Name: "Tony", Age: 28}
	s := Story{Id: 0, Category: "sport", Content: "Yankees won the superbowl"}
	go n.PublishStory("sport", &s)
	err = n.Save()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	n1 := Ns{}
	//fmt.Println(n1)
	err = n1.LoadFile()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	//fmt.Println(n1)
	fmt.Println("removing file")
	n1.RemoveFile()
}
