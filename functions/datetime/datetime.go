// Package datetime provides methods for manipulating dates and times in templates
package datetime

import (
	"text/template"
	"time"

	"github.com/mantidtech/tplr/functions/helper"
)

// Format is the default format to use for timestamps
const Format = time.RFC3339

// Functions operate on time and dates
func Functions() template.FuncMap {
	return template.FuncMap{
		"now":        Now,
		"timeFormat": TimeFormat,
		"timeParse":  TimeParse,
		"timeToUnix": TimeToUnix,
		"unixToTime": UnixToTime,
	}
}

// Now returns the current time in the format "2006-01-02T15:04:05Z07:00" (RFC3339)
func Now(format ...string) string {
	f := Format
	if len(format) > 0 {
		f = format[0]
	}
	return helper.Now().Format(f)
}

// TimeFormat formats the given timestamp with the given format
func TimeFormat(format string, ts time.Time) string {
	return ts.Format(format)
}

// TimeParse parses the given string using the given format
func TimeParse(format string, ts string) (time.Time, error) {
	return time.Parse(format, ts)
}

// TimeToUnix converts the given timestamp to the number of seconds since the unix epoch
func TimeToUnix(ts time.Time) int64 {
	return ts.Unix()
}

// UnixToTime converts the given number of seconds, as a unix epoch, to a timestamp
func UnixToTime(s int) time.Time {
	return time.Unix(int64(s), 0)
}

//
// func TimeToUnixMS(ts time.Time) int64 {
// 	return ts.UnixNano() / 1_000_000
// }
//
// func UnixToTimeMS(ms int) time.Time {
// 	return time.Unix(0, int64(ms)*1_000_000)
// }
//
// func TimeToUnixNS(ts time.Time) int64 {
// 	return ts.UnixNano()
// }
//
// func UnixToTimeNS(s int) time.Time {
// 	return time.Unix(0, int64(s))
// }
