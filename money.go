package goutils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

// FenToYuan将金额分转换成元. 返回一个float64类型.
func FenToYuanFloat(money int) float64 {
	if money == 0 {
		return 0
	}
	yuan, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(money)/100), 64)
	return yuan
}

// FenToYuanString将金额分转换成元. 返回一个string类型.
func FenToYuanString(money int) string {
	if money == 0 {
		return "0"
	}
	if money%100 == 0 {
		return fmt.Sprintf("%.0f", float64(money)/100)
	}
	return fmt.Sprintf("%.2f", float64(money)/100)
}

// YuanToFen将金额元转换为分.
// 如果需要四舍五入，请使用YuanToFenRounding.
func YuanToFen(money float64) int64 {
	d := decimal.New(1, 2) //分转元乘以100
	//df := decimal.NewFromFloat(price).Mul(d).DivRound(d1,2).String()
	df := decimal.NewFromFloat(money).Mul(d).IntPart()
	return df
}

// YuanToFenRounding将金额元转换为分.
// 它满足：当乘以100后，仍然有小数位，四舍五入法后，再取整数部分.
func YuanToFenRounding(money float64) int64 {
	// 分转元乘以100
	d := decimal.New(1, 2)
	// 乘完之后，保留2为小数，需要一个中间参数
	d1 := decimal.New(1, 0)
	dff := decimal.NewFromFloat(money).Mul(d).DivRound(d1, 0).IntPart()
	return dff
}

// DiscountCalculation计算折扣. realMoney实际价格，tagPrice吊牌价.
func DiscountCalculation(realMoney, tagPrice float64) string {
	if tagPrice <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.1f", realMoney/tagPrice*10)
}

// DiscountCalculationInt计算折扣. realMoney实际价格，tagPrice吊牌价.
func DiscountCalculationInt(realMoney, tagPrice int) string {
	if tagPrice <= 0 {
		return "0"
	}
	return DiscountCalculation(FenToYuanFloat(realMoney), FenToYuanFloat(tagPrice))
}
