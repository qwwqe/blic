package blic

import "iter"

type PlayerMat struct {
	CoalMineTiles  []IndustryTile
	IronWorksTiles []IndustryTile
	BreweryTiles   []IndustryTile

	ManufacturerTiles []IndustryTile
	CottonMillTiles   []IndustryTile
	PotteryTiles      []IndustryTile
}

func (mat *PlayerMat) Industries() iter.Seq2[IndustryType, []IndustryTile] {
	return func(yield func(IndustryType, []IndustryTile) bool) {
		if !yield(IndustryTypeCoalMine, mat.CoalMineTiles) {
			return
		}
		if !yield(IndustryTypeIronWorks, mat.IronWorksTiles) {
			return
		}
		if !yield(IndustryTypeBrewery, mat.BreweryTiles) {
			return
		}
		if !yield(IndustryTypeManufacturer, mat.ManufacturerTiles) {
			return
		}
		if !yield(IndustryTypeCottonMill, mat.CottonMillTiles) {
			return
		}
		if !yield(IndustryTypePottery, mat.PotteryTiles) {
			return
		}
	}
}

func (mat *PlayerMat) HasDevelopableIndustry() bool {
	for _, industryTiles := range mat.Industries() {
		if len(industryTiles) > 0 && industryTiles[0].CanDevelop {
			return true
		}
	}

	return false
}

func (mat *PlayerMat) CanDevelopIndustries(industryTypes []IndustryType) bool {
	industryTypeOffset := map[IndustryType]int{}

	for _, industryType := range industryTypes {
		offset := industryTypeOffset[industryType]
		industryTiles := mat.IndustryTilesByType(industryType)

		if offset >= len(industryTiles) || !industryTiles[offset].CanDevelop {
			return false
		}

		industryTypeOffset[industryType]++
	}

	return true
}

func (mat *PlayerMat) Develop(industryType IndustryType) bool {
	industryTiles := mat.industryTilesByType(industryType)

	if len(*industryTiles) == 0 || !(*industryTiles)[0].CanDevelop {
		return false
	}

	*industryTiles = (*industryTiles)[1:]

	return true
}

func (mat *PlayerMat) industryTilesByType(industryType IndustryType) *[]IndustryTile {
	switch industryType {
	case IndustryTypeCoalMine:
		return &mat.CoalMineTiles
	case IndustryTypeIronWorks:
		return &mat.IronWorksTiles
	case IndustryTypeBrewery:
		return &mat.BreweryTiles
	case IndustryTypeManufacturer:
		return &mat.ManufacturerTiles
	case IndustryTypeCottonMill:
		return &mat.CottonMillTiles
	case IndustryTypePottery:
		return &mat.PotteryTiles
	default:
		return nil
	}
}

func (mat *PlayerMat) IndustryTilesByType(industryType IndustryType) []IndustryTile {
	return *mat.industryTilesByType(industryType)
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
