package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	viewCard := card
    value := 0

    switch viewCard {
    case "ace":
        value = 11
    case "two":
        value = 2
    case "three":
        value = 3
    case "four":
        value = 4
    case "five":
        value = 5
    case "six":
        value = 6
    case "seven":
        value = 7
    case "eight":
        value = 8
    case "nine":
        value = 9
    case "ten", "jack", "queen", "king":
        value = 10
    default:
        value = 0  
    }

    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    // S = stand H = hit P = split W = win
    result := ""
    sum := ParseCard(card1) + ParseCard(card2)
    dealer := ParseCard(dealerCard)
	switch {
    case card1 == "ace" && card2 == "ace":
        result = "P"
    case sum == 21 && dealer < 10:
        result = "W"
    case sum == 21:
        result = "S"
    case sum >= 17 && sum <= 20:
        result = "S"
    case sum >= 12 && sum <= 16 && dealer >= 7:
        result = "H"
    case sum >= 12 && sum <= 16 && dealer < 7:
        result = "S"
    case sum <= 11:
        result = "H"
    default:
        result = "W"    
	}
    return result
}
