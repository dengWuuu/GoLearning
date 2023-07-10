package main

import (
	"errors"
)

func testMain() string {
	var s string
	if err := test1(); err != nil {
		s = "alhdfsadkflhaslfaslgh"
		return s
	}
	if err := test2(); err != nil {
		s = "akdhkahdfjgkhasdf"
		return s
	}

	return s
}

func test1() error {
	return errors.New("lakjsd")
}

func test2() error {
	return errors.New("lakjsd")
}
