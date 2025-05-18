package main

import "fmt"

type messageToSend struct {
	message   string
	sender    sender
	recipient sender
}

type sender struct {
	rateLimit int
	user
}

type user struct {
	phoneNumber int
	contactName string
}

func canSend(m messageToSend) bool {
	if m.sender.phoneNumber == 0 {
		return false
	}
	if m.recipient.phoneNumber == 0 {
		return false
	}
	if m.sender.contactName == "" {
		return false
	}
	if m.recipient.contactName == "" {
		return false
	}
	if m.message == "" {
		return false
	}
	if m.sender.phoneNumber == m.recipient.phoneNumber {
		return false
	}
	if m.sender.rateLimit <= 0 {
		return false
	}

	return true
}

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("User %s: Password %s", authI.username, authI.password)
}

func testAuthentication(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func test(m messageToSend) {
	if canSend(m) {
		fmt.Printf("Sending %s to %v from %v\n", m.message, m.recipient.contactName, m.sender.contactName)
		fmt.Println("====================================")
	} else {
		fmt.Printf("Cannot send %s to %v from %v\n", m.message, m.recipient.contactName, m.sender.contactName)
		fmt.Println("====================================")
	}
}

func main() {
	senders := sender{rateLimit: 1000, user: user{phoneNumber: 1234567890, contactName: "Alice"}}
	recipient := sender{rateLimit: 1000, user: user{phoneNumber: 9876543210, contactName: "Bob"}}
	message := messageToSend{message: "Hello, Bob!", sender: senders, recipient: recipient}

	testAuthentication(
		authenticationInfo{
			username: "user",
			password: "password",
		},
	)
	test(message)
}
