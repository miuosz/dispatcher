package main

import (
	"fmt"

	"playground/dispatch"
)

func main() {
	d := dispatch.NewDispatcher()

	d.Register(dispatch.HandlerOneData{}, (&dispatch.HandlerOne{}).Handler())
	d.Register(dispatch.HandlerTwoData{}, (&dispatch.HandlerTwo{}).Handler())

	errOne := d.Exec(dispatch.HandlerOneData{Message: "got there 1"})
	fmt.Println(errOne)

	errTwo := d.Exec(dispatch.HandlerTwoData{Message: "got there 2"})
	fmt.Println(errTwo)
}
