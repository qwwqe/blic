package blic

type IndustryType string

func (t IndustryType) Clone() IndustryType {
	return t
}

const (
	IndustryTypeCoalMine     IndustryType = "coalmine"
	IndustryTypeIronWorks    IndustryType = "ironworks"
	IndustryTypeBrewery      IndustryType = "brewery"
	IndustryTypeCottonMill   IndustryType = "cottonmill"
	IndustryTypeManufacturer IndustryType = "manufacturer"
	IndustryTypePottery      IndustryType = "pottery"
)

type IndustryTile struct {
	Type     IndustryType
	Level    int
	NumLinks int

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

func (tile IndustryTile) Clone() IndustryTile {
	if tile.RequiredEra != nil {
		era := *tile.RequiredEra
		tile.RequiredEra = &era
	}

	return tile
}
