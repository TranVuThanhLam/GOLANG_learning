package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name and birthday is require!")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.lastName, " ", u.firstName, " ", u.birthdate, " ", u.createAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
