package banking

type Banking struct {
	V int
}

var dd int

func (b *Banking) Balance() int {
	return b.V
}

func (b *Banking) Deposit(amount int) {
	b.V += amount
}

func (b *Banking) Withdraw(amount int) {
	b.V -= amount
}
