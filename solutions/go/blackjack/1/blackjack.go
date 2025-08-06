package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    cardValue  := 0
	switch card {
    case "ace":
    	cardValue = 11
    case "two":
    	cardValue = 2
    case "three":
    	cardValue = 3
    case "four":
    	cardValue = 4
    case "five":
    	cardValue = 5
    case "six":
    	cardValue = 6
    case "seven":
    	cardValue = 7
    case "eight":
    	cardValue = 8
    case "nine":
    	cardValue = 9
    case "ten":
    	fallthrough
    case "jack":
    	fallthrough
    case "queen":
    	fallthrough
    case "king":
    	cardValue = 10
    }
	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    
    switch {
    case handValue == 22:
    	return "P"
    case handValue == 21 && dealerValue < 10:
    	return "W"
    case handValue == 21:
    	return "S"
    case handValue >= 17 && handValue <= 20:
    	return "S"
    case handValue >= 12 && handValue <= 16 && dealerValue < 7:
    	return "S"
    default:
    	return "H"
    }
}
