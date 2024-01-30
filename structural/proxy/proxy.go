package proxy

import "fmt"

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (u *UserListProxy) FindUser(id int32) (User, error) {

	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)

	fmt.Println("returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}
