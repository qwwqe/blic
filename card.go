package main

type Card struct {
	IndustryTypes []IndustryType
	LocationName  string

	IsWildLocation bool
	IsWildIndustry bool
}

func (c Card) Clone() Card {
	industryTypes := make([]IndustryType, len(c.IndustryTypes))
	copy(industryTypes, c.IndustryTypes)
	c.IndustryTypes = industryTypes

	return c
}
