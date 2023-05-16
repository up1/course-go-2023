package main

import (
	"testing"
)

func TestBankingDeposit(t *testing.T) {
	// Arrange
	banking := Banking{}
	// Act
	banking.Deposit(THB(1000))
	got := banking.Balance()
	// Assert
	want := THB(1000)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestBankingWithdraw(t *testing.T) {
	// Arrange
	banking := Banking{
		balance: 5000,
	}
	// Act
	banking.Withdraw(THB(1000))
	got := banking.Balance()
	// Assert
	want := THB(4000)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
