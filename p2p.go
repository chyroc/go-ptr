package ptr

import "time"

// PointerInt64ToTime converts a pointer of type *int64 to a time.Time.
func PointerInt64ToTime(i *int64) *time.Time {
	if i == nil {
		return nil
	}
	return Time(time.Unix(*i, 0))
}

// PointerTimeToInt64 converts a pointer of type *time.Time to a int64.
func PointerTimeToInt64(i *time.Time) *int64 {
	if i == nil {
		return nil
	}
	return Int64(i.Unix())
}
