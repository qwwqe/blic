package main

type Era string

const (
	EraCanal Era = "canal"
	EraRail  Era = "rail"
)

type GamePhase string

const (
	GamePhaseAction GamePhase = "action"
)

type Game struct {
	Id     string
	Events []Event

	Players   []Player
	Locations []Location

	CoalInMarket int
	IronInMarket int

	Era Era

	Deck              []Card
	WildLocationCards []Card
	WildIndustryCards []Card

	PlayerIndex int
	Round       int
	Phase       GamePhase
}

func (g *Game) HandleGameCreatedEvent(e GameCreatedEvent) {
	g.Id = e.Id
	g.Events = []Event{}

	g.Players = CloneSlice(g.Players)
	g.Deck = CloneSlice(e.Deck)
	// Locations
}
