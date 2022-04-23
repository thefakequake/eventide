package eventide

import (
	"reflect"
)

var eventCodec = map[string][]Handler

func NewEventCodec() *EventUnmarshaller {
	return &HandlerManager{
		callbacks: map[string][]reflect.Value{},
	}
}

type EventCodec struct {
	events map[string]Event
}

func NewEventCodec(events ...Event) *EventCodec {
	c := &EventCodec{events: make(map[string]Event)}
	for _, e := range events {
		c.events[e.Type()] = e
	}
}

func (e *EventCodec) DecodeEvent(type string, dat []byte) Event {
	e.events[]
}

func (h *HandlerManager) Call(event string) {
	
	
	c, ok := h.callbacks[event]
	if !ok {
		// event not supported
		return
	}


} 

func (s *Session)
