package timeutil

import (
	"testing"
	"time"
)

func TestAtStartOfDay(t *testing.T) {
	start, _ := time.Parse("2006-01-02 15:04:05", "2016-02-02 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05.999999999", "2016-02-02 23:59:59.999999999")
	exp, _ := time.Parse("2006-01-02 15:04:05.999999999", "2016-02-02 00:00:00.000000000")

	for start.Before(end) {
		if diff := exp.Sub(AtStartOfDay(start)); diff != 0 {
			t.Errorf("Expected %d, but was %d", 0, diff)
		}
		start = start.Add(time.Second)
	}
}

func TestAtEndOfDay(t *testing.T) {
	start, _ := time.Parse("2006-01-02 15:04:05", "2016-02-02 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05.999999999", "2016-02-02 23:59:59.999999999")
	exp, _ := time.Parse("2006-01-02 15:04:05.999999999", "2016-02-02 23:59:59.999999999")

	for start.Before(end) {
		if diff := exp.Sub(AtEndOfDay(end)); diff != 0 {
			t.Errorf("Expected %d, but was %d", 0, diff)
		}
		start = start.Add(time.Second)
	}
}

func TestAtTimeOfDay(t *testing.T) {
	dt, _ := time.Parse("2006-01-02 15:04:05", "2016-02-02 00:00:00")
	exp, _ := time.Parse("2006-01-02 15:04:05.999999999", "2016-02-02 23:59:59.999999999")
	if diff := exp.Sub(AtTime(dt, 23, 59, 59, 999999999)); diff != 0 {
		t.Errorf("Expected %d, but was %d", 0, diff)
	}
}

func TestAddDateDay(t *testing.T) {
	// start + dur = exp
	type rec struct {
		start string
		year  int
		mon   int
		day   int
		exp   string
	}
	tests := []rec{
		rec{"2010-01-28", 0, 1, 0, "2010-02-28"},
		rec{"2010-01-29", 0, 1, 0, "2010-02-28"},
		rec{"2010-01-30", 0, 1, 0, "2010-02-28"},
		rec{"2010-02-28", 0, -1, 0, "2010-01-28"},
		// add day to mix
		rec{"2010-01-28", 0, 1, 1, "2010-03-01"},
		rec{"2010-01-29", 0, 1, 1, "2010-03-01"},
		rec{"2010-03-01", 0, -1, -1, "2010-01-31"},
		// leap years
		rec{"2011-01-29", 1, 1, 0, "2012-02-29"},
		rec{"2011-01-30", 1, 1, 0, "2012-02-29"},
		rec{"2011-01-31", 1, 1, 0, "2012-02-29"},
		rec{"2013-03-30", -1, -1, -10, "2012-02-19"},
		rec{"2012-03-30", -1, -1, -10, "2011-02-18"},
		// year turn
		rec{"2010-12-31", 0, 0, 1, "2011-01-01"},
		rec{"2011-01-01", 0, 0, -1, "2010-12-31"},
		// 12 months
		rec{"2010-12-31", 0, 12, 0, "2011-12-31"},
		rec{"2010-12-31", 0, 26, 0, "2013-02-28"},
		// 12 months + years
		rec{"2010-12-31", 1, 12, 0, "2012-12-31"},
		rec{"2010-12-31", 3, 26, 0, "2016-02-29"},
	}

	for _, r := range tests {
		dt, _ := time.Parse("2006-01-02", r.start)
		act := AddDate(dt, r.year, r.mon, r.day)
		exp, _ := time.Parse("2006-01-02", r.exp)
		if exp != act {
			t.Errorf("Expected %s, but was %s (%s, %d, %d, %d)", exp, act, r.start, r.year, r.mon, r.day)
		}
	}
}
