package core

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	*slog.Logger
	JSON *slog.Logger
}

type LogHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *LogHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.BlueString(level)
	case slog.LevelInfo:
		level = color.GreenString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006-01-02 15:04:05.000]")

	if len(fields) > 0 {
		h.l.Println(timeStr, level, r.Message, color.WhiteString(string(b)))
	} else {
		h.l.Println(timeStr, level, r.Message)
	}

	return nil
}

func NewLogHandler(w io.Writer, opts *slog.HandlerOptions) *LogHandler {
	return &LogHandler{
		Handler: slog.NewJSONHandler(w, opts),
		l:       log.New(w, "", 0),
	}
}

func NewLogger(cfg *Config) *Logger {
	var level slog.Level

	switch cfg.LogLevel {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelWarn
	}

	jsonLog := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	log := slog.New(NewLogHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return &Logger{
		Logger: log,
		JSON:   jsonLog,
	}
}
