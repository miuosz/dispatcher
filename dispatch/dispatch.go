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
	d.handlers[handlerKey(data)] = handler
}

func (d *Dispatcher) Exec(data any) error {
	handler := d.handlers[handlerKey(data)]

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

func handlerKey(t any) string {
	return reflect.TypeOf(t).String()
}
