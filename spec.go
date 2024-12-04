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
	ErrAsymmetricalEdge      = errors.New("Asymmetrical edge")
	ErrDuplicateLink         = errors.New("Duplicate link")
	ErrTooFewMerchantTiles   = errors.New("Too few merchant tiles")
	ErrTooManyMerchantTiles  = errors.New("Too many merchant tiles")
)

// TODO: Try generalizing the Spec interface to gracefully accomodate
// omissions due to player count and whatnot
type GameSpec struct {
	Name    string
	Version string

	CardSpecs         []CardSpec
	LocationSpecs     []LocationSpec
	MerchantTileSpecs []MerchantTileSpec
	PlayerMatSpec     PlayerMatSpec

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

	// Deck

	deck := []Card{}

	for _, cardSpec := range s.CardSpecs {
		if cardSpec.Type != CardTypeLocation && cardSpec.Type != CardTypeIndustry {
			return Game{}, fmt.Errorf("%w: %d", ErrInvalidCardType, playerCount)
		}

		// TODO: Just return a Card and a number from Build()?
		deck = append(deck, cardSpec.Build(playerCount)...)
	}

	if len(deck)%playerCount != 0 {
		return Game{}, fmt.Errorf("%w: %d %d", ErrIndivisibleDeckSize, len(deck), playerCount)
	}

	if len(deck) < playerCount*(s.HandSize+1) {
		return Game{}, fmt.Errorf("%w: %d < %d * (%d + 1)", ErrDeckTooSmall, len(deck), playerCount, s.HandSize)
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// Locations

	locations := make([]Location, len(s.LocationSpecs))
	for i, locationSpec := range s.LocationSpecs {
		locations[i] = locationSpec.Build(playerCount)
	}

	if err := validateLocations(locations); err != nil {
		return Game{}, err
	}

	// Merchants

	merchantTiles := []MerchantTile{}
	for _, merchantTileSpec := range s.MerchantTileSpecs {
		if tile := merchantTileSpec.Build(playerCount); tile != nil {
			merchantTiles = append(merchantTiles, *tile)
		}
	}

	rand.Shuffle(len(merchantTiles), func(i, j int) {
		merchantTiles[i], merchantTiles[j] = merchantTiles[j], merchantTiles[i]
	})

	for _, location := range locations {
		if location.Merchant == nil {
			continue
		}

		for i := 0; i < len(location.Merchant.Spaces); i++ {
			if len(merchantTiles) == 0 {
				return Game{}, ErrTooFewMerchantTiles
			}

			location.Merchant.Spaces[i].Tile = merchantTiles[0]
			merchantTiles = merchantTiles[1:]
		}
	}

	if len(merchantTiles) > 0 {
		return Game{}, fmt.Errorf("%w: %d additional tiles", ErrTooManyMerchantTiles, len(merchantTiles))
	}

	// Players

	players := []Player{}

	for range playerCount {
		player := Player{
			Id: uuid.NewString(),

			Mat:            s.PlayerMatSpec.Build(),
			Money:          s.StartingMoney,
			SpentMoney:     0,
			IncomeSpace:    s.StartingIncomeSpace,
			VictoryPoints:  0,
			RemainingLinks: s.LinksPerPlayer,

			Cards: make([]Card, 0, s.HandSize),
		}

		for range s.HandSize {
			player.Cards = append(player.Cards, deck[len(deck)-1])
			deck = deck[:len(deck)-1]
		}

		hiddenDiscard := deck[len(deck)-1]
		player.HiddenDiscard = &hiddenDiscard
		deck = deck[:len(deck)-1]

		players = append(players, player)
	}

	game := Game{}
	game.HandleGameCreatedEvent(GameCreatedEvent{
		Id:   uuid.NewString(),
		Type: EventTypeGameCreated,

		GameSpecName:    s.Name,
		GameSpecVersion: s.Version,

		Deck:      deck,
		Locations: locations,
		Players:   players,

		NumWildLocationCards: s.NumWildLocationCards,
		NumWildIndustryCards: s.NumWildIndustryCards,

		InitialCoalInMarket: s.InitialCoalInMarket,
		InitialIronInMarket: s.InitialIronInMarket,
	})

	return game, nil
}

func validateLocations(locations []Location) error {
	canalEraNeighbours := map[string]map[string]bool{}
	railEraNeighbours := map[string]map[string]bool{}

	for _, location := range locations {
		if _, ok := canalEraNeighbours[location.Name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateLocationName, location.Name)
		}

		if _, ok := railEraNeighbours[location.Name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateLocationName, location.Name)
		}

		canalEraNeighbours[location.Name] = map[string]bool{}
		for _, neighbour := range location.CanalEraNeighbours {
			if canalEraNeighbours[location.Name][neighbour] {
				return fmt.Errorf("%w: %s -> %s", ErrDuplicateLink, location.Name, neighbour)
			}
			canalEraNeighbours[location.Name][neighbour] = true
		}

		railEraNeighbours[location.Name] = map[string]bool{}
		for _, neighbour := range location.RailEraNeighbours {
			if railEraNeighbours[location.Name][neighbour] {
				return fmt.Errorf("%w: %s -> %s", ErrDuplicateLink, location.Name, neighbour)
			}
			railEraNeighbours[location.Name][neighbour] = true
		}
	}

	for _, location := range locations {
		for _, neighbour := range location.CanalEraNeighbours {
			neighbourLookup, ok := canalEraNeighbours[neighbour]

			if !ok {
				return fmt.Errorf("%w: %s", ErrNonExistentLocation, neighbour)
			}

			if _, ok := neighbourLookup[location.Name]; !ok {
				return fmt.Errorf("%w: %s -> %s", ErrAsymmetricalEdge, location.Name, neighbour)
			}
		}

		for _, neighbour := range location.RailEraNeighbours {
			neighbourLookup, ok := railEraNeighbours[neighbour]

			if !ok {
				return fmt.Errorf("%w: %s", ErrNonExistentLocation, neighbour)
			}

			if _, ok := neighbourLookup[location.Name]; !ok {
				return fmt.Errorf("%w: %s -> %s", ErrAsymmetricalEdge, location.Name, neighbour)
			}
		}
	}

	return nil
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
}
