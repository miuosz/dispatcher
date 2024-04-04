package dispatch

import "fmt"

type (
	HandlerTwo     struct{}
	HandlerTwoData struct {
		Message string
	}
)

func (h *HandlerTwo) Handler() func(d HandlerTwoData) error {
	return func(d HandlerTwoData) error {
		fmt.Println(d.Message, "it was called in two")
		return nil
	}
}
