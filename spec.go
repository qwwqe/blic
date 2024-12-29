package blic

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

var (
	ErrInvalidPlayerCount    = errors.New("Invalid player count")
	ErrInvalidCardType       = errors.New("Invalid card type")
	ErrIndivisibleDeckSize   = errors.New("Deck size not divisible by player count")
	ErrDeckTooSmall          = errors.New("Deck too small to deal a full hand to every player")
	ErrNonExistentLocation   = errors.New("Nonexistent location")
	ErrDuplicateLocationName = errors.New("Duplicate location name")
	ErrTooFewMerchantTiles   = errors.New("Too few merchant tiles")
	ErrTooManyMerchantTiles  = errors.New("Too many merchant tiles")
	ErrIncorrectIndustryType = errors.New("Incorrect industry type")
)

// TODO: Try generalizing the Spec interface to gracefully accomodate
// omissions due to player count and whatnot
type GameSpec struct {
	Name    string
	Version string

	CardSpecs               []CardSpec
	LocationSpecs           []LocationSpec
	CanalEraConnectionSpecs []ConnectionSpec
	RailEraConnectionSpecs  []ConnectionSpec
	MerchantTileSpecs       []MerchantTileSpec
	PlayerMatSpec           PlayerMatSpec

	NumWildLocationCards int
	NumWildIndustryCards int

	MinPlayerCount int
	MaxPlayerCount int

	InitialCoalInMarket int
	InitialIronInMarket int

	StartingMoney       int
	StartingIncomeSpace int
	HandSize            int
	LinksPerPlayer      int
}

func (s *GameSpec) Build(playerCount int) (Game, error) {
	if playerCount < s.MinPlayerCount || playerCount > s.MaxPlayerCount {
		return Game{}, fmt.Errorf("%w: %d", ErrInvalidPlayerCount, playerCount)
	}

	deck, err := buildDeck(*s, playerCount)
	if err != nil {
		return Game{}, err
	}

	locations, err := buildLocations(*s, playerCount)
	if err != nil {
		return Game{}, err
	}

	canalEraConnections, err := buildCanalEraConnections(*s, locations)
	if err != nil {
		return Game{}, err
	}

	railEraConnections, err := buildRailEraConnections(*s, locations)
	if err != nil {
		return Game{}, err
	}

	// TODO: Return merchant tiles and define a separate function for their "placement"
	if err := buildMerchants(*s, playerCount, locations); err != nil {
		return Game{}, err
	}

	// TODO: Define a separate function for drawing of cards from the deck?
	players, err := buildPlayers(*s, playerCount, &deck)
	if err != nil {
		return Game{}, err
	}

	game := Game{}
	game.HandleGameCreatedEvent(GameCreatedEvent{
		Id:   uuid.NewString(),
		Type: EventTypeGameCreated,

		GameSpecName:    s.Name,
		GameSpecVersion: s.Version,

		Deck:                deck,
		Locations:           locations,
		CanalEraConnections: canalEraConnections,
		RailEraConnections:  railEraConnections,
		Players:             players,

		NumWildLocationCards: s.NumWildLocationCards,
		NumWildIndustryCards: s.NumWildIndustryCards,

		InitialCoalInMarket: s.InitialCoalInMarket,
		InitialIronInMarket: s.InitialIronInMarket,
	})

	return game, nil
}

func validateConnections(locations []Location, connections []Connection) error {
	locationLookup := map[string]bool{}
	for _, location := range locations {
		locationLookup[location.Name] = true
	}

	for _, connection := range connections {
		for _, locationName := range connection.LocationNames {
			if !locationLookup[locationName] {
				return fmt.Errorf("%w: %s", ErrNonExistentLocation, locationName)
			}
		}
	}

	return nil
}

func validatePlayerMat(playerMat PlayerMat) error {
	for _, tile := range playerMat.BreweryTiles {
		if tile.Type != IndustryTypeBrewery {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypeBrewery, tile.Type)
		}
	}
	for _, tile := range playerMat.CoalMineTiles {
		if tile.Type != IndustryTypeCoalMine {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypeCoalMine, tile.Type)
		}
	}
	for _, tile := range playerMat.CottonMillTiles {
		if tile.Type != IndustryTypeCottonMill {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypeCottonMill, tile.Type)
		}
	}
	for _, tile := range playerMat.IronWorksTiles {
		if tile.Type != IndustryTypeIronWorks {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypeIronWorks, tile.Type)
		}
	}
	for _, tile := range playerMat.ManufacturerTiles {
		if tile.Type != IndustryTypeManufacturer {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypeManufacturer, tile.Type)
		}
	}
	for _, tile := range playerMat.PotteryTiles {
		if tile.Type != IndustryTypePottery {
			return fmt.Errorf("%w: expected %s but got %s", ErrIncorrectIndustryType, IndustryTypePottery, tile.Type)
		}
	}

	return nil
}

