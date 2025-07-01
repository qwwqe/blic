package blic

type EventType string

const (
	GameCreatedEventType      EventType = "gamecreated"
	LoanActionTakenEventType  EventType = "loanactiontaken"
	PassActionTakenEventType  EventType = "passactiontaken"
	ScoutActionTakenEventType EventType = "scoutactiontaken"
	TurnEndedEventType        EventType = "turnended"
)

// TODO: Think about what should appear here and whether it needs to exist at all.
type Event interface{}

type GameCreatedEvent struct {
	Type EventType

	GameId          string
	GameSpecName    string
	GameSpecVersion string

	Deck                []Card
	Locations           []Location
	CanalEraConnections []Connection
	RailEraConnections  []Connection
	Players             []Player
	IncomeTrack         []int

	NumWildLocationCards int
	NumWildIndustryCards int

	InitialCoalInMarket int
	InitialIronInMarket int

	ActionsPerTurn int

	LoanAmount             int
	LoanIncomeLevelPenalty int
	HandSize               int
}

type LoanActionTakenEvent struct {
	Type EventType

	PlayerId        string
	DiscardedCardId string
}

type PassActionTakenEvent struct {
	Type EventType

	PlayerId        string
	DiscardedCardId string
}

type ScoutActionTakenEvent struct {
	Type EventType

	PlayerId         string
	DiscardedCardIds []string
}

type TurnEndedEvent struct {
	Type EventType

	PlayerId string
}

var _ Event = (*GameCreatedEvent)(nil)
