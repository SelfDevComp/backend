package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/sklyar-vlad/selfDev/internal/config"
)

func NewLogger(cfgLogger config.ConfigLogger) (*zap.Logger, error) {
	var cfg zap.Config

	switch cfgLogger.Env {
	case "production":
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "development":
		cfg = zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		return nil, errors.New("failed create config logger")
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
