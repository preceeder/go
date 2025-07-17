package jigou

import "testing"

func Test_generateSignature(t *testing.T) {
	type args struct {
		appId          string
		serverSecret   string
		signatureNonce string
		timeStamp      int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test_generateSignature",
			args: args{appId: "962422020", serverSecret: "e321843ea81546836e48e889af5e051b", signatureNonce: "f90d10b0f67e1831", timeStamp: 1747131817},
			want: "7461022d28f797be6a13605e71fa430b",
		},
	}
	//962422020 f90d10b0f67e1831 1747131817
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSignature(tt.args.appId, tt.args.serverSecret, tt.args.signatureNonce, tt.args.timeStamp); got != tt.want {
				t.Errorf("generateSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
