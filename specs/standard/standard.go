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

	LoanAmount:     30,
	HandSize:       8,
	LinksPerPlayer: 14,

	StartingMoney:          17,
	StartingIncomeSpace:    10,
	LoanIncomeLevelPenalty: 3,

	StartingActions: 1,
	ActionsPerTurn:  2,

	CardSpecs:               cardSpecs,
	PlayerMatSpec:           playerMatSpec,
	IncomeTrackSpec:         incomeTrackSpec,
	MerchantTileSpecs:       merchantTileSpecs,
	LocationSpecs:           locationSpecs,
	CanalEraConnectionSpecs: canalEraConnectionSpecs,
	RailEraConnectionSpecs:  railEraConnectionSpecs,
	CoalMarketSpec:          coalMarketSpec,
	IronMarketSpec:          ironMarketSpec,
}
