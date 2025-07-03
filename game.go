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
	ErrInvalidPhaseAction        = errors.New("Action taken outside of action phase")
	ErrOutOfTurn                 = errors.New("Action taken out of turn")
	ErrActionPlayerNotFound      = errors.New("Action player not found")
	ErrActionDiscardNotFound     = errors.New("Action discard not found")
	ErrNoRemainingActions        = errors.New("No remaining actions")
	ErrRemainingActions          = errors.New("Actions still remaining")
	ErrInsufficientLoanCredit    = errors.New("Insufficient loan credit")
	ErrNoRemainingWildCards      = errors.New("No remaining wild cards")
	ErrScoutDiscardAmount        = errors.New("Not enough discards to scout")
	ErrScoutWildCard             = errors.New("Cannot scout with wild card in hand")
	ErrNoDevelopableIndustry     = errors.New("No industries can be developed")
	ErrCannotConsumeIron         = errors.New("Cannot consume iron")
	ErrInvalidIndustryTypeAmount = errors.New("Invalid amount of industry types")
)

type GamePhase string

const (
	GamePhaseAction                  GamePhase = "action"
	GamePhasePickIndustriesToDevelop GamePhase = "pickindustriestodevelop"
	GamePhaseConsumeIronToDevelop    GamePhase = "consumeirontodevelop"
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

	CoalMarket Market
	IronMarket Market

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

	g.CoalMarket = e.CoalMarket
	g.IronMarket = e.IronMarket

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
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if player.RemainingActions == 0 {
		return ErrNoRemainingActions
	}

	if _, err := getEventCardIndex(g, player, discardedCardId); err != nil {
		return ErrActionDiscardNotFound
	}

	newIncomeSpace := calculateDeductedIncomeSpace(g.IncomeTrack, player.IncomeSpace, g.LoanIncomeLevelPenalty)
	if newIncomeSpace < 0 {
		return ErrInsufficientLoanCredit
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
			fmt.Sprintf("Negative income space index: %d", newIncomeSpace),
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

func (g *Game) TakePassAction(playerId, discardedCardId string) error {
	if g.Phase != GamePhaseAction {
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if player.RemainingActions == 0 {
		return ErrNoRemainingActions
	}

	if _, err := getEventCardIndex(g, player, discardedCardId); err != nil {
		return ErrActionDiscardNotFound
	}

	event := PassActionTakenEvent{
		Type:            PassActionTakenEventType,
		PlayerId:        playerId,
		DiscardedCardId: discardedCardId,
	}

	if err := g.handlePassActionTakenEvent(event); err != nil {
		return err
	}

	return nil
}

func (g *Game) handlePassActionTakenEvent(e PassActionTakenEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	cardIndex, err := getEventCardIndex(g, player, e.DiscardedCardId)
	if err != nil {
		return err
	}

	processEventDiscard(g, player, cardIndex)

	player.RemainingActions--

	g.Events = append(g.Events, e)

	return nil
}

func (g *Game) TakeScoutAction(playerId string, discardedCardIds []string) error {
	if g.Phase != GamePhaseAction {
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if player.RemainingActions == 0 {
		return ErrNoRemainingActions
	}

	if len(discardedCardIds) != 3 {
		return ErrScoutDiscardAmount
	}

	if len(g.WildIndustryCards) == 0 || len(g.WildLocationCards) == 0 {
		return ErrNoRemainingWildCards
	}

	for _, discardedCardId := range discardedCardIds {
		if _, err := getEventCardIndex(g, player, discardedCardId); err != nil {
			return ErrActionDiscardNotFound
		}
	}

	for _, card := range player.Cards {
		if card.Type == CardTypeWildIndustry || card.Type == CardTypeWildLocation {
			return ErrScoutWildCard
		}
	}

	event := ScoutActionTakenEvent{
		Type:             ScoutActionTakenEventType,
		PlayerId:         playerId,
		DiscardedCardIds: discardedCardIds,
	}

	if err := g.handleScoutActionTakenEvent(event); err != nil {
		return err
	}

	return nil
}

func (g *Game) handleScoutActionTakenEvent(e ScoutActionTakenEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	for _, discardedCardId := range e.DiscardedCardIds {
		cardIndex, err := getEventCardIndex(g, player, discardedCardId)
		if err != nil {
			return err
		}

		processEventDiscard(g, player, cardIndex)
	}

	if len(g.WildIndustryCards) == 0 || len(g.WildLocationCards) == 0 {
		return NewHandleEventError(g.Id, len(g.Events), "No remaining wild cards")
	}

	player.Cards = append(
		player.Cards,
		g.WildIndustryCards[len(g.WildIndustryCards)-1],
		g.WildLocationCards[len(g.WildLocationCards)-1],
	)
	g.WildIndustryCards = g.WildIndustryCards[:len(g.WildIndustryCards)-1]
	g.WildLocationCards = g.WildLocationCards[:len(g.WildLocationCards)-1]

	player.RemainingActions--

	g.Events = append(g.Events, e)

	return nil
}

func (g *Game) TakeDevelopAction(playerId, discardedCardId string) error {
	if g.Phase != GamePhaseAction {
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if player.RemainingActions == 0 {
		return ErrNoRemainingActions
	}

	if _, err := getEventCardIndex(g, player, discardedCardId); err != nil {
		return ErrActionDiscardNotFound
	}

	if !player.Mat.HasDevelopableIndustry() {
		return ErrNoDevelopableIndustry
	}

	if !canConsumeIron(g, player, 1) {
		return ErrCannotConsumeIron
	}

	event := DevelopActionTakenEvent{
		Type:            DevelopActionTakenEventType,
		PlayerId:        playerId,
		DiscardedCardId: discardedCardId,
	}

	if err := g.handleDevelopActionTakenEvent(event); err != nil {
		return err
	}

	return nil
}

func (g *Game) handleDevelopActionTakenEvent(e DevelopActionTakenEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	cardIndex, err := getEventCardIndex(g, player, e.DiscardedCardId)
	if err != nil {
		return err
	}

	processEventDiscard(g, player, cardIndex)

	g.Phase = GamePhasePickIndustriesToDevelop

	// player.RemainingActions--

	g.Events = append(g.Events, e)

	return nil
}

func (g *Game) PickIndustriesToDevelop(playerId string, industryTypes []IndustryType) error {
	if g.Phase != GamePhasePickIndustriesToDevelop {
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if len(industryTypes) < 1 || len(industryTypes) > 2 {
		return ErrInvalidIndustryTypeAmount
	}

	if !player.Mat.CanDevelopIndustries(industryTypes) {
		return ErrNoDevelopableIndustry
	}

	if !canConsumeIron(g, player, len(industryTypes)) {
		return ErrCannotConsumeIron
	}

	event := IndustriesPickedToDevelopEvent{
		Type:          IndustriesPickedToDevelopEventType,
		PlayerId:      playerId,
		IndustryTypes: industryTypes,
	}

	if err := g.handleIndustriesPickedToDevelopEvent(event); err != nil {
		return err
	}

	return nil
}

func (g *Game) handleIndustriesPickedToDevelopEvent(e IndustriesPickedToDevelopEvent) error {
	player, err := getEventPlayer(g, e.PlayerId)
	if err != nil {
		return err
	}

	for _, industryType := range e.IndustryTypes {
		if ok := player.Mat.Develop(industryType); !ok {
			return NewHandleEventError(
				g.Id,
				len(g.Events),
				fmt.Sprintf("Player %s unable to develop industry type \"%s\"", player.Id, industryType),
			)
		}
	}

	g.Phase = GamePhaseConsumeIronToDevelop // 狀態怎麼傳給下一個階段？

	g.Events = append(g.Events, e)

	return nil
}

func (g *Game) EndTurn(playerId string) error {
	if g.Phase != GamePhaseAction {
		return ErrInvalidPhaseAction
	}

	if g.Players[g.PlayerIndex].Id != playerId {
		return ErrOutOfTurn
	}

	player, err := getEventPlayer(g, playerId)
	if err != nil {
		return ErrActionPlayerNotFound
	}

	if player.RemainingActions != 0 {
		return ErrRemainingActions
	}

	event := TurnEndedEvent{
		Type:     TurnEndedEventType,
		PlayerId: player.Id,
	}

	if err = g.handleTurnEndedEvent(event); err != nil {
		return err
	}

	// TODO: End of round
	// TODO: End of era

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

func (g *Game) ConsumableIronOnIndustries() int {
	resources := 0
	for _, location := range g.Locations {
		for _, space := range location.IndustrySpaces {
			if space.Tile.Type == IndustryTypeIronWorks {
				resources += space.Resources
			}
		}
	}
	return resources
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
	switch player.Cards[cardIndex].Type {
	case CardTypeIndustry:
		game.WildIndustryCards = append(game.WildIndustryCards, player.Cards[cardIndex])
	case CardTypeWildLocation:
		game.WildLocationCards = append(game.WildLocationCards, player.Cards[cardIndex])
	default:
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
