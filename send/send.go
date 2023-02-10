package send

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/logger"
	"github.com/andey-robins/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to string, user string) {
	if !db.UserExists(user) {
		logger.Log(8, user) //log sender not valid
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		logger.Log(9, user) //log sender pw not valid
		log.Fatalf("Unable to authenticate user")
	}

	if !db.UserExists(to) {
		logger.Log(5, to)
		log.Fatalf("Destination user does not exist")
	}

	message := getUserMessage()
	logger.Log(1, to)
	db.SaveMessage(message, to)
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {
	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return text
}
