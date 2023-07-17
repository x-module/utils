/**
 * Created by Goland.
 * @file   decimal.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/14 13:48
 * @desc   decimal.go
 */

package types

import (
	"github.com/shopspring/decimal"
	"strconv"
)

type Decimal struct {
	decimal.Decimal
}

func NewFromFloat(value float64) Decimal {
	return Decimal{
		Decimal: decimal.NewFromFloat(value),
	}
}

func (d *Decimal) Float64() float64 {
	fl, _ := d.Decimal.Float64()
	return fl
}
func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(d.Float64(), 'f', 2, 64)), nil
}

func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		return nil
	}
	dec, err := decimal.NewFromString(string(decimalBytes))
	if err != nil {
		return err
	}
	*d = Decimal{
		Decimal: dec,
	}
	return nil
}
