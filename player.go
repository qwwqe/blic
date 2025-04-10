package blic

type PlayerMat struct {
	CoalMineTiles  []IndustryTile
	IronWorksTiles []IndustryTile
	BreweryTiles   []IndustryTile

	ManufacturerTiles []IndustryTile
	CottonMillTiles   []IndustryTile
	PotteryTiles      []IndustryTile
}

func (mat PlayerMat) Clone() PlayerMat {
	mat.CoalMineTiles = CloneSlice(mat.CoalMineTiles)
	mat.IronWorksTiles = CloneSlice(mat.IronWorksTiles)
	mat.BreweryTiles = CloneSlice(mat.BreweryTiles)
	mat.ManufacturerTiles = CloneSlice(mat.ManufacturerTiles)
	mat.CottonMillTiles = CloneSlice(mat.CottonMillTiles)
	mat.PotteryTiles = CloneSlice(mat.PotteryTiles)

	return mat
}

type Player struct {
	Id string

	Mat              PlayerMat
	Money            int
	SpentMoney       int
	IncomeSpace      int
	VictoryPoints    int
	RemainingLinks   int
	RemainingActions int

	Cards         []Card
	Discards      []Card
	HiddenDiscard *Card
}

func (p Player) Clone() Player {
	p.Mat = p.Mat.Clone()
	p.Cards = CloneSlice(p.Cards)
	p.Discards = CloneSlice(p.Discards)

	if p.HiddenDiscard != nil {
		hiddenDiscard := *p.HiddenDiscard
		p.HiddenDiscard = &hiddenDiscard
	}

	return p
}
