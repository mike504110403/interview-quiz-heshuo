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

	return convertWithUnit(num)
}

// convertWithUnit 根據數字的大小進行單位換算
func convertWithUnit(num float64) string {
	for i := len(units) - 1; i >= 0; i-- {
		unit := unitThresholds[i]
		if num >= unit {
			// 格式化數字並返回結果
			return formatWithUnit(num, unit, units[i])
		}
	}

	return ""
}

// formatWithUnit 使用指定的單位格式化數字
func formatWithUnit(num, unit float64, unitStr string) string {
	// 計算格式化後的數字
	formattedNum := num / unit

	// 格式化數字
	return fmt.Sprintf("%.3g%s", formattedNum, unitStr)
}

// sol2 : 透過縮放 位移
func PadingConvert(num float64) string {
	units = []string{"", "K", "M", "B", "T"}
	converted := ""
	// 不換算部分直接轉字串
	if num < 10^3 || num > 10^15 {
		converted = fmt.Sprintf("%.15g", num)
		return converted
	}

	// 逐步縮放數值，直到小於 1000
	exp := 0
	for num >= 1000 {
		num /= 1000
		exp++
	}
	unit := units[exp]
	converted += fmt.Sprintf("%f", num/float64(int64(1)*int64(1<<uint(exp*10))))
	converted += unit
	return converted
}

// sol3 : 透過math包
func FormatNumberByMath(num float64) string {
	units = []string{"", "K", "M", "B", "T"}
	converted := ""

	// 不換算部分直接轉字串
	if num < 10^3 || num > 10^15 {
		converted = fmt.Sprintf("%.15g", num)
		return converted
	}
	exp := int(math.Log10(math.Abs(num)) / 3)    // 計算指數
	scaled := num / math.Pow(10, float64(exp*3)) // 計算縮放後的數值
	return fmt.Sprintf("%.1f%s", scaled, units[exp])
}
