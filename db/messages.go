package db

import (
	"crypto/sha256"
	"log"
	"github.com/andey-robins/deaddrop-go/logger"
)

// GetMessagesForUser assumes that a user has already been
// authenticated through a call to session.Authenticate(user)
// and then returns all the messages stored for that user
func GetMessagesForUser(user string) []string {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT * FROM Messages
		WHERE recipient = (
			SELECT id FROM Users WHERE user = ?
		)
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	// marshall rows into an array
	messages := make([]string, 0)
	
	for rows.Next() {
		var message string
		var messageMAC string
		var id int
		var sender string
		var recipient string
		err := rows.Scan(&id, &recipient, &message, &messageMAC, &sender)
		if err != nil {
			log.Fatalf("unable to scan row")
		}
		if !ValidMAC(message, messageMAC) {
			message = "This message cannot be authenticated: " + message
			logger.NoAuth(user)
		}
		message = "Message from " + sender + ": " + message
		
		messages = append(messages, message)
	}
	return messages
}

// saveMessage will process the transaction to place a message
// into the database
func SaveMessage(message string, recipient string, sender string) {
	database := Connect().Db

	mac := sha256.New()
	mac.Write([]byte(message))
	newMac := mac.Sum(nil)
	database.Exec(`
		INSERT INTO Messages (recipient, data, mac, sender)
		VALUES (
			(SELECT id FROM Users WHERE user = ?), 
			?, ?, ?
		);
	`, recipient, message, string(newMac), sender)
}

func ValidMAC(message string, messageMAC string) bool {
	mac := sha256.New()
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return (messageMAC == string(expectedMAC))
}
