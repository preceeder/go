package rsa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestPublicEncryptLong(t *testing.T) {

	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RsaHandler{}
			err := r.SetPublicKey("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzN6tx98b4KZB1uqEuT7P\n/nWHrYqFdiy+Kzs9KZ6JtSQWb3b45loOsdUxFeaCAt+ZJ0+fNJRDnwc7AiKOlgbw\n0HT93WRVZXP6cwQV1Bg1XybBxtQE4OcEq+Uzzmd7RoBkQuNmjIUgDYtWPBSekSpZ\nAhWkk4dh8Nd7Qv2BvJNNOISVFcROFgMgbGz80v6WofR4nnTEdTB+j4pR/Q4dhnIR\nOlaWrai+hBPn95sahQ+Ujf7LZgLyhpyQeS+/xsLv29lDI6D+8neR1tsOYdOp8f8Q\nNwDkOroMlzxkQeYsJDLpLG8p58zHSdcLOsopVe2u41uzdrQ8qjhw4FU9eBOmFite\niwIDAQAB\n-----END PUBLIC KEY-----")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = r.SetPrivateKey("-----BEGIN PRIVATE KEY-----\nMIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDM3q3H3xvgpkHW\n6oS5Ps/+dYetioV2LL4rOz0pnom1JBZvdvjmWg6x1TEV5oIC35knT580lEOfBzsC\nIo6WBvDQdP3dZFVlc/pzBBXUGDVfJsHG1ATg5wSr5TPOZ3tGgGRC42aMhSANi1Y8\nFJ6RKlkCFaSTh2Hw13tC/YG8k004hJUVxE4WAyBsbPzS/pah9HiedMR1MH6PilH9\nDh2GchE6VpatqL6EE+f3mxqFD5SN/stmAvKGnJB5L7/Gwu/b2UMjoP7yd5HW2w5h\n06nx/xA3AOQ6ugyXPGRB5iwkMuksbynnzMdJ1ws6yilV7a7jW7N2tDyqOHDgVT14\nE6YWK16LAgMBAAECggEBAKFLC8yZdixHGPzohHgH4N94jsptjae9kDcfG4dB3y8y\n60r0gv9wlbMiotOYOHGkssKFaFWQCTESEz4aEOJDMqMcCKaeELGgPuUAqWLjcFmq\nfNNaJ0EeAMqI2GG/jQmzmbwjpqApS1P+iHUi0rh9e7gta/YOl2hzbgMO7W6XFivQ\npMIQZQE0WpmpK8cNgev/Xog8ZnHFC6XGUgK+mDVvJMYwmywUPIfLw2fvAZ29Qogt\nqiGeFCJSwAL1VkxryXSjJJBKuoc3cXEcq/hjhz6G9rvd50Lj2kCWMd8iqm/dtFyh\nDnT5WSFYNPIH0Up9qtqeP+TqgI/SrztAVHgUXVB2ABkCgYEA9cSeHG04Pj3p9ZCe\nCc6qb6L2kFphb62BhmUSHZ50p6X1KsSMw4wnzbrgrvcSe97iWZNLC536eQVHE5gL\n4ZjIxylYkp+FuuPHMIDseASR2pNmY2sJ83iTB4C9Y+37+64wBceFiXWBERdJA1t2\nMnzWLR8ijFfmHQ4KX3DJOhR05qUCgYEA1WYroahttPyvMFHdmcCphF9jhF3U6SGu\nVndwTtqaGLHzCmSvHxLFyd8ziw/F344IGIn8fIbOqhFAijyliD53kGMiKSUqMH4Q\neP2RfxGrZqek3f6pvyUtxfjXAh6+7pfL46u0AzmyvcpaGXqQToecCF43MCdbxh7Z\n3CViGfBcWW8CgYEAvJRcufU4ddHuJoYMLfxNHRIPXV5sa1PYEjaVevKuEkGuaF2e\noSF3HU4qvzZIEZJJXnA94jEbEydwjWFapIUmcmOQWhlbdLb4jYgvajwfanc11k04\nuoAnWVd4eygN9OWIZbbeCUaHfYS/ensAq+bMNJ0yVjvQDzVJ0kfpr84okR0CgYEA\nmBroNKTx9ZQ6Zu2jT2lVKuY27+1VygpY0ob1xS7psXp9asYTUMm3s0ll2tQWTV9W\ng+8uya/o9K2xXBcYQgGMhZ0zhzJXXRMuOJ88qt70VgpeaGGRqo4cj0TsNDWoEDag\nfJoxiC8DKWZnTEvhOihM3mYRXkBfmNr6nIEE6Mo7eP8CgYEA8KqzIk+5On3xmeES\nfQcLPYiaO9Hlttc7flyIpUL52Og7S1T/ekdiBVIDlePpjRx5H0iCtANyWmQq0Xbf\nSseQ9SFJ/4DLvDMawhvolmxHs98PNa8xZ9KdXgUNc7RcewUVhK2aLkxUQKNO0lww\nGGDGWfvePWzlVotJd0bM+a/X4qg=\n-----END PRIVATE KEY-----")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			data := []byte("adding再分段加密。除非我们是“定长定量自己可控可理解”的加密不需要padding")
			got, err := PriEncryptLong(r.prikey, data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//fmt.Println(got)
			fmt.Println(base64.StdEncoding.EncodeToString(got))

			iuy, err := PubKeyDecrypt(r.pubkey, got)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(iuy))
		})
	}
}
