package specs

import "github.com/qwwqe/blic"

/**
 * Generate a Location CardSpec.
 *
 * The value at index i of amountByPlayerCount corresponds to the number of
 * cards present in a game with i + 2 players.
 */
func newLocationCardSpec(locationName string, amountByPlayerCount ...int) blic.CardSpec {
	spec := blic.CardSpec{
		Type:                blic.CardTypeLocation,
		LocationName:        locationName,
		AmountByPlayerCount: map[int]int{},
	}

	for n, amount := range amountByPlayerCount {
		spec.AmountByPlayerCount[n+2] = amount
	}

	return spec
}

/**
 * Generate an Industry CardSpec.
 *
 * Consult newLocationCardSpec for details about amountByPlayerCount.
 */
func newIndustryCardSpec(industryType blic.IndustryType, amountByPlayerCount ...int) blic.CardSpec {
	spec := blic.CardSpec{
		Type:                blic.CardTypeIndustry,
		IndustryTypes:       []blic.IndustryType{industryType},
		AmountByPlayerCount: map[int]int{},
	}

	for n, amount := range amountByPlayerCount {
		spec.AmountByPlayerCount[n+2] = amount
	}

	return spec
}
