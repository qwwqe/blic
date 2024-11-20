package blic

import (
	"errors"
	"fmt"
)

// TODO: stub
var StandardGameSpec = GameSpec{}

// TODO: 改成 Spec 實作
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

var standardGameDualLinks = [][]string{
	{"warrington", "stokeontrent"},
	{"stokeontrent", "leek"},
	{"stokeontrent", "stone"},
	{"nottingham", "derby"},
	{"derby", "belper"},
	{"stafford", "stone"},
	{"stafford", "cannock"},
	{"cannock", "walsall"},
	{"cannock", "farmbrewery1"},
	{"burtonupontrent", "derby"},
	{"burtonupontrent", "tamworth"},
	{"burtonupontrent", "stone"},
	{"shrewsbury", "coalbrookdale"},
	{"coalbrookdale", "wolverhampton"},
	{"coalbrookdale", "kidderminster"},
	{"wolverhampton", "dudley"},
	{"dudley", "kidderminster"},
	{"kidderminster", "worcester"},
	{"kidderminster", "farmbrewery2"}, // TODO: support "three-way edges"
	{"worcester", "farmbrewery2"},     // TODO: support "three-way edges"
	{"gloucester", "worcester"},
	{"gloucester", "redditch"},
	{"redditch", "oxford"},
	{"oxford", "birmingham"},
	{"birmingham", "worcester"},
	{"birmingham", "dudley"},
	{"birmingham", "walsall"},
	{"birmingham", "tamworth"},
	{"birmingham", "coventry"},
	{"nuneaton", "tamworth"},
}

var standardGameCanalLinks = [][]string{
	{"walsall", "burtonupontrent"},
}

var standardGameRailLinks = [][]string{
	{"leek", "belper"},
	{"stone", "uttoxeter"},
	{"uttoxeter", "derby"},
	{"cannock", "burtonupontrent"},
	{"walsall", "tamworth"},
	{"nuneaton", "birmingham"},
	{"nuneaton", "coventry"},
	{"birmingham", "redditch"},
}

var (
	ErrDuplicateLocationName = errors.New("Duplicate location name")
	ErrInvalidLink           = errors.New("Invalid link")
	ErrDuplicateLink         = errors.New("Duplicate link")
	ErrOrphanLocation        = errors.New("Location has no neighbours")
)

// TODO: 改成 Spec 實作
func addLink(link []string,
	locationLookup map[string]*Location, neighbourLookup map[string]map[string]bool, isCanal bool) error {
	if len(link) != 2 {
		return fmt.Errorf("%w: %v", ErrInvalidLink, link)
	}

	sourceName, destName := link[0], link[1]

	source, sourceOk := locationLookup[sourceName]
	if !sourceOk {
		return fmt.Errorf("%w: %s", ErrNonExistentLocation, sourceName)
	}

	dest, destOk := locationLookup[destName]
	if !destOk {
		return fmt.Errorf("%w: %s", ErrNonExistentLocation, destName)
	}

	if neighbourLookup[sourceName][destName] || neighbourLookup[destName][sourceName] {
		return fmt.Errorf("%w: (%s, %s)", ErrDuplicateLink, sourceName, destName)
	}

	neighbourLookup[sourceName][destName] = true
	neighbourLookup[destName][sourceName] = true

	sourceNeighbours := &source.CanalEraNeighbours
	destNeighbours := &dest.CanalEraNeighbours
	if !isCanal {
		sourceNeighbours = &source.RailEraNeighbours
		destNeighbours = &dest.RailEraNeighbours
	}

	(*sourceNeighbours) = append(*sourceNeighbours, destName)
	(*destNeighbours) = append(*destNeighbours, sourceName)

	return nil
}

func populateLocationLinks(
	locations []Location, dualLinks [][]string,
	canalLinks [][]string, railLinks [][]string,
) error {
	locationLookup := map[string]*Location{}
	locationCanalNeighbourLookup := map[string]map[string]bool{}
	locationRailNeighbourLookup := map[string]map[string]bool{}
	for i, location := range locations {
		if _, ok := locationLookup[location.Name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateLocationName, location.Name)
		}

		locationLookup[location.Name] = &locations[i]
		locationCanalNeighbourLookup[location.Name] = map[string]bool{}
		locationRailNeighbourLookup[location.Name] = map[string]bool{}
	}

	for _, link := range dualLinks {
		if err := addLink(link, locationLookup, locationCanalNeighbourLookup, true); err != nil {
			return err
		}

		if err := addLink(link, locationLookup, locationRailNeighbourLookup, false); err != nil {
			return err
		}
	}

	for _, link := range canalLinks {
		if err := addLink(link, locationLookup, locationCanalNeighbourLookup, true); err != nil {
			return err
		}
	}

	for _, link := range railLinks {
		if err := addLink(link, locationLookup, locationRailNeighbourLookup, false); err != nil {
			return err
		}
	}

	return nil
}

func buildStandardLocations() ([]Location, error) {
	locations := standardGameLocations()
	if err := populateLocationLinks(locations, standardGameDualLinks, standardGameCanalLinks, standardGameRailLinks); err != nil {
		return nil, err
	}

	for _, location := range locations {
		if len(location.CanalEraNeighbours) == 0 &&
			len(location.RailEraNeighbours) == 0 {
			return nil, ErrOrphanLocation
		}
	}

	return locations, nil
}

func mustBuildStandardLocations() []Location {
	locations, err := buildStandardLocations()
	if err != nil {
		panic("setup: " + err.Error())
	}

	return locations
}
