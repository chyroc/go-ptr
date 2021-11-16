package ptr

import "time"

func PointerInt64ToTime(i *int64) *time.Time {
	if i == nil {
		return nil
	}
	return Time(time.Unix(*i, 0))
}

func PointerTimeToInt64(i *time.Time) *int64 {
	if i == nil {
		return nil
	}
	return Int64(i.Unix())
}
