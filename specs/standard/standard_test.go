package standard

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	cases := []struct {
		playerCount int
		cardCount   int
	}{
		{2, 40},
		{3, 54},
		{4, 64},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("Card count in %d-player game", testCase.playerCount), func(t *testing.T) {
			count := 0
			for _, spec := range cardSpecs {
				cards := spec.Build(testCase.playerCount)
				count += len(cards)
			}

			if count != testCase.cardCount {
				t.Errorf("Expected %d cards but got %d", testCase.cardCount, count)
			}
		})

	}
}

func TestIndustryTiles(t *testing.T) {
	t.Run("Industry tile count", func(t *testing.T) {
		const standardTileCount = 45

		mat := playerMatSpec.Build()
		count := len(mat.BreweryTiles) + len(mat.CoalMineTiles) + len(mat.CottonMillTiles) +
			len(mat.IronWorksTiles) + len(mat.ManufacturerTiles) + len(mat.PotteryTiles)

		if count != standardTileCount {
			t.Errorf("Expected %d tiles but got %d", standardTileCount, count)
		}
	})
}
