package blic

import (
	"errors"
	"fmt"
)

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

var (
	InvalidPhaseActionErr    = errors.New("Action taken outside of action phase")
	OutOfTurnErr             = errors.New("Action taken out of turn")
	ActionPlayerNotFoundErr  = errors.New("Action player not found")
	ActionDiscardNotFoundErr = errors.New("Action discard not found")
	NoRemainingActionsErr    = errors.New("No remaining actions")
)

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
	LoanAmount             int
	HandSize               int
	LoanIncomeLevelPenalty int
	ActionsPerTurn         int

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

	g.LoanIncomeLevelPenalty = e.LoanIncomeLevelPenalty

	g.ActionsPerTurn = e.ActionsPerTurn

	g.WildLocationCards = make([]Card, 0, e.NumWildLocationCards)
	for range e.NumWildLocationCards {
		g.WildLocationCards = append(g.WildLocationCards, Card{Type: CardTypeWildLocation})
	}

	g.WildIndustryCards = make([]Card, 0, e.NumWildIndustryCards)
	for range e.NumWildIndustryCards {
		g.WildIndustryCards = append(g.WildIndustryCards, Card{Type: CardTypeWildIndustry})
	}

	g.PlayerIndex = 0
	g.Round = 0
	g.Phase = GamePhaseAction

	return g
}

// TODO: Tests
func (g *Game) TakeLoanAction(playerId, discardedCardId string) error {
	if g.Phase != GamePhaseAction {
		return InvalidPhaseActionErr
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return OutOfTurnErr
	}

	// TODO: Handle action sub-phases (choices within each phase)

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ActionPlayerNotFoundErr
	}

	if player.RemainingActions == 0 {
		return NoRemainingActionsErr
	}

	if _, err := getEventCardIndex(g, player, discardedCardId); err != nil {
		return ActionDiscardNotFoundErr
	}

	event := LoanActionTakenEvent{
		Type:            LoanActionTakenEventType,
		PlayerId:        playerId,
		DiscardedCardId: discardedCardId,
	}

	if err := g.handleLoanActionTakenEvent(event); err != nil {
		return err
	}

	/** post-action boilerplate */

	if player.RemainingActions == 0 {
		event := TurnEndedEvent{
			Type:     TurnEndedEventType,
			PlayerId: player.Id,
		}

		if err = g.handleTurnEndedEvent(event); err != nil {
			return err
		}

		// TODO: End of round

		// TODO: End of era
	}

	/** end boilerplate*/

	return nil
}

// TODO: Tests
func (g *Game) handleLoanActionTakenEvent(e LoanActionTakenEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	cardIndex, err := getEventCardIndex(g, player, e.DiscardedCardId)
	if err != nil {
		return err
	}

	newIncomeSpace := calculateDeductedIncomeSpace(g.IncomeTrack, player.IncomeSpace, g.LoanIncomeLevelPenalty)
	if newIncomeSpace < 0 {
		return NewHandleEventError(
			g.Id,
			len(g.Events),
			fmt.Sprintf("Negative income space index"),
		)
	}

	processEventDiscard(g, player, cardIndex)

	player.IncomeSpace = newIncomeSpace
	player.Money += g.LoanAmount

	/** post-action boilerplate */

	player.RemainingActions--

	/** end boilerplate */

	g.Events = append(g.Events, e)

	return nil
}

func (g *Game) handleTurnEndedEvent(e TurnEndedEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	for len(g.Deck) > 0 && len(player.Cards) < g.HandSize {
		player.Cards = append(player.Cards, g.Deck[len(g.Deck)-1])
		g.Deck = g.Deck[:len(g.Deck)-1]
	}

	player.RemainingActions = g.ActionsPerTurn

	g.Events = append(g.Events, e)

	return nil
}

func getEventPlayer(g *Game, playerId string) (*Player, error) {
	playerIndex := -1
	for index, player := range g.Players {
		if player.Id == playerId {
			playerIndex = index
			break
		}
	}

	if playerIndex == -1 {
		return nil, NewHandleEventError(
			g.Id,
			len(g.Events),
			fmt.Sprintf("Player with id %s not found", playerId),
		)
	}

	return &g.Players[playerIndex], nil
}

func getEventCardIndex(game *Game, player *Player, cardId string) (int, error) {
	cardIndex := -1
	for index, card := range player.Cards {
		if card.Id == cardId {
			cardIndex = index
			break
		}
	}

	if cardIndex == -1 {
		return -1, NewHandleEventError(
			game.Id,
			len(game.Events),
			fmt.Sprintf("Card with id %s not found", cardId),
		)
	}

	return cardIndex, nil
}

func processEventDiscard(game *Game, player *Player, cardIndex int) {
	if player.Cards[cardIndex].Type == CardTypeWildIndustry {
		game.WildIndustryCards = append(game.WildIndustryCards, player.Cards[cardIndex])
	} else if player.Cards[cardIndex].Type == CardTypeWildLocation {
		game.WildLocationCards = append(game.WildLocationCards, player.Cards[cardIndex])
	} else {
		player.Discards = append(player.Discards, player.Cards[cardIndex])
	}

	for i := cardIndex; i < len(player.Cards)-1; i++ {
		player.Cards[i] = player.Cards[i+1]
	}

	player.Cards = player.Cards[:len(player.Cards)-1]
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
