package datetime

import (
	"text/template"
	"time"

	"github.com/mantidtech/tplr/functions/helper"
)

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

func TimeFormat(format string, ts time.Time) string {
	return ts.Format(format)
}

func TimeParse(format string, ts string) (time.Time, error) {
	return time.Parse(format, ts)
}

func TimeToUnix(ts time.Time) int64 {
	return ts.Unix()
}

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
