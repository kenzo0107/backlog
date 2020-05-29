package backlog

import "testing"

func TestJSONTime_String(t *testing.T) {
	tests := []struct {
		name string
		t    JSONTime
		want string
	}{
		{
			name: "set JSONTime( date in the response of backlog API )",
			t:    JSONTime("2020-02-19T05:54:32Z"),
			want: "\"2020-02-19T05:54:32Z\"",
		},
		{
			name: "empty JSONTime",
			t:    JSONTime(""),
			want: "\"0001-01-01T00:00:00Z\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("JSONTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
