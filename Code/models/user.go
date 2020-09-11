package models

import (
	"errors"
	"fmt"
)

//User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers func
func GetUsers() []*User {

	return users
}

//AddUser func
func AddUser(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("Neww user must not include ID or  it must be set to 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID func
func GetUserByID(ID int) (User, error) {

	for _, u := range users {
		if u.ID == ID {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", ID)
}

//UpdateUser func
func UpdateUser(u User) (User, error) {

	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUser func
func RemoveUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
