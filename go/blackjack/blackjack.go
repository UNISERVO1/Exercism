package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValue := 0
	switch {
	case card == "ace":
		cardValue = 11
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		cardValue = 10
	case card == "nine":
		cardValue = 9
	case card == "eight":
		cardValue = 8
	case card == "seven":
		cardValue = 7
	case card == "six":
		cardValue = 6
	case card == "five":
		cardValue = 5
	case card == "four":
		cardValue = 4
	case card == "three":
		cardValue = 3
	case card == "two":
		cardValue = 2
	}
	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	hand := ParseCard(card1) + ParseCard(card2)
	decision := "H"

	switch {
	case hand == 21:
		if ParseCard(dealerCard) < 10 {
			decision = "W"
		} else {
			decision = "S"
		}
	case card1 == "ace" && card2 == "ace":
		decision = "P"
	case hand >= 17:
		decision = "S"
	case hand >= 12:
		if ParseCard(dealerCard) >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	}

	return decision
}
