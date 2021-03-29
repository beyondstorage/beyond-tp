package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapFactory init an instance of zap.Logger with given level
func zapFactory() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	logLevel := viper.GetString("log_level")

	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(logLevel))
	if err != nil {
		err = nil
		println(fmt.Sprintf("%s is not a valid loglevel, use info instead", logLevel))
		*l = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, os.Stderr, l)

	// enable caller and error stacktrace for dev level
	if l.Enabled(zap.DebugLevel) {
		logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		return logger
	}

	return zap.New(core)
}
