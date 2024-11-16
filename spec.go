package blic

type GameSpec struct {
	CardSpecs     []CardSpec
	LocationSpecs []LocationSpec
}

type CardSpec struct {
	IndustryTypes []IndustryType
	LocationName  string

	IsWildLocation bool
	IsWildIndustry bool

	AmountByPlayerCount map[int]int
}

func (s *CardSpec) Build(playerCount int) []Card {
	cards := []Card{}

	for i := 0; i < s.AmountByPlayerCount[playerCount]; i++ {
		cards = append(cards, Card{
			// NOTE: Depending on how many games are expected to be processed
			// simultaneously, it may be worthwhile not cloning resources
			// that are in principle expected to be immutable.
			IndustryTypes:  CloneSlice(s.IndustryTypes),
			LocationName:   s.LocationName,
			IsWildLocation: s.IsWildLocation,
			IsWildIndustry: s.IsWildIndustry,
		})
	}

	return cards
}
