package standard

import "github.com/qwwqe/blic"

/**
 * GameSpec corresponds to the standard vanilla game rules and components.
 */
var GameSpec = blic.GameSpec{
	Name:    "standard",
	Version: "0.0.1",

	NumWildLocationCards: 4,
	NumWildIndustryCards: 4,

	MinPlayerCount: 2,
	MaxPlayerCount: 4,

	InitialCoalInMarket: 13,
	InitialIronInMarket: 8,

	StartingMoney:       17,
	StartingIncomeSpace: 10,
	HandSize:            8,
	LinksPerPlayer:      14,

	CardSpecs:               cardSpecs,
	PlayerMatSpec:           playerMatSpec,
	MerchantTileSpecs:       merchantTileSpecs,
	LocationSpecs:           locationSpecs,
	CanalEraConnectionSpecs: canalEraConnectionSpecs,
	RailEraConnectionSpecs:  railEraConnectionSpecs,
}
