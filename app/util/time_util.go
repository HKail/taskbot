package util

import "time"

type YearType int

const (
	YearTypeLeap  YearType = iota // 闰年
	YearTypePeace                 // 平年
)

var (
	daysOfMonthInLeapYear  = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	daysOfMonthInPeaceYear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

// GetYearmonth 获取年月数值, 如: 200601
func GetYearmonth(t time.Time) int {
	return t.Year()*100 + int(t.Month())
}

// GetDateNumberVal 获取日期数值, 如: 20060102
func GetDateNumberVal(t time.Time) int {
	return GetYearmonth(t)*100 + t.Day()
}

// GetDayMask 获取当天日期 bitmap 掩码
func GetDayMask(t time.Time) uint32 {
	return 1 << (t.Day() - 1)
}

// GetDaysOfMonthInYear 获取当前年份每月的天数
func GetDaysOfMonthInYear(year int) []int {
	yearType := getYearType(year)
	if yearType == YearTypeLeap {
		return daysOfMonthInLeapYear
	}

	return daysOfMonthInPeaceYear
}

// GetDaysOfYearmonth 获取当前年月的天数
func GetDaysOfYearmonth(yearmonth int) int {
	year, month := yearmonth/100, yearmonth%100

	return GetDaysOfMonthInYear(year)[month-1]
}

func getYearType(year int) YearType {
	if year%100 != 0 && year%4 == 0 || year%400 == 0 {
		return YearTypeLeap
	}

	return YearTypePeace
}
