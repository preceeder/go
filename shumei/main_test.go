package shumei

import (
	"reflect"
	"testing"
)

func TestNewShumeiClient(t *testing.T) {
	type args struct {
		config ShumeiConfig
	}
	tests := []struct {
		name string
		args args
		want *ShuMei
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShumeiClient(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShumeiClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
