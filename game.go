package blic

type Era string

const (
	EraCanal Era = "canal"
	EraRail  Era = "rail"
)

type GamePhase string

const (
	GamePhaseAction GamePhase = "action"
)

const startingCoalInMarket = 13
const startingIronInMarket = 8
const startingWildCards = 4

type Game struct {
	Id     string
	Events []Event

	Players             []Player
	Locations           []Location
	CanalEraConnections []Connection
	RailEraConnections  []Connection

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

/**
 * HandleGameCreatedEvent handles the event corresponding to the start of a game.
 * Calling this method will effectively clear the Game referenced by the pointer receiver.
 */
func (g *Game) HandleGameCreatedEvent(e GameCreatedEvent) *Game {
	*g = Game{}

	g.Id = e.Id
	g.Events = []Event{e}

	g.Players = CloneSlice(g.Players)
	g.Locations = CloneSlice(e.Locations)
	g.CanalEraConnections = CloneSlice(e.CanalEraConnections)
	g.RailEraConnections = CloneSlice(e.RailEraConnections)

	g.CoalInMarket = startingCoalInMarket
	g.IronInMarket = startingIronInMarket

	g.Era = EraCanal

	g.Deck = CloneSlice(e.Deck)
	g.WildLocationCards = make([]Card, 0, startingWildCards)
	g.WildIndustryCards = make([]Card, 0, startingWildCards)
	for i := 0; i < startingWildCards; i++ {
		g.WildLocationCards = append(g.WildLocationCards, Card{Type: CardTypeWildLocation})
		g.WildIndustryCards = append(g.WildIndustryCards, Card{Type: CardTypeWildIndustry})
	}

	g.PlayerIndex = 0
	g.Round = 0
	g.Phase = GamePhaseAction

	return g
}
