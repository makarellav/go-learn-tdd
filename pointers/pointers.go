package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(depositAmount Bitcoin) {
	w.balance += depositAmount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(withdrawAmount Bitcoin) error {
	if withdrawAmount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= withdrawAmount

	return nil
}
