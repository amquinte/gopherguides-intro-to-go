package assignment04

import (
	"io"
	//"fmt"
)

type Venue struct {
	Audience int
	Log      io.Writer
}

type Entertainer interface {
	Name() string
	Perform(v Venue) error
}

type Setuper interface {
	Setup(v Venue) error
}

type Teardowner interface {
	Teardown(v Venue) error
}
