package blic

import (
	"errors"
	"fmt"
)

func standardGameLocations() []Location {
	return []Location{
		// Merchants
		{
			Name:     "shrewsbury",
			Merchant: NewMerchant(2, 1, 2, MerchantBeerBonusTypeVictoryPoints, 4),
		},
		{
			Name:     "gloucester",
			Merchant: NewMerchant(2, 2, 2, MerchantBeerBonusTypeDevelop, 1),
		},
		{
			Name:     "oxford",
			Merchant: NewMerchant(2, 2, 2, MerchantBeerBonusTypeIncomeBoost, 2),
		},
		{
			Name:     "warrington",
			Merchant: NewMerchant(2, 2, 3, MerchantBeerBonusTypeMoney, 5),
		},
		{
			Name:     "nottingham",
			Merchant: NewMerchant(2, 2, 4, MerchantBeerBonusTypeVictoryPoints, 3),
		},

		// Farm breweries
		{
			Name: "farmbrewery1",
			IndustrySpaces: []IndustrySpace{{
				Types: []IndustryType{IndustryTypeBrewery},
			}},
		},
		{
			Name: "farmbrewery2",
			IndustrySpaces: []IndustrySpace{{
				Types: []IndustryType{IndustryTypeBrewery},
			}},
		},

		// Blue
		{
			Name: "stokeontrent",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypePottery, IndustryTypeIronWorks}},
			},
		},
		{
			Name: "leek",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeCoalMine}},
			},
		},
		{
			Name: "stone",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
			},
		},
		{
			Name: "uttoxeter",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeBrewery}},
			},
		},

		// Teal
		{
			Name: "belper",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypePottery}},
			},
		},
		{
			Name: "derby",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeIronWorks}},
			},
		},

		// Red
		{
			Name: "stafford",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypePottery}},
			},
		},
		{
			Name: "cannock",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeCoalMine}},
			},
		},
		{
			Name: "burtonupontrent",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeBrewery}},
			},
		},
		{
			Name: "tamworth",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeCoalMine}},
			},
		},
		{
			Name: "walsall",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeIronWorks, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeBrewery}},
			},
		},

		// Yellow
		{
			Name: "coalbrookdale",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeIronWorks, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypeIronWorks}},
				{Types: []IndustryType{IndustryTypeCoalMine}},
			},
		},
		{
			Name: "wolverhampton",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
			},
		},
		{
			Name: "dudley",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeIronWorks}},
			},
		},
		{
			Name: "kidderminster",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeCottonMill}},
			},
		},
		{
			Name: "worcester",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill}},
				{Types: []IndustryType{IndustryTypeCottonMill}},
			},
		},

		// Purple
		{
			Name: "birmingham",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeManufacturer}},
				{Types: []IndustryType{IndustryTypeIronWorks}},
			},
		},
		{
			Name: "redditch",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeIronWorks}},
			},
		},
		{
			Name: "coventry",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypePottery}},
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeCoalMine}},
				{Types: []IndustryType{IndustryTypeIronWorks, IndustryTypeManufacturer}},
			},
		},
		{
			Name: "nuneaton",
			IndustrySpaces: []IndustrySpace{
				{Types: []IndustryType{IndustryTypeManufacturer, IndustryTypeBrewery}},
				{Types: []IndustryType{IndustryTypeCottonMill, IndustryTypeCoalMine}},
			},
		},
	}
}
}
