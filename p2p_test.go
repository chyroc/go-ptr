package ptr

import (
	"testing"

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
