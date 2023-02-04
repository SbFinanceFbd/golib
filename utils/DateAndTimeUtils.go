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
