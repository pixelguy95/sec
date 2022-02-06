package sec

import "time"

const SECDateFormat = "2006-01-02"

func StandardSecDateFormatParse(date string) (time.Time, error) {
	return time.Parse(SECDateFormat, date)
}
