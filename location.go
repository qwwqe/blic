package main

type IndustrySpace struct {
	Types     []IndustryType
	Tile      IndustryTile
	Resources int
}

func (s IndustrySpace) Clone() IndustrySpace {
	s.Types = CloneSlice(s.Types)
	s.Tile = s.Tile.Clone()

	return s
}

type MerchantBeerBonusType string

const (
	MerchantBeerBonusTypeDevelop       MerchantBeerBonusType = "develop"
	MerchantBeerBonusTypeIncomeBoost   MerchantBeerBonusType = "incomeboost"
	MerchantBeerBonusTypeVictoryPoints MerchantBeerBonusType = "victorypoints"
	MerchantBeerBonusTypeMoney         MerchantBeerBonusType = "money"
)

type MerchantBeerBonus struct {
	Type   MerchantBeerBonusType
	Amount int
}

type MerchantTile struct {
	IndustryTypes []IndustryType
}

func (t MerchantTile) Clone() MerchantTile {
	t.IndustryTypes = CloneSlice(t.IndustryTypes)
	return t
}

type MerchantSpace struct {
	Tile MerchantTile
	Beer int
}

func (s MerchantSpace) Clone() MerchantSpace {
	s.Tile = s.Tile.Clone()
	return s
}

type Merchant struct {
	Links     int
	BeerBonus MerchantBeerBonus
	Spaces    []*MerchantSpace
}

func (m Merchant) Clone() Merchant {
	m.Spaces = ClonePointerSlice(m.Spaces)
	return m
}

type Location struct {
	Name               string
	CanalEraNeighbours []*Location
	RailEraNeighbours  []*Location

	IndustrySpaces []IndustrySpace
	Merchant       *Merchant
}

func (l Location) Clone() Location {
	l.CanalEraNeighbours = ClonePointerSlice(l.CanalEraNeighbours)
	l.RailEraNeighbours = ClonePointerSlice(l.RailEraNeighbours)
	l.IndustrySpaces = CloneSlice(l.IndustrySpaces)
	if l.Merchant != nil {
		m := l.Merchant.Clone()
		l.Merchant = &m
	}
	return l
}
