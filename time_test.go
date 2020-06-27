package backlog

import (
	"testing"
	"time"
)

var referenceTime time.Time = time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)

func TestTimestamp_String(t *testing.T) {
	tests := []struct {
		name string
		t    Timestamp
		want string
	}{
		{
			name: "set Timestamp( date in the response of backlog API )",
			t:    Timestamp{referenceTime},
			want: "2006-01-02 15:04:05 +0000 UTC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Timestamp.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
