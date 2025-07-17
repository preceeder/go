package producer

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestNewProduct(t *testing.T) {

	_, err := NewProduct("127.0.0.1:4160", "producter", "test")
	if err != nil {
		fmt.Println(err.Error())
	}
	se := map[string]any{
		"name": "王二 10 jisud",
		"time": time.Now().Format(time.DateTime),
	}
	marshal, err := json.Marshal(se)
	if err != nil {
		fmt.Println(err.Error())
	}
	//err = NsqProducer["producter"].Publish("shumei", marshal)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	err = NsqProducer["producter"].DeferredPublish("shumei", time.Second*100, marshal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
