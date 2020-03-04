package date

import "time"

const fmtYYYY = `2006`

type YYYY int32

func (d YYYY) Time() time.Time {
	return time.Date(int(d), time.January, 1, 0, 0, 0, 0, time.UTC)
}

func (d YYYY) String() string {
	return d.Time().Format(fmtYYYY)
}
