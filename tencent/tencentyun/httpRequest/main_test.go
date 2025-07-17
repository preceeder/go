package httpRequest

import (
	"fmt"
	"testing"
	"time"
)

func TestNewIPv4FallbackClient(t *testing.T) {
	client := NewIPv4FallbackClient(30 * time.Second)
	resp, _ := client.R().Get("https://httpbin.org/ip")
	fmt.Println(resp.String())
}
