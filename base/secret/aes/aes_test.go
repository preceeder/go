package aes

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestDecryptCBC(t *testing.T) {
	type args struct {
		data string
		key  []byte
		iv   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: args{data: "DgkDAJ4ZoLJm75qkDSOj9w==", key: []byte("juytrgtegdtsydgw"), iv: []byte("78uytdhsgdtegdys")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dd, _ := base64.StdEncoding.DecodeString(tt.args.data)
			got, err := DecryptCBC(dd, tt.args.key, tt.args.iv)
			fmt.Println(string(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptCBC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptCBC() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptCBC(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
		iv   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: args{data: []byte("hello"), key: []byte("juytrgtegdtsydgw"), iv: []byte("78uytdhsgdtegdys")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptCBC(tt.args.data, tt.args.key, tt.args.iv)
			encodedCiphertext := base64.StdEncoding.EncodeToString(got)
			fmt.Println("Encrypted:", encodedCiphertext)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptCBC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptCBC() got = %v, want %v", got, tt.want)
			}
		})
	}
}
