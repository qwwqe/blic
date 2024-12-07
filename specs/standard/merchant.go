package standard

import "github.com/qwwqe/blic"

var merchantTileSpecs = []blic.MerchantTileSpec{
	newMerchantTileSpec(2),
	newMerchantTileSpec(2),
	newMerchantTileSpec(2, blic.IndustryTypeManufacturer),
	newMerchantTileSpec(2, blic.IndustryTypeCottonMill),
	newMerchantTileSpec(2, blic.IndustryTypeManufacturer, blic.IndustryTypeCottonMill, blic.IndustryTypePottery),
	newMerchantTileSpec(3),
	newMerchantTileSpec(3, blic.IndustryTypePottery),
	newMerchantTileSpec(4, blic.IndustryTypeManufacturer),
	newMerchantTileSpec(4, blic.IndustryTypeCottonMill),
}

func newMerchantTileSpec(minPlayers int, industryTypes ...blic.IndustryType) blic.MerchantTileSpec {
	return blic.MerchantTileSpec{
		MerchantTile: blic.MerchantTile{
			IndustryTypes: industryTypes,
		},
		MinPlayers: minPlayers,
	}
}
