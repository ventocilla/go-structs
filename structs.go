package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// u is the parameter name  / user is the value type
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	UserBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user

	appUser = newUser(userFirstName, userLastName, UserBirthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
