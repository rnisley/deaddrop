package read

import (
	"fmt"
	"log"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/logger"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {
	if !db.UserExists(user) {
		logger.Log(4, user)
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		logger.Log(3, user)
		log.Fatalf("Unable to authenticate user")
	}

	logger.Log(0, user)
	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		fmt.Println(message)
	}
}
