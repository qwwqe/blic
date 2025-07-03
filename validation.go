package blic

func canConsumeIron(game *Game, player *Player) bool {
	return game.isIronOnBoard() || game.IronMarket.BuyPrice() <= player.Money
}
