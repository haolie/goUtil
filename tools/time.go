package tools

import (
	"fmt"
	"time"
)

func GetDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func ToTimeStr(t time.Time) string {
	return fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
}

func ToDateStr(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%d-%d", y, m, d)
}

func ToDateTimeStr(t time.Time) string {
	return fmt.Sprintf("%s %s", ToDateStr(t), ToTimeStr(t))
}
