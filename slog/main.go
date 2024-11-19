package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/samber/lo"
)

func init() {
	SetUpLogger(LoggerOpts{
		Level: slog.LevelError,
		Out:   os.Stdout,
	})
}

func main() {
	// no output
	log.Println("hello world")
	// output
	slog.ErrorContext(context.Background(), "hello world")
}

type LoggerOpts struct {
	Level slog.Level
	Out   io.Writer
}

func SetUpLogger(opts LoggerOpts) {
	out := lo.Ternary[io.Writer](
		opts.Out != nil,
		opts.Out,
		os.Stdout,
	)
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		// ログに出力される最小のレベル
		Level:     opts.Level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key != slog.LevelKey {
				return a
			}
			return a
		},
	})
	l := slog.New(jsonHandler)
	slog.SetDefault(l)
}
