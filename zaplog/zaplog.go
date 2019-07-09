package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLog struct {
	Level zapcore.Level

	MaxSize int

	MaxAge int

	Outpath string

	Project string

	Module string
}

func (zl ZapLog) Build() {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "@timestamp",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	/********************* file core start *************************/
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename: zl.Outpath,
		MaxSize:  zl.MaxSize, // megabytes
		MaxAge:   zl.MaxAge,  // days
	})
	fileCore := zapcore.NewCore(encoder, ws, zl.Level)
	/********************* file core end *************************/

	/********************* stdout core start *************************/
	sink, closeOut, err := zap.Open("stdout")
	if err != nil {
		closeOut()
	}
	stdoutCore := zapcore.NewCore(encoder, sink, zl.Level)
	/********************* stdout core start *************************/

	core := zapcore.NewTee(fileCore, stdoutCore)
	logger := zap.New(core, zap.Fields(zap.String("project", zl.Project), zap.String("module", zl.Module)))
	zap.ReplaceGlobals(logger)
}
