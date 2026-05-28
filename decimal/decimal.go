package decimal

import "fmt"

type Decimal interface {
	ToTHB() string
}

type decimal struct {
	price   float64
	digits  []string
	numbers []string
}

func NewFromFloat(price float64) Decimal {
	return decimal{
		price:   price,
		numbers: []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"},
		digits:  []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"},
	}
}

func (d decimal) String() string {
	return fmt.Sprintf("%g", d.price)
}

func (d decimal) ToTHB() string {
	integer := int(d.price)
	satang := int((d.price-float64(integer))*100 + 0.5)

	res := d._convert(integer) + "บาท"

	if satang == 0 {
		return res + "ถ้วน"
	}

	return res + d._convert(satang) + "สตางค์"
}

func (d decimal) _convert(n int) string {
	if n == 0 {
		return d.numbers[0]
	}

	res := ""

	nums := []int{}
	for ; n > 0; n /= 10 {
		nums = append([]int{n % 10}, nums...)
	}

	length := len(nums)

	for i, num := range nums {
		if num == 0 {
			continue
		}

		pos := length - i - 1

		switch pos {
		// avoid "หนึ่ง" in ending digits of 10s
		case 0:
			if num == 1 && length > 1 {
				res += "เอ็ด"
			} else {
				res += d.numbers[num]
			}

		case 1:
			switch num {
			// avoid "หนึ่งสิบ" and "สองสิบ" in 10s
			case 1:
				res += "สิบ"
			case 2:
				res += "ยี่สิบ"
			default:
				res += d.numbers[num] + "สิบ"
			}

		default:
			res += d.numbers[num] + d.digits[pos]
		}
	}

	return res
}
