package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func ParseTime(timeStr string, layout string, timezone string) (time.Time, error) {
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func MinutesToFormatted(minutes int) string {
	dur := time.Duration(minutes) * time.Minute
	return fmt.Sprintf("%vh %vm", int(dur.Hours()), int(dur.Minutes())%60)
}

func CalculateDelay(min, max, te int) time.Duration {
	random := rand.Intn(max-min) + min
	return time.Duration(random-te) * time.Millisecond
}
