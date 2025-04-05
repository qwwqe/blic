package blic

import "fmt"

type Era string

const (
	EraCanal Era = "canal"
	EraRail  Era = "rail"
)

type HandleEventError struct {
	GameId     string
	EventIndex int
	Reason     string
}

func (e HandleEventError) Error() string {
	return fmt.Sprintf("Error handling event %d for game %s: %s", e.EventIndex, e.GameId, e.Reason)
}

func NewHandleEventError(gameId string, eventIndex int, reason string) HandleEventError {
	return HandleEventError{
		GameId:     gameId,
		EventIndex: eventIndex,
		Reason:     reason,
	}
}

type GamePhase string

const (
	GamePhaseAction GamePhase = "action"
)

type Game struct {
	Id     string
	Events []Event

	Players             []Player
	Locations           []Location
	CanalEraConnections []Connection
	RailEraConnections  []Connection

	// TODO: Consider whether settings like this should be encapsulated
	// in a config object or something.
	LoanAmount int
	HandSize   int

	CoalInMarket int
	IronInMarket int

	Era Era

	Deck              []Card
	WildLocationCards []Card
	WildIndustryCards []Card
	IncomeTrack       []int

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

	g.Id = e.GameId
	g.Events = []Event{e}

	g.Players = CloneSlice(g.Players)
	g.Locations = CloneSlice(e.Locations)
	g.CanalEraConnections = CloneSlice(e.CanalEraConnections)
	g.RailEraConnections = CloneSlice(e.RailEraConnections)

	g.CoalInMarket = e.InitialCoalInMarket
	g.IronInMarket = e.InitialIronInMarket

	g.Era = EraCanal

	g.Deck = CloneSlice(e.Deck)

	g.IncomeTrack = make([]int, len(e.IncomeTrack))
	copy(g.IncomeTrack, e.IncomeTrack)

	g.WildLocationCards = make([]Card, 0, e.NumWildLocationCards)
	for i := 0; i < e.NumWildLocationCards; i++ {
		g.WildLocationCards = append(g.WildLocationCards, Card{Type: CardTypeWildLocation})
	}

	g.WildIndustryCards = make([]Card, 0, e.NumWildIndustryCards)
	for i := 0; i < e.NumWildIndustryCards; i++ {
		g.WildIndustryCards = append(g.WildIndustryCards, Card{Type: CardTypeWildIndustry})
	}

	g.PlayerIndex = 0
	g.Round = 0
	g.Phase = GamePhaseAction

	return g
}

func (g *Game) HandleLoanActionTakenEvent(e LoanActionTakenEvent) error {
	/** action boilerplate */
	playerIndex := -1
	for index, player := range g.Players {
		if player.Id == e.PlayerId {
			playerIndex = index
			break
		}
	}

	if playerIndex == -1 {
		return NewHandleEventError(
			g.Id,
			len(g.Events),
			fmt.Sprintf("Player with id %s not found", e.PlayerId),
		)
	}

	player := &g.Players[playerIndex]

	cardIndex := -1
	for index, card := range player.Cards {
		if card.Id == e.DiscardedCardId {
			cardIndex = index
			break
		}
	}

	if cardIndex == -1 {
		return NewHandleEventError(
			g.Id,
			len(g.Events),
			fmt.Sprintf("Card with id %s not found", e.DiscardedCardId),
		)
	}

	if player.Cards[cardIndex].Type == CardTypeWildIndustry {
		g.WildIndustryCards = append(g.WildIndustryCards, player.Cards[cardIndex])
	} else if player.Cards[cardIndex].Type == CardTypeWildLocation {
		g.WildLocationCards = append(g.WildLocationCards, player.Cards[cardIndex])
	} else {
		player.Discards = append(player.Discards, player.Cards[cardIndex])
	}

	for i := cardIndex; i < len(player.Cards)-1; i++ {
		player.Cards[i] = player.Cards[i+1]
	}

	player.Cards = player.Cards[:len(player.Cards)-1]

	/** end boilerplate */

	player.Money += g.LoanAmount

	// TODO: Define income track in spec
	// TODO: Define minimum income level

	g.Events = append(g.Events, e)

	return nil
}

func calculateDeductedIncomeSpace(incomeTrack []int, currentIncomeSpace int, deductedLevels int) int {
	newIncomeSpace := -1
	incomeLevelsToDeduct := deductedLevels
	lastIncomeLevel := incomeTrack[currentIncomeSpace]

	for i := currentIncomeSpace - 1; i >= 0; i-- {
		if incomeTrack[i] == lastIncomeLevel {
			continue
		}

		lastIncomeLevel = incomeTrack[i]
		incomeLevelsToDeduct--
		if incomeLevelsToDeduct == 0 {
			newIncomeSpace = i
			break
		}
	}

	return newIncomeSpace
}
