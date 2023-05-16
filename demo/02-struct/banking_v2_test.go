package main

import (
	"testing"
)

func TestBanking(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		b := Banking{}
		b.Deposit(THB(1000))
		assertBalance(t, b, THB(1000))
	})

	t.Run("withdraw", func(t *testing.T) {
		b := Banking{balance: 5000}
		err := b.Withdraw(THB(1000))
		assertBalance(t, b, THB(4000))
		assertNoError(t, err)
	})

	t.Run("withdraw not enough money", func(t *testing.T) {
		startingBalance := THB(1000)
		b := Banking{startingBalance}
		err := b.Withdraw(THB(2000))

		assertBalance(t, b, startingBalance)
		assertError(t, err, ErrNotEnoughMoney)
	})
}

func assertBalance(t testing.TB, b Banking, want THB) {
	t.Helper()
	got := b.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
