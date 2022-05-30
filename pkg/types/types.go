package types

type Money int

type Currency string

type PAN string

type Payment struct {
	ID int
	Amount Money
}

type Card struct{
	Balance Money
	Currency Currency
	Color string
	Active bool
}
const(
	Platinum = "Platinum"
	Gold = "Gold"
	Silver = "Silver"
)

const(
	Somoni =  "TJS"
	Rubls =   "RUB"
	Dollars = "USD"
)
