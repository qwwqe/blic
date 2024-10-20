package main

type PlayerMat struct {
	CoalMineTiles  []IndustryTile
	IronWorksTiles []IndustryTile
	BreweryTiles   []IndustryTile

	ManufacturerTiles []IndustryTile
	CottonMillTiles   []IndustryTile
	PotteryTiles      []IndustryTile
}

type Player struct {
	Mat           PlayerMat
	Money         int
	IncomeSpace   int
	VictoryPoints int
	Links         int
	Cards         []Card
}
