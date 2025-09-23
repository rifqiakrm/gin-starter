//revive:disable:var-naming
package helper

import (
	"database/sql"
	"fmt"
	"time"
)

// StringToNullString convert string to sql null string
func StringToNullString(d string) sql.NullString {
	if d == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return sql.NullString{
		String: d,
		Valid:  true,
	}
}

// BoolToNullBool convert bool to sql null bool
func BoolToNullBool(d bool) sql.NullBool {
	return sql.NullBool{
		Bool:  d,
		Valid: true,
	}
}

// Float64ToNullFloat64 convert float64 to sql null float64
func Float64ToNullFloat64(d float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: d,
		Valid:   true,
	}
}

// Int32ToNullInt32 convert int32 to sql null int32
func Int32ToNullInt32(d int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: d,
		Valid: true,
	}
}

// Int64ToNullInt64 convert int64 to sql null int64
func Int64ToNullInt64(d int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: d,
		Valid: true,
	}
}

// TimeToNullTime convert time to sql null time
func TimeToNullTime(d time.Time) sql.NullTime {
	if d.IsZero() {
		return sql.NullTime{
			Time:  d,
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  d,
		Valid: true,
	}
}

// DateStringToTime convert date string to time
func DateStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}

	layout := "2006-01-02"
	t, errParse := time.Parse(layout, date)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t, nil
}

// DateTimeStringToTime convert date string to time
func DateTimeStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}

	layout := "2006-01-02 15:04:05"
	t, errParse := time.Parse(layout, date)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t, nil
}

// ImageFullPath define image full path
func ImageFullPath(imgHost string, path string) string {
	return fmt.Sprintf("%s%s", imgHost, path)
}
