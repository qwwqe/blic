package standard

import (
	"github.com/qwwqe/blic"
)

var playerMatSpec = blic.PlayerMatSpec{
	CoalMineTiles: []blic.IndustryTileSpec{
		coalMineTileSpecLevel1,
		coalMineTileSpecLevel2,
		coalMineTileSpecLevel2,
		coalMineTileSpecLevel3,
		coalMineTileSpecLevel3,
		coalMineTileSpecLevel4,
		coalMineTileSpecLevel4,
	},

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

var coalMineTileSpecLevel1 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeCoalMine,
		Level:    1,
		NumLinks: 2,

		CanalEraResources: 2,
		RailEraResources:  2,

		VictoryPoints:      1,
		IncomeBoost:        4,
		BeerRequiredToSell: 0,

		RequiredEra: address(blic.EraCanal),
		CanDevelop:  true,

		CostInPounds: 5,
		CostInCoal:   0,
		CostInIron:   0,
	},
}

var coalMineTileSpecLevel2 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeCoalMine,
		Level:    2,
		NumLinks: 1,

		CanalEraResources: 3,
		RailEraResources:  3,

		VictoryPoints:      2,
		IncomeBoost:        7,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 7,
		CostInCoal:   0,
		CostInIron:   0,
	},
}

var coalMineTileSpecLevel3 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeCoalMine,
		Level:    3,
		NumLinks: 1,

		CanalEraResources: 4,
		RailEraResources:  4,

		VictoryPoints:      3,
		IncomeBoost:        6,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 8,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var coalMineTileSpecLevel4 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeCoalMine,
		Level:    4,
		NumLinks: 1,

		CanalEraResources: 5,
		RailEraResources:  5,

		VictoryPoints:      4,
		IncomeBoost:        5,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 10,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

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
