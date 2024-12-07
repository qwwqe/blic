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

	IronWorksTiles: []blic.IndustryTileSpec{
		ironWorksTileSpecLevel1,
		ironWorksTileSpecLevel2,
		ironWorksTileSpecLevel3,
		ironWorksTileSpecLevel4,
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

	ManufacturerTiles: []blic.IndustryTileSpec{
		manufacturerTileSpecLevel1,
		manufacturerTileSpecLevel2,
		manufacturerTileSpecLevel2,
		manufacturerTileSpecLevel3,
		manufacturerTileSpecLevel4,
		manufacturerTileSpecLevel5,
		manufacturerTileSpecLevel5,
		manufacturerTileSpecLevel6,
		manufacturerTileSpecLevel7,
		manufacturerTileSpecLevel8,
		manufacturerTileSpecLevel8,
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

var ironWorksTileSpecLevel1 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeIronWorks,
		Level:    1,
		NumLinks: 1,

		CanalEraResources: 4,
		RailEraResources:  4,

		VictoryPoints:      3,
		IncomeBoost:        3,
		BeerRequiredToSell: 0,

		RequiredEra: address(blic.EraCanal),
		CanDevelop:  true,

		CostInPounds: 5,
		CostInCoal:   1,
		CostInIron:   0,
	},
}

var ironWorksTileSpecLevel2 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeIronWorks,
		Level:    2,
		NumLinks: 1,

		CanalEraResources: 4,
		RailEraResources:  4,

		VictoryPoints:      5,
		IncomeBoost:        3,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 7,
		CostInCoal:   1,
		CostInIron:   0,
	},
}

var ironWorksTileSpecLevel3 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeIronWorks,
		Level:    3,
		NumLinks: 1,

		CanalEraResources: 5,
		RailEraResources:  5,

		VictoryPoints:      7,
		IncomeBoost:        2,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 9,
		CostInCoal:   1,
		CostInIron:   0,
	},
}

var ironWorksTileSpecLevel4 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeIronWorks,
		Level:    4,
		NumLinks: 1,

		CanalEraResources: 6,
		RailEraResources:  6,

		VictoryPoints:      9,
		IncomeBoost:        1,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 12,
		CostInCoal:   1,
		CostInIron:   0,
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

var manufacturerTileSpecLevel1 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    1,
		NumLinks: 2,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      3,
		IncomeBoost:        5,
		BeerRequiredToSell: 1,

		RequiredEra: address(blic.EraCanal),
		CanDevelop:  true,

		CostInPounds: 8,
		CostInCoal:   1,
		CostInIron:   0,
	},
}

var manufacturerTileSpecLevel2 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    2,
		NumLinks: 1,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      5,
		IncomeBoost:        1,
		BeerRequiredToSell: 1,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 10,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var manufacturerTileSpecLevel3 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    3,
		NumLinks: 0,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      4,
		IncomeBoost:        4,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 12,
		CostInCoal:   2,
		CostInIron:   0,
	},
}

var manufacturerTileSpecLevel4 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    4,
		NumLinks: 1,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      3,
		IncomeBoost:        6,
		BeerRequiredToSell: 1,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 8,
		CostInCoal:   0,
		CostInIron:   1,
	},
}

var manufacturerTileSpecLevel5 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    5,
		NumLinks: 2,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      8,
		IncomeBoost:        2,
		BeerRequiredToSell: 2,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 16,
		CostInCoal:   1,
		CostInIron:   0,
	},
}

var manufacturerTileSpecLevel6 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    6,
		NumLinks: 1,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      7,
		IncomeBoost:        6,
		BeerRequiredToSell: 1,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 20,
		CostInCoal:   0,
		CostInIron:   0,
	},
}

var manufacturerTileSpecLevel7 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    7,
		NumLinks: 0,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      9,
		IncomeBoost:        4,
		BeerRequiredToSell: 0,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 16,
		CostInCoal:   1,
		CostInIron:   1,
	},
}

var manufacturerTileSpecLevel8 = blic.IndustryTileSpec{
	IndustryTile: blic.IndustryTile{
		Type:     blic.IndustryTypeManufacturer,
		Level:    8,
		NumLinks: 1,

		CanalEraResources: 0,
		RailEraResources:  0,

		VictoryPoints:      11,
		IncomeBoost:        1,
		BeerRequiredToSell: 1,

		RequiredEra: nil,
		CanDevelop:  true,

		CostInPounds: 20,
		CostInCoal:   0,
		CostInIron:   2,
	},
}
