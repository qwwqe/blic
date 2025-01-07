package blic

type CardType string

const (
	CardTypeLocation     CardType = "location"
	CardTypeIndustry     CardType = "industry"
	CardTypeWildLocation CardType = "wildlocation"
	CardTypeWildIndustry CardType = "wildindustry"
)

type Card struct {
	Id            string
	IndustryTypes []IndustryType
	LocationName  string
	Type          CardType
}

func (c Card) Clone() Card {
	industryTypes := make([]IndustryType, len(c.IndustryTypes))
	copy(industryTypes, c.IndustryTypes)
	c.IndustryTypes = industryTypes

	return c
}
