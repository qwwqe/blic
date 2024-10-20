package main

type IndustryType string

const (
	IndustryTypeCoalMine     IndustryType = "coalmine"
	IndustryTypeIronWorks    IndustryType = "ironworks"
	IndustryTypeBrewery      IndustryType = "brewery"
	IndustryTypeCottonMill   IndustryType = "cottonmill"
	IndustryTypeManufacturer IndustryType = "manufacturer"
	IndustryTypePottery      IndustryType = "pottery"
)

type IndustryTile struct {
	Type  IndustryType
	Level int
	Links int

	CanalEraResources int
	RailEraResources  int

	VictoryPoints      int
	IncomeBoost        int
	BeerRequiredToSell int

	RequiredEra *Era
	CanDevelop  bool

	CostInPounds int
	CostInCoal   int
	CostInIron   int
}
