package helper
import "strings"

func ValidateUserInput(userFirstName string, userLastName string, userMailID string, userTickets uint,remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userMailID, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
