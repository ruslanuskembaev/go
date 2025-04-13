package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "king", "queen":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardVal1 := ParseCard(card1)
	cardVal2 := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)

	if cardVal1 == 11 && cardVal2 == 11 {
		return "P"
	}
	if cardVal1+cardVal2 == 21 {
		if dealerCardVal < 10 {
			return "W"
		} else {
			return "S"
		}
	}
	if cardVal1+cardVal2 <= 11 {
		return "H"
	}
	if (cardVal1+cardVal2 >= 12) && (cardVal1+cardVal2 <= 16) {
		if dealerCardVal >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
	if (cardVal1+cardVal2 >= 17) && (cardVal1+cardVal2 <= 20) {
		return "S"
	}
	return "0"
}
