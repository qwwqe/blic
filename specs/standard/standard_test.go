package standard

import (
	"fmt"
	"testing"

	"github.com/qwwqe/blic"
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

func TestMerchantTiles(t *testing.T) {
	cases := []struct {
		playerCount int
		tileCount   int
	}{
		{2, 5},
		{3, 7},
		{4, 9},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("Merchant tile count in %d-player game", testCase.playerCount), func(t *testing.T) {
			count := 0
			for _, spec := range merchantTileSpecs {
				tile := spec.Build(testCase.playerCount)
				if tile != nil {
					count++
				}
			}

			if count != testCase.tileCount {
				t.Errorf("Expected %d tiles but got %d", testCase.tileCount, count)
			}
		})
	}
}

func TestConnections(t *testing.T) {
	cases := []struct {
		era   blic.Era
		specs []blic.ConnectionSpec
		count int
	}{
		{blic.EraCanal, canalEraConnectionSpecs, 32},
		{blic.EraRail, railEraConnectionSpecs, 39},
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%v era connections", testCase.era), func(t *testing.T) {
			count := 0
			for _, spec := range testCase.specs {
				connection := spec.Build()
				if len(connection.LocationNames) > 0 {
					count++
				}
			}

			if count != testCase.count {
				t.Errorf("Expected %d connections but got %d", testCase.count, count)
			}
		})
	}
}

func TestIncomeTrack(t *testing.T) {
	t.Run("Income spaces", func(t *testing.T) {
		incomeTrack := incomeTrackSpec.Build()
		if len(incomeTrack) != 100 {
			t.Errorf("Expected 100 income spaces but got %d", len(incomeTrack))
		}
	})

	t.Run("Income levels", func(t *testing.T) {
		incomeTrack := incomeTrackSpec.Build()
		levels := map[int]bool{}

		for _, level := range incomeTrack {
			levels[level] = true
		}

		if len(levels) != 41 {
			t.Errorf("Expected 41 income levels but got %d", len(levels))
		}
	})
}

func TestGame(t *testing.T) {
	minPlayers, maxPlayers := 2, 4
	playerIds := []string{"a", "b", "c", "d"}
	for numPlayers := minPlayers; numPlayers <= maxPlayers; numPlayers++ {
		t.Run(fmt.Sprintf("Game creation for %d players", numPlayers), func(t *testing.T) {
			_, err := GameSpec.Build(playerIds[0:numPlayers])
			if err != nil {
				t.Errorf("Received error when building game from spec: %v", err)
			}
		})
	}
}
