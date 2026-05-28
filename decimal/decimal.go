package decimal

import "fmt"

type Decimal interface{}

type decimal struct {
	price float64
}

func NewFromFloat(price float64) Decimal {
	return decimal{
		price: price,
	}
}

func (d decimal) String() string {
	return fmt.Sprintf("%g", d.price)
}
