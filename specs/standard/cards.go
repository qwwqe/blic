package standard

import (
	"github.com/qwwqe/blic"
)

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
