package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	if customer == "" {
		return "Welcome to the Tech Palace!"
	}
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	if numStarsPerLine <= 0 {
		return welcomeMsg
	}
	border := ""
	for i := 0; i < numStarsPerLine; i++ {
		border += "*"
	}
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	if oldMsg == "" {
		return ""
	}
	cleanedMsg := ""
	for _, char := range oldMsg {
		if char != '*' && char != '\n' {
			cleanedMsg += string(char)
		}
	}
	return strings.TrimSpace(cleanedMsg)
}
