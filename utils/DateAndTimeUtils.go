package utils

import (
	"fmt"
	"strconv"

	"time"
)

func GetYears(difference time.Duration) int64 {
	return int64(difference.Hours() / 24 / 365)
}

func GetMonths(difference time.Duration) int64 {
	return int64(difference.Hours() / 24 / 30)
}

func GetWeeks(difference time.Duration) int64 {
	return int64(difference.Hours() / 24 / 7)
}

func GetDays(difference time.Duration) int64 {
	return int64(difference.Hours() / 24)
}

func GetHours(difference time.Duration) (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%.f", difference.Hours()), 10, 64)
}

func GetMinutes(difference time.Duration) (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%.f", difference.Minutes()), 10, 64)
}

func GetSeconds(difference time.Duration) (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%.f", difference.Seconds()), 10, 64)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

var daysBefore = [...]int32{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

func daysIn(m time.Month, year int) int {
	if m == time.February && isLeap(year) {
		return 29
	}
	return int(daysBefore[m] - daysBefore[m-1])
}

func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}
func AddDateX(t time.Time, years int, months int, days int, normalizes ...bool) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	if len(normalizes) > 0 && normalizes[0] == true {
		return time.Date(year+years, month+time.Month(months), day+days, hour, min, sec, 0, t.Location())
	}

	// Normalize month, overflowing into year.
	m := int(month+time.Month(months)) - 1
	year, m = norm(year+years, m, 12)
	month = time.Month(m) + 1

	if lDay := daysIn(month, year); day > lDay {
		// If the desired month does not have this day,
		// temporarily set the day as the maximum day of the desired month,
		// and then process the param days.
		day = lDay
	}
	return time.Date(year, month, day+days, hour, min, sec, 0, t.Location())
}
