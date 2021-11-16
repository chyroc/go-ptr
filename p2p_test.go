package ptr

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPointerInt64ToTime(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name string
		args *int64
		want []string
	}{
		{
			name: "to-time-1617879700",
			args: nil,
			want: []string{},
		},
		{
			name: "to-time-1617879700",
			args: Int64(1617879700),
			want: []string{"2021-04-08T19:01:40", "2021-04-08T11:01:40"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PointerInt64ToTime(tt.args)
			if len(tt.want) == 0 {
				as.Nil(got)
			} else {
				as.Contains(tt.want, got.Format("2006-01-02T15:04:05"))
			}
		})
	}
}

func TestPointerTimeToInt64(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *int64
	}{
		{
			name: "nil time to nil int64",
			args: nil,
			want: nil,
		},
		{
			name: "pointer time to pointer int64",
			args: Time(time.Unix(24580, 0)),
			want: Int64(24580),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointerTimeToInt64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointerTimeToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
