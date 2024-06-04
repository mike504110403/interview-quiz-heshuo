package quiz5

import (
	"fmt"
	"math"
)

var units = []string{"K", "M", "B", "T"}
var unitThresholds = []float64{1e3, 1e6, 1e9, 1e12}

// sol1 : 透長度判斷
func Format(num float64) string {
	// 轉數字
	// todo: 大位數格式化調整
	converted := fmt.Sprintf("%.16g", num)
	// 不換算直接回傳
	if num < 1e3 || num > 1e15 {
		return converted
	}
	convertUnit := func(floatNum float64) (converted string) {
		for i := len(units) - 1; i >= 0; i-- {
			unit := unitThresholds[i]
			if num >= unit {
				// 格式化數字並返回結果
				return formatWithUnit(num, unit, units[i])
			}
		}
		return converted
	}(num)

	return convertUnit
}

// formatWithUnit 使用指定的單位格式化數字
func formatWithUnit(num, unit float64, unitStr string) string {
	// 計算格式化後的數字
	formattedNum := num / unit

	// 格式化數字
	return fmt.Sprintf("%.3g%s", formattedNum, unitStr)
}

// sol2 : 透過縮放
func PadingConvert(num float64) string {
	units = []string{"", "K", "M", "B", "T"}
	converted := ""
	// 不換算部分直接轉字串
	if num < 1e3 || num > 1e15 {
		converted = fmt.Sprintf("%.16g", num)
		return converted
	}

	// 逐步縮放數值，直到小於 1000
	exp := 0
	for num >= 1000 {
		num /= 1000
		exp++
	}
	unit := units[exp]
	converted += fmt.Sprintf("%.3g", num)
	converted += unit
	return converted
}

// sol3 : 透過math包
func FormatNumberByMath(num float64) string {
	units = []string{"", "K", "M", "B", "T"}
	converted := ""

	// 不換算部分直接轉字串
	if num < 1e3 || num > 1e15 {
		converted = fmt.Sprintf("%.16g", num)
		return converted
	}
	exp := int(math.Log10(math.Abs(num)) / 3)    // 計算指數
	scaled := num / math.Pow(10, float64(exp*3)) // 計算縮放後的數值
	return fmt.Sprintf("%.3g%s", scaled, units[exp])
}
