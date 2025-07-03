package blic

func canConsumeIron(game *Game, player *Player, amount int) bool {
	return game.ConsumableIronOnIndustries() >= amount ||
		game.IronMarket.BuyPrice() <= player.Money
}
