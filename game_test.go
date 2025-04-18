package blic

import (
	"reflect"
	"testing"
)

func TestHandleGameCreatedEvent(t *testing.T) {
	t.Run("records event in store", func(t *testing.T) {
		e := GameCreatedEvent{GameId: "test"}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if !reflect.DeepEqual(game.Events[0], e) {
			t.Errorf("expected %v but got %v", e, game.Events[0])
		}
	})

	t.Run("clears existing events in store", func(t *testing.T) {
		e := GameCreatedEvent{GameId: "test"}
		game := Game{
			Events: []Event{GameCreatedEvent{}, GameCreatedEvent{}},
		}
		game.HandleGameCreatedEvent(e)

		if !reflect.DeepEqual(game.Events, []Event{e}) {
			t.Errorf("expected %v but got %v", []Event{e}, game.Events)
		}
	})

	t.Run("allocates correct number of wild location cards", func(t *testing.T) {
		e := GameCreatedEvent{NumWildLocationCards: 24}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if len(game.WildLocationCards) != e.NumWildLocationCards {
			t.Errorf("expected %v but got %v", e.NumWildLocationCards, len(game.WildLocationCards))
		}
	})

	t.Run("allocates correct number of wild industry cards", func(t *testing.T) {
		e := GameCreatedEvent{NumWildIndustryCards: 99}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if len(game.WildIndustryCards) != e.NumWildIndustryCards {
			t.Errorf("expected %v but got %v", e.NumWildIndustryCards, len(game.WildIndustryCards))
		}
	})

	t.Run("initializes starting player index", func(t *testing.T) {
		e := GameCreatedEvent{}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if game.PlayerIndex != 0 {
			t.Errorf("expected %v but got %v", 0, game.PlayerIndex)
		}
	})

	t.Run("initializes starting round", func(t *testing.T) {
		e := GameCreatedEvent{}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if game.Round != 0 {
			t.Errorf("expected %v but got %v", 0, game.Round)
		}
	})

	t.Run("initializes starting phase", func(t *testing.T) {
		e := GameCreatedEvent{}
		game := Game{}
		game.HandleGameCreatedEvent(e)

		if game.Phase != GamePhaseAction {
			t.Errorf("expected %v but got %v", 0, GamePhaseAction)
		}
	})
}

func TestCalculateDeductedIncomeSpace(t *testing.T) {
	cases := []struct {
		incomeTrack         []int
		currentIncomeSpace  int
		deductedLevels      int
		expectedIncomeSpace int
	}{
		{[]int{-10, -9, -8, -7}, 3, 3, 0},
		{[]int{-10, -9, -8}, 2, 3, -1},
		{[]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4}, 7, 3, 1},
		{[]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4}, 8, 3, 1},
		{[]int{1, 1, 1, 2, 3, 3, 4, 5, 5, 5, 6}, 9, 3, 3},
	}

	for _, testCase := range cases {
		newIncomeSpace := calculateDeductedIncomeSpace(testCase.incomeTrack, testCase.currentIncomeSpace, testCase.deductedLevels)
		if newIncomeSpace != testCase.expectedIncomeSpace {
			t.Errorf("expected %v but got %v", testCase.expectedIncomeSpace, newIncomeSpace)
		}
	}
}
