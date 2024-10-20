package main

type IndustrySpace struct {
	Types     []IndustryType
	Tile      IndustryTile
	Resources int
}

type MerchantBeerBonusType string

const (
	MerchantBeerBonusTypeDevelop       MerchantBeerBonusType = "develop"
	MerchantBeerBonusTypeIncomeBoost   MerchantBeerBonusType = "incomeboost"
	MerchantBeerBonusTypeVictoryPoints MerchantBeerBonusType = "victorypoints"
	MerchantBeerBonusTypeMoney         MerchantBeerBonusType = "money'"
)

type MerchantBeerBonus struct {
	Type   MerchantBeerBonusType
	Amount int
}

type MerchantTile struct {
	IndustryTypes []IndustryType
}

type MerchantSpace struct {
	Tile MerchantTile
	Beer int
}

type Merchant struct {
	Links     int
	BeerBonus MerchantBeerBonus
	Spaces    []*MerchantSpace
}

type Location struct {
	Name               string
	CanalEraNeighbours []*Location
	RailEraNeighbours  []*Location

	IndustrySpaces []IndustrySpace
	Merchant       *Merchant
}
