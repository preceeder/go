package rsa

import (
	"fmt"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	tests := []struct {
		name       string
		wantPrvkey []byte
		wantPubkey []byte
	}{
		// TODO: Add test cases.
		{name: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrvkey, gotPubkey := GenRsaKey()
			fmt.Println(string(gotPrvkey))
			fmt.Println(string(gotPubkey))
		})
	}
}
