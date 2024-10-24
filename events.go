package main

type EventType string

const (
	EventTypeGameCreated EventType = "gamecreated"
)

// TODO: Think about what should appear here and whether it needs to exist at all.
type Event interface{}

type GameCreatedEvent struct {
	Deck      []Card
	Locations []Location
	Players   []Player
}

var _ Event = (*GameCreatedEvent)(nil)
