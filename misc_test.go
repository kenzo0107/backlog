package backlog

import "testing"

func TestProjectIDOrKey(t *testing.T) {
	type args struct {
		projectIDOrKey interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "set projectIDOrKey int",
			args: args{
				projectIDOrKey: 12345,
			},
			want:    "12345",
			wantErr: false,
		},
		{
			name: "set projectIDOrKey string",
			args: args{
				projectIDOrKey: "AIUEO",
			},
			want:    "AIUEO",
			wantErr: false,
		},
		{
			name: "set projectIDOrKey bool",
			args: args{
				projectIDOrKey: true,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "set projectIDOrKey []string",
			args: args{
				projectIDOrKey: []string{"a"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := projIDOrKey(tt.args.projectIDOrKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectIDOrKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("projectIDOrKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
