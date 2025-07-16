package base

import "testing"

func TestVersionCompare(t *testing.T) {
	type args struct {
		v1 string
		v2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{
			v1: "",
			v2: "2.3.4.1",
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionCompare(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("VersionCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
