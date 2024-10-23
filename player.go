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
	Id int

	Mat           PlayerMat
	Money         int
	SpentMoney    int
	IncomeSpace   int
	VictoryPoints int
	Links         int

	Cards         []Card
	Discards      []Card
	HiddenDiscard *Card
}