func buildDeck(spec GameSpec, numPlayers int) ([]Card, error) {
	deck := []Card{}

	for _, cardSpec := range spec.CardSpecs {
		if cardSpec.Type != CardTypeLocation && cardSpec.Type != CardTypeIndustry {
			return nil, fmt.Errorf("%w: %d", ErrInvalidCardType, numPlayers)
		}

		// TODO: Just return a Card and a number from Build()?
		deck = append(deck, cardSpec.Build(numPlayers)...)
	}

	if len(deck)%numPlayers != 0 {
		return nil, fmt.Errorf("%w: %d %d", ErrIndivisibleDeckSize, len(deck), numPlayers)
	}

	if len(deck) < numPlayers*(spec.HandSize+1) {
		return nil, fmt.Errorf("%w: %d < %d * (%d + 1)", ErrDeckTooSmall, len(deck), numPlayers, spec.HandSize)
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck, nil
}

func buildLocations(spec GameSpec, numPlayers int) ([]Location, error) {
	locations := make([]Location, len(spec.LocationSpecs))
	for i, locationSpec := range spec.LocationSpecs {
		locations[i] = locationSpec.Build(numPlayers)
	}

	locationLookup := map[string]bool{}
	for _, location := range locations {
		if locationLookup[location.Name] {
			return nil, fmt.Errorf("%w: %s", ErrDuplicateLocationName, location.Name)
		}
		locationLookup[location.Name] = true
	}

	return locations, nil
}

func buildCanalEraConnections(spec GameSpec, locations []Location) ([]Connection, error) {
	canalEraConnections := make([]Connection, len(spec.CanalEraConnectionSpecs))
	for i, connectionSpec := range spec.CanalEraConnectionSpecs {
		canalEraConnections[i] = connectionSpec.Build()
	}

	if err := validateConnections(locations, canalEraConnections); err != nil {
		return nil, err
	}

	return canalEraConnections, nil
}

func buildRailEraConnections(spec GameSpec, locations []Location) ([]Connection, error) {
	railEraConnections := make([]Connection, len(spec.RailEraConnectionSpecs))
	for i, connectionSpec := range spec.RailEraConnectionSpecs {
		railEraConnections[i] = connectionSpec.Build()
	}

	if err := validateConnections(locations, railEraConnections); err != nil {
		return nil, err
	}

	return railEraConnections, nil
}

func buildMerchants(spec GameSpec, numPlayers int, locations []Location) error {
	merchantTiles := []MerchantTile{}
	for _, merchantTileSpec := range spec.MerchantTileSpecs {
		if tile := merchantTileSpec.Build(numPlayers); tile != nil {
			merchantTiles = append(merchantTiles, *tile)
		}
	}

	rand.Shuffle(len(merchantTiles), func(i, j int) {
		merchantTiles[i], merchantTiles[j] = merchantTiles[j], merchantTiles[i]
	})

	count := 0
	for _, l := range locations {
		if l.Merchant != nil {
			count += len(l.Merchant.Spaces)
		}
	}

	for i := range len(locations) {
		if locations[i].Merchant == nil {
			continue
		}

		for j := 0; j < len(locations[i].Merchant.Spaces); j++ {
			if len(merchantTiles) == 0 {
				fmt.Println(locations[i].Merchant.Spaces)
				return ErrTooFewMerchantTiles
			}

			locations[i].Merchant.Spaces[j].Tile = merchantTiles[0]
			merchantTiles = merchantTiles[1:]
		}
	}

	if len(merchantTiles) > 0 {
		return fmt.Errorf("%w: %d additional tiles", ErrTooManyMerchantTiles, len(merchantTiles))
	}

	return nil
}

func buildPlayers(spec GameSpec, numPlayers int, deck *[]Card) ([]Player, error) {
	players := []Player{}

	for range numPlayers {
		player := Player{
			Id: uuid.NewString(),

			Mat:            spec.PlayerMatSpec.Build(),
			Money:          spec.StartingMoney,
			SpentMoney:     0,
			IncomeSpace:    spec.StartingIncomeSpace,
			VictoryPoints:  0,
			RemainingLinks: spec.LinksPerPlayer,

			Cards: make([]Card, 0, spec.HandSize),
		}

		for range spec.HandSize {
			player.Cards = append(player.Cards, (*deck)[len(*deck)-1])
			*deck = (*deck)[:len(*deck)-1]
		}

		hiddenDiscard := (*deck)[len(*deck)-1]
		player.HiddenDiscard = &hiddenDiscard
		*deck = (*deck)[:len(*deck)-1]

		if err := validatePlayerMat(player.Mat); err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

type CardSpec struct {
	IndustryTypes []IndustryType
	LocationName  string
	Type          CardType

	AmountByPlayerCount map[int]int
}

func (s *CardSpec) Build(playerCount int) []Card {
	cards := []Card{}

	for i := 0; i < s.AmountByPlayerCount[playerCount]; i++ {
		cards = append(cards, Card{
			// NOTE: Depending on how many games are expected to be processed
			// simultaneously, it may be worthwhile not cloning resources
			// that are in principle expected to be immutable.
			IndustryTypes: CloneSlice(s.IndustryTypes),
			LocationName:  s.LocationName,
			Type:          s.Type,
		})
	}

	return cards
}

type LocationSpec struct {
	Name string

	IndustrySpaces []IndustrySpaceSpec
	Merchant       *MerchantSpec
}

func (s *LocationSpec) Build(playerCount int) Location {
	location := Location{
		Name: s.Name,
	}

	for _, industrySpaceSpec := range s.IndustrySpaces {
		location.IndustrySpaces = append(location.IndustrySpaces, industrySpaceSpec.Build(playerCount))
	}

	if s.Merchant != nil {
		location.Merchant = s.Merchant.Build(playerCount)
	}

	return location
}

type ConnectionSpec struct {
	LocationNames []string
}

func (s *ConnectionSpec) Build() Connection {
	c := Connection{
		LocationNames: make([]string, len(s.LocationNames)),
	}
	copy(c.LocationNames, s.LocationNames)

	return c
}

type IndustrySpaceSpec struct {
	Types []IndustryType
	// Tile      IndustryTileSpec
	Resources int
}

func (s *IndustrySpaceSpec) Build(playerCount int) IndustrySpace {
	return IndustrySpace{
		Types: CloneSlice(s.Types),
		// Tile:      s.Tile.Build(),
		Resources: s.Resources,
	}
}

type IndustryTileSpec struct {
	IndustryTile
}

func (s *IndustryTileSpec) Build() IndustryTile {
	return s.Clone()
}

type MerchantSpec struct {
	BeerBonus  MerchantBeerBonusSpec
	NumLinks   int
	NumSpaces  int
	MinPlayers int
}

func (s *MerchantSpec) Build(playerCount int) *Merchant {
	merchant := Merchant{
		NumLinks:  s.NumLinks,
		BeerBonus: s.BeerBonus.Build(),
	}

	if playerCount >= s.MinPlayers {
		for i := 0; i < s.NumSpaces; i++ {
			merchant.Spaces = append(merchant.Spaces, MerchantSpace{})
		}
	}

	return &merchant
}

type MerchantBeerBonusSpec struct {
	MerchantBeerBonus
}

func (s *MerchantBeerBonusSpec) Build() MerchantBeerBonus {
	return s.Clone()
}

type MerchantTileSpec struct {
	MerchantTile
	MinPlayers int
}

func (s *MerchantTileSpec) Build(playerCount int) *MerchantTile {
	if playerCount < s.MinPlayers {
		return nil
	}

	tile := s.MerchantTile.Clone()
	return &tile
}

type PlayerMatSpec struct {
	CoalMineTiles  []IndustryTileSpec
	IronWorksTiles []IndustryTileSpec
	BreweryTiles   []IndustryTileSpec

	ManufacturerTiles []IndustryTileSpec
	CottonMillTiles   []IndustryTileSpec
	PotteryTiles      []IndustryTileSpec
}

func (s *PlayerMatSpec) Build() PlayerMat {
	mat := PlayerMat{
		CoalMineTiles:  make([]IndustryTile, 0, len(s.CoalMineTiles)),
		IronWorksTiles: make([]IndustryTile, 0, len(s.IronWorksTiles)),
		BreweryTiles:   make([]IndustryTile, 0, len(s.BreweryTiles)),

		ManufacturerTiles: make([]IndustryTile, 0, len(s.ManufacturerTiles)),
		CottonMillTiles:   make([]IndustryTile, 0, len(s.CottonMillTiles)),
		PotteryTiles:      make([]IndustryTile, 0, len(s.PotteryTiles)),
	}

	for _, t := range s.CoalMineTiles {
		mat.CoalMineTiles = append(mat.CoalMineTiles, t.Build())
	}
	for _, t := range s.IronWorksTiles {
		mat.IronWorksTiles = append(mat.IronWorksTiles, t.Build())
	}
	for _, t := range s.BreweryTiles {
		mat.BreweryTiles = append(mat.BreweryTiles, t.Build())
	}

	for _, t := range s.ManufacturerTiles {
		mat.ManufacturerTiles = append(mat.ManufacturerTiles, t.Build())
	}
	for _, t := range s.CottonMillTiles {
		mat.CottonMillTiles = append(mat.CottonMillTiles, t.Build())
	}
	for _, t := range s.PotteryTiles {
		mat.PotteryTiles = append(mat.PotteryTiles, t.Build())
	}

	return mat
}
