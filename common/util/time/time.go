package time

import "time"

func Parse(t time.Time) time.Time {
	parse, _ := time.Parse("2006-01-02 15:04:05", time.Unix(t.Unix(), 0).Format("2006-01-02 15:04:05"))
	return parse
}
