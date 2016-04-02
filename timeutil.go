// Package timeutil uses a different arithmetic for time.Time and has some utilities functions.
package timeutil

import "time"

// AddDate adds years, months and days to t, similar to time.AddDate.
//
// AddDate truncates the result, instead of nomalizing like the time.AddDate,
// so, for example, adding one month to October 31 yields November 30,
// truncated from November 31.
func AddDate(t time.Time, years, months, days int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	month-- // make months as zero-based index for modulo operation
	years += (int(month) + months) / 12
	tm := (month + time.Month(months)) % 12
	if mdays[month] > mdays[tm] && day+days > mdays[tm] {
		month--
	}

	// truncate
	if !IsLeap(year+years) && day > mdays[tm] {
		day = mdays[tm]
	} else if day > mdaysLeap[tm] {
		day = mdaysLeap[tm]
	}

	tm++ // revert to 1-based index
	return time.Date(year+years, tm, day+days, hour, min, sec, t.Nanosecond(), t.Location())
}

// AtStartOfDay returns t with zeroed time parts (00:00:00.000000000).
func AtStartOfDay(t time.Time) time.Time {
	d := time.Duration(-t.Hour()) * time.Hour
	return t.Add(d).Truncate(time.Hour)
}

// AtStartOfDay returns t at its last moment of the day (23:59:59.999999999)
func AtEndOfDay(t time.Time) time.Time {
	d := time.Duration(24-t.Hour()) * time.Hour
	return t.Add(d).Truncate(time.Hour).Add(-time.Nanosecond)
}

// AtTime returns t with the time parts set to hour, min, sec and nsec.
func AtTime(t time.Time, hour, min, sec, nsec int) time.Time {
	h := time.Duration(hour) * time.Hour
	m := time.Duration(min) * time.Minute
	s := time.Duration(sec) * time.Second
	n := time.Duration(nsec) * time.Nanosecond
	delta := h + m + s + n
	return AtStartOfDay(t).Add(delta)
}

// Check if year is a leap year.
// Same code as the standard year, but public.
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

var mdays = []int{
	31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

var mdaysLeap = []int{
	31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}
