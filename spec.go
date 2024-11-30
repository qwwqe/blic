package blic

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPlayerCount  = errors.New("Invalid player count")
	ErrInvalidCardType     = errors.New("Invalid card type")
	ErrNonExistentLocation = errors.New("Nonexistent location")
)

// TODO: Try generalizing the Spec interface to gracefully accomodate
// omissions due to player count and whatnot
type GameSpec struct {
	Name string

	CardSpecs         []CardSpec
	LocationSpecs     []LocationSpec
	MerchantTileSpecs []MerchantTileSpec

	NumWildLocationCards int
	NumWildIndustryCards int

	MinPlayerCount int
	MaxPlayerCount int

	InitialCoalInMarket int
	InitialIronInMarket int
}

func (s *GameSpec) Build(playerCount int) (Game, error) {
	if playerCount < s.MinPlayerCount || playerCount > s.MaxPlayerCount {
		return Game{}, fmt.Errorf("%w: %d", ErrInvalidPlayerCount, playerCount)
	}

	deck := []Card{}

	for _, cardSpec := range s.CardSpecs {
		if cardSpec.Type != CardTypeLocation && cardSpec.Type != CardTypeIndustry {
			return Game{}, fmt.Errorf("%w: %d", ErrInvalidCardType, playerCount)
		}

		// TODO: Just return a Card and a number from Build()?
		deck = append(deck, cardSpec.Build(playerCount)...)
	}

	// TODO: Factor out validation into its own testable function
	locations := make([]Location, len(s.LocationSpecs))
	locationLookup := map[string]bool{}
	for i, locationSpec := range s.LocationSpecs {
		locations[i] = locationSpec.Build(playerCount)
		locationLookup[locations[i].Name] = true
	}

	// TODO: Check for bidirectional edges?
	for _, location := range locations {
		for _, neighbour := range location.CanalEraNeighbours {
			if !locationLookup[neighbour] {
				return Game{}, fmt.Errorf("%w: %s", ErrNonExistentLocation, neighbour)
			}
		}

		for _, neighbour := range location.RailEraNeighbours {
			if !locationLookup[neighbour] {
				return Game{}, fmt.Errorf("%w: %s", ErrNonExistentLocation, neighbour)
			}
		}
	}

	merchantTiles := []MerchantTile{}
	for _, merchantTileSpec := range s.MerchantTileSpecs {
		if tile := merchantTileSpec.Build(playerCount); tile != nil {
			merchantTiles = append(merchantTiles, *tile)
		}
	}

	//

	game := Game{}
	game.HandleGameCreatedEvent(GameCreatedEvent{
		uuid.NewString(),
		game.Deck,
		game.Locations,
		game.Players,
	})

	return game, nil
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
