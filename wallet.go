package learn_go_with_tests

import (
	"errors"
	"fmt"
)
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() (i Bitcoin){
	return w.balance
}

func (w *Wallet) Withdraw(i Bitcoin) (error){
	if i > w.balance {
		return InsufficientFundsError
	}
	w.balance -= i
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}