package blic

import (
	"errors"
	"math"
)

type Market struct {
	NumTiers         int
	NumResources     int
	ResourcesPerTier int
}

func (m *Market) BuyPrice() int {
	return m.NumTiers - int(math.Ceil(float64(m.NumResources)/float64(m.ResourcesPerTier)))
}

func (m *Market) SellPrice() int {
	return m.NumTiers - int(math.Ceil(float64(m.NumResources+1)/float64(m.ResourcesPerTier)))
}

func (m *Market) Buy() int {
	price := m.BuyPrice()

	if m.NumResources > 0 {
		m.NumResources--
	}

	return price
}

func (m *Market) Sell() int {
	price := m.SellPrice()

	if m.CanSell() {
		m.NumResources++
	}

	return price
}

func (m *Market) CanSell() bool {
	return m.NumResources < (m.NumTiers-1)*m.ResourcesPerTier
}

func NewMarket(numTiers, numResources, resourcesPerTier int) (*Market, error) {
	if numTiers <= 0 {
		return nil, errors.New("Number of tiers must be greater than 0")
	}

	if numResources < 0 || numResources > (numTiers-1)*resourcesPerTier {
		return nil, errors.New("Invalid starting resources")
	}

	if resourcesPerTier <= 0 {
		return nil, errors.New("Num of resources per tier must be greater than 0")
	}

	return &Market{
		NumTiers:         numTiers,
		NumResources:     numResources,
		ResourcesPerTier: resourcesPerTier,
	}, nil
}
