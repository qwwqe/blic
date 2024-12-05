package specs

import "github.com/qwwqe/blic"

/**
 * StandardGameSpec corresponds to the vanilla game rules and components.
 */
var StandardGameSpec = blic.GameSpec{
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
}

var cardSpecs = []blic.CardSpec{}
