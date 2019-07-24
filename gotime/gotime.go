package gotime

import "time"

const (
	SingaporeLocal = "Asia/Singapore"
	LayoutTime     = "2006-01-02 15:04:05"
)

func Local(name string) (*time.Location, error) {
	l, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return nil, err
	}
	return l, nil
}

func ParseTimeInLocation(value string) (time.Time, error) {
	local, err := Local(SingaporeLocal)
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation(LayoutTime, value, local)
}

func FormatLocation(t time.Time) string {
	return t.Format(LayoutTime)
}

func FormatTimestamp(i int64) time.Time {
	return time.Unix(0, i*int64(time.Millisecond))
}

func ToMillisecond(t time.Time) int64 {
	return t.Local().UnixNano() / int64(time.Millisecond)
}
