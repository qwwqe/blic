package blic

type IndustrySpace struct {
	Types        []IndustryType
	Tile         IndustryTile
	TilePlayerId string // TODO: 考慮把 PlayerdId 掛在 IndustryTile 身上
	Resources    int
}

func (s IndustrySpace) Clone() IndustrySpace {
	s.Types = CloneSlice(s.Types)
	s.Tile = s.Tile.Clone()

	return s
}

type MerchantBeerBonusType string

const (
	MerchantBeerBonusTypeDevelop       MerchantBeerBonusType = "develop"
	MerchantBeerBonusTypeIncomeBoost   MerchantBeerBonusType = "incomeboost"
	MerchantBeerBonusTypeVictoryPoints MerchantBeerBonusType = "victorypoints"
	MerchantBeerBonusTypeMoney         MerchantBeerBonusType = "money"
)

type MerchantBeerBonus struct {
	Type   MerchantBeerBonusType
	Amount int
}

func (s MerchantBeerBonus) Clone() MerchantBeerBonus {
	return s
}

type MerchantTile struct {
	IndustryTypes []IndustryType
}

func (t MerchantTile) Clone() MerchantTile {
	t.IndustryTypes = CloneSlice(t.IndustryTypes)
	return t
}

type MerchantSpace struct {
	Tile MerchantTile
	Beer int
}

func (s MerchantSpace) Clone() MerchantSpace {
	s.Tile = s.Tile.Clone()
	return s
}

type Merchant struct {
	BeerBonus MerchantBeerBonus
	NumLinks  int
	Spaces    []MerchantSpace
}

func (m Merchant) Clone() Merchant {
	m.Spaces = CloneSlice(m.Spaces)
	return m
}

type Location struct {
	Name string

	IndustrySpaces []IndustrySpace
	Merchant       *Merchant
}

func (l Location) Clone() Location {
	l.IndustrySpaces = CloneSlice(l.IndustrySpaces)
	if l.Merchant != nil {
		m := l.Merchant.Clone()
		l.Merchant = &m
	}
	return l
}

type Connection struct {
	LocationNames []string
	LinkPlayerId  string // 考慮正式定義包含 PlayerId 和 Era 的 LinkTile
}

func (c Connection) Clone() Connection {
	locationNames := make([]string, len(c.LocationNames))
	copy(locationNames, c.LocationNames)
	c.LocationNames = locationNames

	return c
}
