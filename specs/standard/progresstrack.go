package standard

import "github.com/qwwqe/blic"

var incomeTrackSpec = blic.IncomeTrackSpec{
	StartingIncomeLevel: -10,
	GraduationSpecs: []blic.IncomeTrackGraduationSpec{
		{SpacesPerIncomeLevel: 1, NumIncomeLevels: 11},
		{SpacesPerIncomeLevel: 2, NumIncomeLevels: 10},
		{SpacesPerIncomeLevel: 3, NumIncomeLevels: 10},
		{SpacesPerIncomeLevel: 4, NumIncomeLevels: 9},
		{SpacesPerIncomeLevel: 3, NumIncomeLevels: 1},
	},
}
