package blic

type EventType string

const (
	EventTypeGameCreated EventType = "gamecreated"
)

// TODO: Think about what should appear here and whether it needs to exist at all.
type Event interface{}

type GameCreatedEvent struct {
	Id   string
	Type EventType

	GameSpecName    string
	GameSpecVersion string

	Deck                []Card
	Locations           []Location
	CanalEraConnections []Connection
	RailEraConnections  []Connection
	Players             []Player

	NumWildLocationCards int
	NumWildIndustryCards int

	InitialCoalInMarket int
	InitialIronInMarket int
}

var _ Event = (*GameCreatedEvent)(nil)
