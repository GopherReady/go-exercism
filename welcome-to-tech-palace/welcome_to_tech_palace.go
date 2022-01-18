package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	parserStr := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", parserStr, welcomeMsg, parserStr)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
	msg = strings.TrimSpace(msg)
	return msg
}
