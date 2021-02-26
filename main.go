package main

import (
	"context"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/aos-dev/dm/cmd"
	ilog "github.com/aos-dev/dm/pkg/logger"
)

const (
	logLvlInEnv   = "log_level"
	defaultLogLvl = "warn"
)

func main() {
	if err := setEnvConfig(); err != nil {
		log.Printf("set env config failed: %s", err)
		os.Exit(1)
	}

	logger, err := initLogger(viper.GetString(logLvlInEnv))
	if err != nil {
		log.Printf("init logger failed: %s", err)
		os.Exit(1)
	}

	if err := cmd.Init(); err != nil {
		logger.Error("init command line failed:", zap.Error(err))
		os.Exit(1)
	}

	ctx := ilog.ContextWithLogger(context.Background(), logger)
	if err := cmd.RootCmd.ExecuteContext(ctx); err != nil {
		logger.Error("execute command failed:", zap.Error(err))
		os.Exit(1)
	}
}

// setEnvConfig bind config from env, and set with viper
func setEnvConfig() error {
	viper.SetEnvPrefix(cmd.Name)
	viper.AutomaticEnv()
	// ignore this error because it's always nil
	_ = viper.BindEnv(logLvlInEnv)
	viper.SetDefault(logLvlInEnv, defaultLogLvl)
	return nil
}

// initLogger init an instance of zap.Logger with given level
func initLogger(lvl string) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(lvl))
	if err != nil {
		return nil, err
	}
	core := zapcore.NewCore(encoder, os.Stderr, l)

	// enable caller and error stacktrace for debug level
	if l.Enabled(zap.DebugLevel) {
		logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		return logger, nil
	}

	return zap.New(core), nil
}
