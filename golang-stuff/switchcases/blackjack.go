package switchcases

import "strings"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	card = strings.ToLower(card)
	switch card {
	case "ace":
		return 11
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)
	switch {
	case playerTotal == 22:
		return "P"
	case playerTotal == 21:
		if dealerHand >= 10 {
			return "S"
		} else {
			return "W"
		}
	case playerTotal >= 17 && playerTotal <= 20:
		return "S"
	case playerTotal >= 12 && playerTotal <= 16:
		if dealerHand < 7 {
			return "S"
		} else {
			return "H"
		}
	default:
		return "H"
	}
}
