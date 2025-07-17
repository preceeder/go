package tencentyun

import (
	"reflect"
	"testing"
)

func TestNewTencentIdentityClient(t *testing.T) {
	type args struct {
		config TencentIdentityConfig
	}
	tests := []struct {
		name string
		args args
		want TencentIdentityClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTencentIdentityClient(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTencentIdentityClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
