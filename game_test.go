package blic

import (
	"reflect"
	"testing"
)

func TestHandleGameCreatedEvent(t *testing.T) {
	t.Run("records event in store", func(t *testing.T) {
		e := GameCreatedEvent{GameId: "test"}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if !reflect.DeepEqual(game.Events[0], e) {
			t.Errorf("expected %v but got %v", e, game.Events[0])
		}
	})

	t.Run("clears existing events in store", func(t *testing.T) {
		e := GameCreatedEvent{GameId: "test"}
		game := Game{
			Events: []Event{GameCreatedEvent{}, GameCreatedEvent{}},
		}
		game.HandleGameCreatedEvent(e)

		if !reflect.DeepEqual(game.Events, []Event{e}) {
			t.Errorf("expected %v but got %v", []Event{e}, game.Events)
		}
	})

}
