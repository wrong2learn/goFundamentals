package models

type Wallet struct {
	Balance float64
}

// Value receiver (copy)
func (w Wallet) Add(amount float64) {
	w.Balance += amount
}

// Pointer receiver (reference)
func (w *Wallet) AddPtr(amount float64) {
	w.Balance += amount
}
