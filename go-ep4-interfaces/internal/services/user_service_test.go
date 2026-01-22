package services

import "testing"

func TestUserRegistration(t *testing.T) {
	mock := &MockNotifier{}
	service := NewUserService(mock)

	service.RegisterUser("Alice")

	if len(mock.Messages) != 1 {
		t.Fatal("expected notification")
	}
}
