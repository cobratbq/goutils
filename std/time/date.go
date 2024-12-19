package time

import (
	"strings"
	"time"
	"unicode"

	"github.com/cobratbq/goutils/codec/bytes/digit"
	"github.com/cobratbq/goutils/std/errors"
	"github.com/cobratbq/goutils/std/log"
	"github.com/cobratbq/goutils/std/strconv"
	strings_ "github.com/cobratbq/goutils/std/strings"
)

// FIXME provide convert to time.Duration type?
type DateDuration struct {
	Days   int
	Months int
	Years  int
}

func (d *DateDuration) AddToTime(t *time.Time) time.Time {
	return t.AddDate(int(d.Years), int(d.Months), int(d.Days))
}

func (d *DateDuration) Add(o *DateDuration) {
	d.Days += o.Days
	d.Months += o.Months
	d.Years += o.Years
}

// ParseDateDuration is similar to ParseDuration but parses years, months, days. This is useful in `AddDate`
// function, that can processes larger durations.
//
// Valid suffixes: y, year, years, m, month, months, w, week, weeks, d, day, days
//
// An error is returned iff no valid duration could be parsed from the provided text.
//
// Returns DateDuration, error
func ParseDateDuration(text string) (DateDuration, error) {
	if len(text) == 0 || !digit.IsDigit(text[0]) {
		return DateDuration{}, errors.ErrIllegal
	}
	switch {
	case strings_.AnySuffix(text, "y", "year", "years"):
		digits := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(text, "y"), "year"), "years")
		years, err := strconv.ParseInt[int](digits, 10)
		if err != nil {
			return DateDuration{}, errors.Context(errors.ErrIllegal, "Illegal 'year' duration")
		}
		return DateDuration{Years: years}, nil
	case strings_.AnySuffix(text, "m", "month", "months"):
		digits := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(text, "m"), "month"), "months")
		months, err := strconv.ParseInt[int](digits, 10)
		if err != nil {
			return DateDuration{}, errors.Context(errors.ErrIllegal, "Illegal 'month' duration")
		}
		return DateDuration{Months: months}, nil
	case strings_.AnySuffix(text, "w", "week", "weeks"):
		digits := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(text, "w"), "week"), "weeks")
		weeks, err := strconv.ParseInt[int](digits, 10)
		if err != nil {
			return DateDuration{}, errors.Context(errors.ErrIllegal, "Illegal 'week' duration")
		}
		return DateDuration{Days: weeks * 7}, nil
	case strings_.AnySuffix(text, "d", "day", "days"):
		digits := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(text, "d"), "day"), "days")
		days, err := strconv.ParseInt[int](digits, 10)
		if err != nil {
			return DateDuration{}, errors.Context(errors.ErrIllegal, "Illegal 'day' duration")
		}
		return DateDuration{Days: days}, nil
	default:
		return DateDuration{}, errors.ErrIllegal
	}
}

// ParseDateDurations is similar to ParseDuration but parses years, months, days. This is useful in `AddDate`
// function, that can processes larger durations. ParseDateDurations splits input string in fields and parses
// each individual field. Fields that produce an error are skipped, incrementing the number of failures.
//
// If failures == 0, all fields were properly parsed with no surprises, meaning that for each field a suitable
// suffix was recognized and a decimal value processed. For failures > 0, some --possibly all-- fields failed
// to process.
//
// Valid suffixes: y, year, years, m, month, months, w, week, weeks, d, day, days
// Valid separators: any whitespace (see `unicode.IsSpace`), ',', ';'
//
// Returns DateDuration, number of failures
func ParseDateDurations(s string) (DateDuration, uint) {
	parts := strings.FieldsFunc(s, func(c rune) bool {
		return c == ',' || c == ';' || unicode.IsSpace(c)
	})
	var duration DateDuration
	var failures uint
	for p := range parts {
		if d, err := ParseDateDuration(parts[p]); err == nil {
			duration.Add(&d)
		} else {
			log.Traceln("Skipping part", parts[p], "with invalid duration:", err.Error())
			failures++
		}
	}
	return duration, failures
}
