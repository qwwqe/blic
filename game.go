package main

type Era string

const (
	EraCanal Era = "canal"
	EraRail  Era = "rail"
)

type Game struct {
	Id int

	Players   []Player
	Locations []Location

	CoalInMarket int
	IronInMarket int

	Era Era

	Deck              []Card
	WildLocationCards []Card
	WildIndustryCards []Card
}
