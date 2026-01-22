package services

type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) UserService {
	return UserService{notifier: n}
}

func (s UserService) RegisterUser(name string) {
	s.notifier.Notify("User registered: " + name)
}
