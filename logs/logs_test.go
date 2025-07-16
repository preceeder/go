package logs

import (
	"encoding/json"
	"fmt"
	"github.com/preceeder/go/base"
	"log/slog"
	"strconv"
	"strings"
	"testing"
)

type Dar string

func (d Dar) MarshalJSON() ([]byte, error) {
	return []byte(d), nil
}

func BenchmarkTestNewSlog(b *testing.B) {
	cfg := SlogConfig{
		ErrorFileName:           "logs/error.log",
		InfoFileName:            "logs/out.log",
		TransparentTransmission: true,
		StdOut:                  "0",
		MaxSize:                 1,
		MaxAge:                  3,
		MaxBackups:              4,
	}
	NewSlog(cfg)
	//i := 1000
	//j := 0
	ctx := base.Context{RequestId: "sdsd"}
	slog.InfoContext(ctx, "sds")

	marshal, _ := json.Marshal(map[string]any{"Name": "ois", "Age": 23})
	de := Dar(marshal)
	slog.InfoContext(ctx, "", "data", de)
	//for j <= b.N {
	//	j++
	//	slog.InfoContext(ctx, "sds")
	//	slog.ErrorContext(ctx, "hahah", "error", "ssss")
	//	//slog.Info("", "sss", j)
	//}

}

func BenchmarkNewSlog(b *testing.B) {
	file := "scabhvuiehviuusdvuauisgvuiav"
	line := 3
	//j := 0
	//for j <= b.N {
	//	j++
	//	//fmt.Sprintf("%s:%d", file, line)
	//	//
	//	_ = strings.Join([]string{file, strconv.Itoa(line)}, ":")
	//
	//}

	dd := strings.Join([]string{file, strconv.Itoa(line)}, ":")
	fmt.Println(dd)
}
