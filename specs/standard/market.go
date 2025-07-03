package standard

import "github.com/qwwqe/blic"

var coalMarketSpec = blic.MarketSpec{
	NumTiers:          8,
	StartingResources: 13,
	ResourcesPerTier:  2,
}

var ironMarketSpec = blic.MarketSpec{
	NumTiers:          6,
	StartingResources: 8,
	ResourcesPerTier:  2,
}
