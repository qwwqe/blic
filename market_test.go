package blic

import "testing"

func TestMarket(t *testing.T) {
	t.Run("BuyPrice", func(t *testing.T) {
		cases := []struct {
			NumTiers         int
			NumResources     int
			ResourcesPerTier int
			BuyPrice         int
		}{{
			NumTiers:         6,
			NumResources:     0,
			ResourcesPerTier: 2,
			BuyPrice:         6,
		}, {
			NumTiers:         6,
			NumResources:     10,
			ResourcesPerTier: 2,
			BuyPrice:         1,
		}, {
			NumTiers:         6,
			NumResources:     7,
			ResourcesPerTier: 2,
			BuyPrice:         2,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if err != nil {
				t.Fatal(err)
			}

			if market.BuyPrice() != testCase.BuyPrice {
				t.Errorf("expected %v but got %v", testCase.BuyPrice, market.BuyPrice())
			}
		}
	})

	t.Run("SellPrice", func(t *testing.T) {
		cases := []struct {
			NumTiers         int
			NumResources     int
			ResourcesPerTier int
			SellPrice        int
		}{{
			NumTiers:         6,
			NumResources:     0,
			ResourcesPerTier: 2,
			SellPrice:        5,
		}, {
			NumTiers:         6,
			NumResources:     10,
			ResourcesPerTier: 2,
			SellPrice:        0,
		}, {
			NumTiers:         6,
			NumResources:     6,
			ResourcesPerTier: 2,
			SellPrice:        2,
		}, {
			NumTiers:         6,
			NumResources:     7,
			ResourcesPerTier: 2,
			SellPrice:        2,
		}, {
			NumTiers:         6,
			NumResources:     8,
			ResourcesPerTier: 2,
			SellPrice:        1,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if err != nil {
				t.Fatal(err)
			}

			if market.SellPrice() != testCase.SellPrice {
				t.Errorf("expected %v but got %v", testCase.SellPrice, market.SellPrice())
			}
		}
	})

	t.Run("Buy", func(t *testing.T) {
		cases := []struct {
			NumTiers           int
			NumResources       int
			ResourcesPerTier   int
			BuyPrice           int
			RemainingResources int
		}{{
			NumTiers:           6,
			NumResources:       0,
			ResourcesPerTier:   2,
			BuyPrice:           6,
			RemainingResources: 0,
		}, {
			NumTiers:           6,
			NumResources:       10,
			ResourcesPerTier:   2,
			BuyPrice:           1,
			RemainingResources: 9,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if err != nil {
				t.Fatal(err)
			}

			price := market.Buy()
			if price != testCase.BuyPrice {
				t.Errorf("expected %v but got %v", testCase.BuyPrice, price)
			}

			if market.NumResources != testCase.RemainingResources {
				t.Errorf("expected %v but got %v", testCase.RemainingResources, market.NumResources)
			}
		}
	})

	t.Run("Sell", func(t *testing.T) {
		cases := []struct {
			NumTiers           int
			NumResources       int
			ResourcesPerTier   int
			SellPrice          int
			RemainingResources int
		}{{
			NumTiers:           6,
			NumResources:       0,
			ResourcesPerTier:   2,
			SellPrice:          5,
			RemainingResources: 1,
		}, {
			NumTiers:           6,
			NumResources:       10,
			ResourcesPerTier:   2,
			SellPrice:          0,
			RemainingResources: 10,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if err != nil {
				t.Fatal(err)
			}

			price := market.Sell()
			if price != testCase.SellPrice {
				t.Errorf("expected %v but got %v", testCase.SellPrice, price)
			}

			if market.NumResources != testCase.RemainingResources {
				t.Errorf("expected %v but got %v", testCase.RemainingResources, market.NumResources)
			}
		}
	})

	t.Run("CanSell", func(t *testing.T) {
		cases := []struct {
			NumTiers         int
			NumResources     int
			ResourcesPerTier int
			CanSell          bool
		}{{
			NumTiers:         6,
			NumResources:     0,
			ResourcesPerTier: 2,
			CanSell:          true,
		}, {
			NumTiers:         6,
			NumResources:     10,
			ResourcesPerTier: 2,
			CanSell:          false,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if err != nil {
				t.Fatal(err)
			}

			if market.CanSell() != testCase.CanSell {
				t.Errorf("expected %v but got %v", testCase.CanSell, market.CanSell())
			}
		}
	})

	t.Run("NewMarket", func(t *testing.T) {
		cases := []struct {
			NumTiers         int
			NumResources     int
			ResourcesPerTier int
			ShouldErr        bool
		}{{
			NumTiers:         0,
			NumResources:     0,
			ResourcesPerTier: 2,
			ShouldErr:        true,
		}, {
			NumTiers:         6,
			NumResources:     -1,
			ResourcesPerTier: 2,
			ShouldErr:        true,
		}, {
			NumTiers:         6,
			NumResources:     11,
			ResourcesPerTier: 2,
			ShouldErr:        true,
		}, {
			NumTiers:         6,
			NumResources:     0,
			ResourcesPerTier: 0,
			ShouldErr:        true,
		}}

		for _, testCase := range cases {
			market, err := NewMarket(
				testCase.NumTiers,
				testCase.NumResources,
				testCase.ResourcesPerTier,
			)

			if (err == nil) == testCase.ShouldErr {
				t.Fatal(err)
			}

			if (market == nil) != testCase.ShouldErr {
				t.Fatal(err)
			}
		}
	})
}
