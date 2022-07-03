package providers

import (
	"os/exec"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLoggerProvider(config *viper.Viper) *zap.SugaredLogger {
	env := config.GetString("app.Env")
	var log *zap.Logger
	if env != "production" {
		logConfig := zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		log, _ = logConfig.Build()
		log = log.WithOptions(
			zap.Hooks(func(entry zapcore.Entry) error {
				if entry.Level == zapcore.PanicLevel {
					cmd := exec.Command("/usr/bin/say", entry.Message)
					// 执行命令，返回命令是否执行成功
					_ = cmd.Run()
				}
				return nil
			}),
		)
	} else {
		log, _ = zap.NewProduction()
	}
	return log.Sugar()
}
