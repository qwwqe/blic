package standard

import (
	"github.com/qwwqe/blic"
)

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
