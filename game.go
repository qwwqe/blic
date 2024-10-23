package main

type Era string

const (
	EraCanal Era = "canal"
	EraRail  Era = "rail"
)

type RoundPhase string

type TurnPhase string

type Game struct {
	Id string

	Players   []Player
	Locations []Location

	CoalInMarket int
	IronInMarket int

	Era Era

	Deck              []Card
	WildLocationCards []Card
	WildIndustryCards []Card

	CurrentPlayerIndex int
	CurrentRoundPhase  RoundPhase
	CurrentTurnPhase   TurnPhase
}
