package decimal_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDecimal_ToTHB(t *testing.T) {
	t.Run("price without satang", func(t *testing.T) {
		tests := []struct {
			input    float64
			expected string
		}{
			{0, "ศูนย์บาทถ้วน"},
			{1, "หนึ่งบาทถ้วน"},
			{10, "สิบบาทถ้วน"},
			{11, "สิบเอ็ดบาทถ้วน"},
			{20, "ยี่สิบบาทถ้วน"},
			{21, "ยี่สิบเอ็ดบาทถ้วน"},
			{105, "หนึ่งร้อยห้าบาทถ้วน"},
			{1234, "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
			{21.00, "ยี่สิบเอ็ดบาทถ้วน"},
		}

		for _, tt := range tests {
			d := decimal.NewFromFloat(tt.input)
			result := d.ToTHB()

			assert.Equal(t, tt.expected, result)
		}
	})

	t.Run("price with satang", func(t *testing.T) {
		tests := []struct {
			input    float64
			expected string
		}{
			{33333.75, "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
			{100.01, "หนึ่งร้อยบาทหนึ่งสตางค์"},
			{100.11, "หนึ่งร้อยบาทสิบเอ็ดสตางค์"},
			{100.21, "หนึ่งร้อยบาทยี่สิบเอ็ดสตางค์"},
			{21.50, "ยี่สิบเอ็ดบาทห้าสิบสตางค์"},
		}

		for _, tt := range tests {
			d := decimal.NewFromFloat(tt.input)
			result := d.ToTHB()

			assert.Equal(t, tt.expected, result)
		}
	})

	t.Run("special thai wording rules", func(t *testing.T) {
		tests := []struct {
			input    float64
			expected string
		}{
			{11, "สิบเอ็ดบาทถ้วน"},
			{21, "ยี่สิบเอ็ดบาทถ้วน"},
			{101, "หนึ่งร้อยหนึ่งบาทถ้วน"},
			{121, "หนึ่งร้อยยี่สิบเอ็ดบาทถ้วน"},
		}

		for _, tt := range tests {
			d := decimal.NewFromFloat(tt.input)
			result := d.ToTHB()

			assert.Equal(t, tt.expected, result)
		}
	})
}
