package datepicker

import (
	"strconv"
	"time"
)

func intToString(i int) string {
	return strconv.Itoa(i)
}

func resolveMinYear(min int) int {
	if min == 0 {
		return time.Now().Year() - 100
	}
	return min
}

func resolveMaxYear(max int) int {
	if max == 0 {
		return time.Now().Year() + 20
	}
	return max
}

func resolvePlaceholder(p string) string {
	if p == "" {
		return "Select a date..."
	}
	return p
}

func formatDisplayDate(iso string) string {
	if iso == "" {
		return ""
	}
	t, err := time.Parse("2006-01-02", iso)
	if err != nil {
		return iso
	}
	return t.Format("Jan 2, 2006")
}
