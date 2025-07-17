package tencentIm

import (
	"github.com/go-resty/resty/v2"
	"github.com/preceeder/go/base"
	"reflect"
	"testing"
)

func TestNewTencentIm(t *testing.T) {
	type args struct {
		config TencentImConfig
	}
	tests := []struct {
		name string
		args args
		want TencentImClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTencentIm(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTencentIm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTencentImClient_GetUserSign(t1 *testing.T) {
	type fields struct {
		Config   TencentImConfig
		Client   *resty.Client
		UserSign string
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantUserSign string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TencentImClient{
				Config:   tt.fields.Config,
				Client:   tt.fields.Client,
				UserSign: tt.fields.UserSign,
			}
			gotUserSign, err := t.GetUserSign(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetUserSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserSign != tt.wantUserSign {
				t1.Errorf("GetUserSign() gotUserSign = %v, want %v", gotUserSign, tt.wantUserSign)
			}
		})
	}
}

func TestTencentImClient_SendImRequest(t1 *testing.T) {
	type fields struct {
		Config   TencentImConfig
		Client   *resty.Client
		UserSign string
	}
	type args struct {
		ctx         base.Context
		serverName  string
		requestData any
		respBody    any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TencentImClient{
				Config:   tt.fields.Config,
				Client:   tt.fields.Client,
				UserSign: tt.fields.UserSign,
			}
			if err := t.SendImRequest(tt.args.ctx, tt.args.serverName, tt.args.requestData, tt.args.respBody); (err != nil) != tt.wantErr {
				t1.Errorf("SendImRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTencentImClient_getEcdsaSign(t1 *testing.T) {
	type fields struct {
		Config   TencentImConfig
		Client   *resty.Client
		UserSign string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TencentImClient{
				Config:   tt.fields.Config,
				Client:   tt.fields.Client,
				UserSign: tt.fields.UserSign,
			}
			if got := t.getEcdsaSign(); got != tt.want {
				t1.Errorf("getEcdsaSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTencentImClient_getHmacSign(t1 *testing.T) {
	type fields struct {
		Config   TencentImConfig
		Client   *resty.Client
		UserSign string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TencentImClient{
				Config:   tt.fields.Config,
				Client:   tt.fields.Client,
				UserSign: tt.fields.UserSign,
			}
			if got := t.getHmacSign(); got != tt.want {
				t1.Errorf("getHmacSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTencentImClient_setUrl(t1 *testing.T) {
	type fields struct {
		Config   TencentImConfig
		Client   *resty.Client
		UserSign string
	}
	type args struct {
		serverName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TencentImClient{
				Config:   tt.fields.Config,
				Client:   tt.fields.Client,
				UserSign: tt.fields.UserSign,
			}
			if got := t.setUrl(tt.args.serverName); got != tt.want {
				t1.Errorf("setUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
