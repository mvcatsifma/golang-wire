//+build wireinject

package events

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	//wire.Build(NewEvent, NewGreeter, NewMessage)
	wire.Build(NewEvent, NewMessage)
	return Event{}, nil
}
