package account

type account struct {
	amount int
}

type AccountAccess interface {
	Credit(amount int) bool
	Debit(amount int) bool
	Balance() int
}

func NewAccountAccess() AccountAccess {
	return &account{amount: 0}
}

func (acc *account) Credit(amount int) bool {
	acc.amount += amount
	return true
}

func (acc *account) Debit(amount int) bool {
	acc.amount -= amount
	return true
}

func (acc *account) Balance() int {
	return acc.amount
}
