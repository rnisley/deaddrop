# deaddrop-go

A deaddrop utility written in Go. Put files in a database behind a password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:
- [Javascript](https://github.com/andey-robins/deaddrop-js)
- [Rust](https://github.com/andey-robins/deaddrop-rs)

## Versioning

`deaddrop-go` is built with:
- go version go1.19.4 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new -user <username here>` and you will be prompted to create the initial password.

## Database

Data gets stored into the local database file dd.db. This file will not by synched to git repos. Delete this file if you don't set up a user properly on the first go

## Logging Strategy

Added logger directory with logger.go in it

Added logging for 
• sending and reading a message to a user that exists
• creating a new user
• reading messages with the wrong password
• reading the messages for a user which doesn’t exist
• sending messages to a user which doesn’t exist
• creating a user with an account that doesn't exist
• creating a user with an invalid password

## Mitigation

Added logger discussed in HW 1
DB backup solution would require an external application to be developed. 
Discussion with Andey revealed that any valid user can create additional users so that's been addressed already.

Added user login requirement for sending messages