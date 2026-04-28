package services

type MockNotifier struct {
	Messages []string
}

func (m *MockNotifier) Notify(message string) error {
	m.Messages = append(m.Messages, message)
	return nil
}
