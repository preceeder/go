package common

import (
	"log/slog"
)

type Logger struct {
	Logger *slog.Logger
}

func (l Logger) Output(depth int, s string) error {
	switch depth {
	case 0:
		l.Logger.Debug(s)
	case 1:
		l.Logger.Info(s)
	case 2:
		l.Logger.Warn(s)
	case 3:
		l.Logger.Error(s)
	}
	return nil
}
