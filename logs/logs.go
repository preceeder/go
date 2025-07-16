package logs

import (
	"github.com/preceeder/logs/lumberjack"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type SlogConfig struct {
	DebugFileName           string `json:"debugFileName"`           // 配置这个 debug 等级的日志 就会写入这个文件
	InfoFileName            string `json:"infoFileName"`            // 配置这个 info 等级的日志 就会写入这个文件
	WarnFileName            string `json:"warnFileName"`            // 配置这个 warn 等级的日志 就会写入这个文件
	ErrorFileName           string `json:"errorFileName"`           // 配置这个 error 等级的日志 就会写入这个文件
	MaxSize                 int    `json:"maxSize"`                 // 日志轮转的最大size
	MaxAge                  int    `json:"maxAge"`                  // 历史日志保存的最大时间 天
	MaxBackups              int    `json:"maxBackups"`              // 保存的 历史日志数量
	TransparentTransmission bool   `json:"transparentTransmission"` // 日志输出是否传递到高等级等级  level 越小等级越高, ErrorFileName 日志可以写入到 WarnFileName|InfoFileName|DebugFileName
	StdOut                  string `json:"stdOut"`                  // 是否在终端输出   0 不输出, 1 输出
	Compress                bool   `json:"compress"`                // 备份的日志是否压缩
	NotSetDefaultLog        bool   `json:"notSetDefaultLog"`        // 不设置为默认的slog   true 不设置， false设置
}

// LogWriters 目前没啥用
var LogWriters []*lumberjack.Logger = make([]*lumberjack.Logger, 0)

func NewSlog(cfg SlogConfig) *slog.Logger {
	return mHandlerSlog(cfg)
}

func mHandlerSlog(cfg SlogConfig) *slog.Logger {
	MinLevel, iohandler := getIoWriter(cfg)
	if cfg.StdOut == "1" {
		for level, iod := range iohandler {
			iohandler[level] = io.MultiWriter(iod, os.Stdout)
		}
	}
	opt := &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			d := a.Value.Any().(*slog.Source)
			//d.Function = ""
			//d.File = fmt.Sprintf("%s:%d", d.File, d.Line)
			//strings.Join(d.File, d.Line)
			a.Value = slog.AnyValue(strings.Join([]string{d.File, strconv.Itoa(d.Line)}, ":"))
		} else if a.Key == slog.TimeKey {
			a.Value = slog.StringValue(a.Value.Time().Format("20060102150405.000000"))
		}
		return a
	}}

	var log *slog.Logger
	log = slog.New(NewMoreHandler(iohandler, MinLevel, cfg.TransparentTransmission, opt))
	if !cfg.NotSetDefaultLog {
		slog.SetDefault(log)
	}
	return log
}

func getIoWriter(cfg SlogConfig) (slog.Level, map[slog.Level]io.Writer) {
	MinLevel := slog.LevelDebug

	iohandler := map[slog.Level]io.Writer{}

	if cfg.ErrorFileName != "" && len(cfg.ErrorFileName) > 0 {
		el := &lumberjack.Logger{
			Filename:   cfg.ErrorFileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}
		LogWriters = append(LogWriters, el)
		iohandler[slog.LevelError] = el
		MinLevel = slog.LevelError
	}

	if cfg.WarnFileName != "" && len(cfg.WarnFileName) > 0 {
		wl := &lumberjack.Logger{
			Filename:   cfg.WarnFileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}
		LogWriters = append(LogWriters, wl)
		iohandler[slog.LevelWarn] = wl
		MinLevel = slog.LevelWarn
	}
	if cfg.InfoFileName != "" && len(cfg.InfoFileName) > 0 {
		ll := &lumberjack.Logger{
			Filename:   cfg.InfoFileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}
		LogWriters = append(LogWriters, ll)
		iohandler[slog.LevelInfo] = ll
		MinLevel = slog.LevelInfo
	}

	if cfg.DebugFileName != "" && len(cfg.DebugFileName) > 0 {
		dl := &lumberjack.Logger{
			Filename:   cfg.DebugFileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}
		LogWriters = append(LogWriters, dl)
		iohandler[slog.LevelDebug] = dl
		MinLevel = slog.LevelDebug
	}
	return MinLevel, iohandler
}

// json 字符串可以不加转议符"\" 输出
type LogStr string

func (d LogStr) MarshalJSON() ([]byte, error) {
	return []byte(d), nil
}
