package learn_go_with_tests

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func (w Wallet, want Bitcoin){
		got := w.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func (err error, expectedErr error){
		if err != expectedErr {
			t.Errorf("should throw out err when over drafting, got %s, but expect %s", err, expectedErr)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		fmt.Println("address of balance in test is", &wallet.balance)

		assertBalance(wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(6))
		assertBalance(wallet, Bitcoin(4))
		assertError(err, nil)
	})

	t.Run("over drafting", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(20))
		assertBalance(wallet, Bitcoin(10))
		expectedError := InsufficientFundsError
		assertError(err, expectedError)
	})

}
