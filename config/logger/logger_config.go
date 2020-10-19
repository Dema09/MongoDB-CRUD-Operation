package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var(
	log *zap.Logger
)

func init(){
	loggingConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "timestamp",
			MessageKey: "message",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = loggingConfig.Build(); err != nil{
		panic(err)
	}
}

func Info(message string, tags ...zap.Field){
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field){
	if err != nil{
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(message, tags...)
	log.Sync()
}
