package config

import (
	"testing"
)

func Test_getAppPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{path: ""},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAppPath(tt.args.path); got[0:1] != tt.want {
				t.Errorf("getAppPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
