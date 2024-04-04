package dispatch

import (
	"reflect"
)

type Handler interface {
	Exec(data any)
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: map[string]any{},
	}
}

type Dispatcher struct {
	handlers map[string]any
}

func (d *Dispatcher) Register(data any, handler any) {
	key := reflect.TypeOf(data).String()
	d.handlers[key] = handler
}

func (d *Dispatcher) Exec(data any) error {
	key := reflect.TypeOf(data).String()
	handler := d.handlers[key]

	x := reflect.ValueOf(handler)
	values := []reflect.Value{
		reflect.ValueOf(data),
	}

	res := x.Call(values)[0].Interface()
	if res == nil {
		return nil
	}

	return res.(error)
}
