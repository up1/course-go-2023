package main

import (
	"errors"
	"fmt"
)

type THB int

func (thb THB) String() string {
	return fmt.Sprintf("%d THB", thb)
}

type Banking struct {
	balance THB
}

func (b *Banking) Balance() THB {
	return b.balance
}

func (b *Banking) Deposit(amount THB) {
	b.balance += amount
}

var ErrNotEnoughMoney = errors.New("Cannot withdraw, money not enough")

func (b *Banking) Withdraw(amount THB) error {
	if amount > b.balance {
		return ErrNotEnoughMoney
	}

	b.balance -= amount
	return nil
}
