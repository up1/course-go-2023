package main

import "fmt"

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

func (b *Banking) Withdraw(amount THB) {
	b.balance -= amount
}
