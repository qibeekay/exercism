package techpalace

import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace" + ", " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var star string

    for i := 0; i < numStarsPerLine; i++ {
        star += "*"
    }

    return star + "\n" + welcomeMsg + "\n" + star
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    noStar := strings.ReplaceAll(oldMsg, "*", "")

    return strings.TrimSpace(noStar)
}
