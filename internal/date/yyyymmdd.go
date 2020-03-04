package date

import "time"

const fmtYYYYMMDD = `2006-01-02`

// YYYYMMDD provides a year, month and day type.
type YYYYMMDD time.Time

// MarshalJSON outputs JSON.
func (d YYYYMMDD) MarshalJSON() ([]byte, error) {
	return []byte(wrap(d.String())), nil
}

// UnmarshalJSON handles incoming JSON.
func (d *YYYYMMDD) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(wrap(fmtYYYYMMDD), string(b))
	if err != nil {
		return err
	}
	*d = YYYYMMDD(t)
	return nil
}

func (d YYYYMMDD) String() string {
	return time.Time(d).Format(fmtYYYYMMDD)
}
