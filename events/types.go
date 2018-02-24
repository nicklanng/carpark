package events

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Event must apply to all events.
// All events have a sequence number and a timestamp.
type Event interface {
	proto.Message
	GetAt() *timestamp.Timestamp
}

// EventType returns the type of an event
func EventType(event Event) string {
	t := reflect.TypeOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
