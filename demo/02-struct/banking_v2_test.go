package main

import (
	"testing"
)

func TestBanking(t *testing.T) {
	assertBalance := func(t testing.TB, b Banking, want THB) {
		t.Helper()
		got := b.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		b := Banking{}
		b.Deposit(THB(1000))
		assertBalance(t, b, THB(1000))
	})

	t.Run("withdraw", func(t *testing.T) {
		b := Banking{balance: 5000}
		b.Withdraw(THB(1000))
		assertBalance(t, b, THB(4000))
	})

	t.Run("withdraw not enough money", func(t *testing.T) {
		startingBalance := THB(1000)
		b := Banking{startingBalance}
		err := b.Withdraw(THB(2000))

		assertBalance(t, b, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

}
