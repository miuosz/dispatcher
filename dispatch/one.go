package dispatch

import (
	"errors"
	"fmt"
)

type (
	HandlerOne     struct{}
	HandlerOneData struct {
		Message string
	}
)

func (h *HandlerOne) Handler() func(d HandlerOneData) error {
	return func(d HandlerOneData) error {
		fmt.Println(d.Message, "it was called in one")
		return errors.New("error here")
	}
}
