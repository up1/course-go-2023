package banking_test

import (
	"banking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	// Arrange
	b := banking.Banking{
		V: 1000,
	}
	// Act
	b.Deposit(3000)
	r := b.Balance()
	// Assert
	assert.Equal(t, 4000, r)
}

func TestWithdraw(t *testing.T) {
	// Arrange
	b := banking.Banking{
		V: 1000,
	}
	// Act
	b.Withdraw(700)
	r := b.Balance()
	// Assert
	assert.Equal(t, 300, r)
}
