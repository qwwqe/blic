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

	CardSpecs:     cardSpecs,
	PlayerMatSpec: playerMatSpec,
}

var cardSpecs = []blic.CardSpec{
	newLocationCardSpec("birmingham", 3, 3, 3),
	newLocationCardSpec("coventry", 3, 3, 3),
	newLocationCardSpec("nuneaton", 1, 1, 1),
	newLocationCardSpec("redditch", 1, 1, 1),
	newLocationCardSpec("coalbrookdale", 3, 3, 3),
	newLocationCardSpec("dudley", 2, 2, 2),
	newLocationCardSpec("kidderminster", 2, 2, 2),
	newLocationCardSpec("wolverhampton", 2, 2, 2),
	newLocationCardSpec("worcester", 2, 2, 2),
	newLocationCardSpec("burtonupontrent", 2, 2, 2),
	newLocationCardSpec("cannock", 2, 2, 2),
	newLocationCardSpec("stafford", 2, 2, 2),
	newLocationCardSpec("tamworth", 1, 1, 1),
	newLocationCardSpec("walsall", 1, 1, 1),
	newLocationCardSpec("leek", 0, 2, 2),
	newLocationCardSpec("stokeontrent", 0, 3, 3),
	newLocationCardSpec("stone", 0, 2, 2),
	newLocationCardSpec("uttoxeter", 0, 1, 2),
	newLocationCardSpec("belper", 0, 0, 2),
	newLocationCardSpec("derby", 0, 0, 3),
	newIndustryCardSpec(blic.IndustryTypeBrewery, 5, 5, 5),
	newIndustryCardSpec(blic.IndustryTypeCoalMine, 2, 2, 3),
	newIndustryCardSpec(blic.IndustryTypeIronWorks, 4, 4, 4),
	newIndustryCardSpec(blic.IndustryTypePottery, 2, 2, 3),
	{
		Type: blic.CardTypeIndustry,
		IndustryTypes: []blic.IndustryType{
			blic.IndustryTypeManufacturer,
			blic.IndustryTypeCottonMill,
		},
		AmountByPlayerCount: map[int]int{
			3: 6,
			4: 8,
		},
	},
}

var playerMatSpec = blic.PlayerMatSpec{
	BreweryTiles: []blic.IndustryTileSpec{
		breweryTileSpecLevel1,
		breweryTileSpecLevel1,
		breweryTileSpecLevel2,
		breweryTileSpecLevel2,
		breweryTileSpecLevel3,
		breweryTileSpecLevel3,
		breweryTileSpecLevel4,
	},
}

// TODO: Consider using either code generation or populating the specs from YAML?

var breweryTileSpecLevel1 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeBrewery,
		Level:    1,
		NumLinks: 2,

		CanalEraResources: 1,
		RailEraResources:  1,

		VictoryPoints:      4,
		IncomeBoost:        4,
		BeerRequiredToSell: 0,

		RequiredEra: address(blic.EraCanal),
		CanDevelop:  true,

		CostInPounds: 5,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var breweryTileSpecLevel2 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeBrewery,
		Level:    2,
		NumLinks: 2,

		CanalEraResources: 1,
		RailEraResources:  2,

		VictoryPoints:      5,
		IncomeBoost:        5,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 7,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var breweryTileSpecLevel3 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeBrewery,
		Level:    3,
		NumLinks: 2,

		CanalEraResources: 1,
		RailEraResources:  2,

		VictoryPoints:      7,
		IncomeBoost:        5,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 9,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var breweryTileSpecLevel4 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeBrewery,
		Level:    4,
		NumLinks: 2,

		CanalEraResources: 1,
		RailEraResources:  2,

		VictoryPoints:      10,
		IncomeBoost:        5,
		BeerRequiredToSell: 0,

		RequiredEra: address(blic.EraRail),
		CanDevelop:  true,

		CostInPounds: 9,
		CostInCoal:   0,
		CostInIron:   1,
	},
}
