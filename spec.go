package blic

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPlayerCount = errors.New("Invalid player count")
)

type GameSpec struct {
	Name string

	CardSpecs         []CardSpec
	LocationSpecs     []LocationSpec
	MerchantTileSpecs []MerchantTileSpec

	MinPlayerCount int
	MaxPlayerCount int

	InitialCoalInMarket int
	InitialIronInMarket int
}

func (s *GameSpec) Build(playerCount int) (Game, error) {
	if playerCount < s.MinPlayerCount || playerCount > s.MaxPlayerCount {
		return Game{}, fmt.Errorf("%w: %d", ErrInvalidPlayerCount, playerCount)
	}

	// TODO: stub
	// TODO: Create event and pass it to an empty Game{}
	/*
		game := Game{}
		game.HandleGameCreatedEvent(GameCreatedEvent{
			uuid.NewString(),
			game.Deck,
			game.Locations,
			game.Players,
		})
	*/

	return Game{}, nil
}

type CardSpec struct {
	IndustryTypes []IndustryType
	LocationName  string

	IsWildLocation bool
	IsWildIndustry bool

	AmountByPlayerCount map[int]int
}

func (s *CardSpec) Build(playerCount int) []Card {
	cards := []Card{}

	for i := 0; i < s.AmountByPlayerCount[playerCount]; i++ {
		cards = append(cards, Card{
			// NOTE: Depending on how many games are expected to be processed
			// simultaneously, it may be worthwhile not cloning resources
			// that are in principle expected to be immutable.
			IndustryTypes:  CloneSlice(s.IndustryTypes),
			LocationName:   s.LocationName,
			IsWildLocation: s.IsWildLocation,
			IsWildIndustry: s.IsWildIndustry,
		})
	}

	return cards
}

type LocationSpec struct {
	Name string

	CanalEraNeighbours []string
	RailEraNeighbours  []string

	IndustrySpaces []IndustrySpaceSpec
	Merchant       *MerchantSpec
}

func (s *LocationSpec) Build(playerCount int) Location {
	location := Location{
		Name:               s.Name,
		CanalEraNeighbours: make([]string, len(s.CanalEraNeighbours)),
		RailEraNeighbours:  make([]string, len(s.RailEraNeighbours)),
	}

	copy(location.CanalEraNeighbours, s.CanalEraNeighbours)
	copy(location.RailEraNeighbours, s.RailEraNeighbours)

	for _, industrySpaceSpec := range s.IndustrySpaces {
		location.IndustrySpaces = append(location.IndustrySpaces, industrySpaceSpec.Build(playerCount))
	}

	if s.Merchant != nil {
		location.Merchant = s.Merchant.Build(playerCount)
	}

	return location
}

type IndustrySpaceSpec struct {
	Types     []IndustryType
	Tile      IndustryTileSpec
	Resources int
}

func (s *IndustrySpaceSpec) Build(playerCount int) IndustrySpace {
	return IndustrySpace{
		Types:     CloneSlice(s.Types),
		Tile:      s.Tile.Build(),
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
