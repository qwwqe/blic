package standard

import "github.com/qwwqe/blic"

var locationSpecs = []blic.LocationSpec{
	// Merchants
	{
		Name: "shrewsbury",
		Merchant: &blic.MerchantSpec{
			NumLinks: 2,
			BeerBonus: blic.MerchantBeerBonusSpec{
				MerchantBeerBonus: blic.MerchantBeerBonus{
					Type:   blic.MerchantBeerBonusTypeVictoryPoints,
					Amount: 4,
				},
			},
			NumSpaces:  1,
			MinPlayers: 2,
		},
	},
	{
		Name: "gloucester",
		Merchant: &blic.MerchantSpec{
			NumLinks: 2,
			BeerBonus: blic.MerchantBeerBonusSpec{
				MerchantBeerBonus: blic.MerchantBeerBonus{
					Type:   blic.MerchantBeerBonusTypeDevelop,
					Amount: 1,
				},
			},
			NumSpaces:  2,
			MinPlayers: 2,
		},
	},
	{
		Name: "oxford",
		Merchant: &blic.MerchantSpec{
			NumLinks: 2,
			BeerBonus: blic.MerchantBeerBonusSpec{
				MerchantBeerBonus: blic.MerchantBeerBonus{
					Type:   blic.MerchantBeerBonusTypeIncomeBoost,
					Amount: 2,
				},
			},
			NumSpaces:  2,
			MinPlayers: 2,
		},
	},
	{
		Name: "warrington",
		Merchant: &blic.MerchantSpec{
			NumLinks: 2,
			BeerBonus: blic.MerchantBeerBonusSpec{
				MerchantBeerBonus: blic.MerchantBeerBonus{
					Type:   blic.MerchantBeerBonusTypeMoney,
					Amount: 5,
				},
			},
			NumSpaces:  2,
			MinPlayers: 3,
		},
	},
	{
		Name: "nottingham",
		Merchant: &blic.MerchantSpec{
			NumLinks: 2,
			BeerBonus: blic.MerchantBeerBonusSpec{
				MerchantBeerBonus: blic.MerchantBeerBonus{
					Type:   blic.MerchantBeerBonusTypeVictoryPoints,
					Amount: 3,
				},
			},
			NumSpaces:  2,
			MinPlayers: 4,
		},
	},

	// Farm breweries
	{
		Name: "farmbrewery1",
		IndustrySpaces: []blic.IndustrySpaceSpec{{
			Types: []blic.IndustryType{blic.IndustryTypeBrewery},
		}},
	},
	{
		Name: "farmbrewery2",
		IndustrySpaces: []blic.IndustrySpaceSpec{{
			Types: []blic.IndustryType{blic.IndustryTypeBrewery},
		}},
	},

	// Blue
	{
		Name: "stoke-on-trent",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypePottery, blic.IndustryTypeIronWorks}},
		},
	},
	{
		Name: "leek",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "stone",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "uttoxeter",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeBrewery}},
		},
	},

	// Teal
	{
		Name: "belper",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypePottery}},
		},
	},
	{
		Name: "derby",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks}},
		},
	},

	// Red
	{
		Name: "stafford",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypePottery}},
		},
	},
	{
		Name: "cannock",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "burton-upon-trent",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeBrewery}},
		},
	},
	{
		Name: "tamworth",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "walsall",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeBrewery}},
		},
	},

	// Yellow
	{
		Name: "coalbrookdale",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks}},
			{Types: []blic.IndustryType{blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "wolverhampton",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
		},
	},
	{
		Name: "dudley",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks}},
		},
	},
	{
		Name: "kidderminster",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill}},
		},
	},
	{
		Name: "worcester",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill}},
		},
	},

	// Purple
	{
		Name: "birmingham",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks}},
		},
	},
	{
		Name: "redditch",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks}},
		},
	},
	{
		Name: "coventry",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypePottery}},
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeCoalMine}},
			{Types: []blic.IndustryType{blic.IndustryTypeIronWorks, blic.IndustryTypeManufacturer}},
		},
	},
	{
		Name: "nuneaton",
		IndustrySpaces: []blic.IndustrySpaceSpec{
			{Types: []blic.IndustryType{blic.IndustryTypeManufacturer, blic.IndustryTypeBrewery}},
			{Types: []blic.IndustryType{blic.IndustryTypeCottonMill, blic.IndustryTypeCoalMine}},
		},
	},
}
