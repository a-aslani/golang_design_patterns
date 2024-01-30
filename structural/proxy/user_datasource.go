package proxy

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserRepository interface {
	FindUser(ID int32) (User, error)
}

type UserList []User

func (u *UserList) FindUser(ID int32) (User, error) {

	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == ID {
			return (*u)[i], nil
		}
	}

	return User{}, fmt.Errorf("user %d could not be found", ID)
}

func (u *UserList) addUser(newUser User) {
	*u = append(*u, newUser)
}
