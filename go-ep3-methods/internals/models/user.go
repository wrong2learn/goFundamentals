package models

type User struct {
	Name string
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}
