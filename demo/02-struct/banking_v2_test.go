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

}
