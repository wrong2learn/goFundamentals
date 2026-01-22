package services

type LoggingNotifier struct {
	Notifier
}

func (l LoggingNotifier) Notify(msg string) error {
	println("[LOG]", msg)
	return l.Notifier.Notify(msg)
}
