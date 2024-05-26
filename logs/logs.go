package logs

import (
	"io"
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level 		string 		`yaml:"level"`
	//Formatter 	*logrus.Formatter 	`yaml:"formatter"`
	Output 		io.Writer			`yaml:"output"`
}

type DocLog struct {
	zapLog *zap.Logger
}

func(d *DocLog) Debug(message string, key string, values... any) {
	if len(values) == 0 {
		d.zapLog.Debug(message)
		return
	}

	for _, value := range values {
		d.zapLog.Debug(message, zap.Any(key, value))
	}

}

var Log *DocLog

func NewLog(config *Config) (logger *DocLog, err error) {
	if err = config.Validate(); err != nil {return}

	var log_lvl zap.AtomicLevel
	if log_lvl, err = zap.ParseAtomicLevel(config.Level); err != nil {return}

	var zap_cfg = zap.Config{
		Level				: zap.NewAtomicLevelAt(log_lvl.Level()),
		Development			: true,
		DisableCaller		: false,
		DisableStacktrace	: false,
		Sampling			: nil,
		Encoding			: "console",
		EncoderConfig		: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths			: []string{"../../logs/logs.txt", "stderr"},
		ErrorOutputPaths	: []string{"stderr"},
		InitialFields		: nil,
	}

	logger = &DocLog{}
	if logger.zapLog, err = zap_cfg.Build(); err != nil {return}
	return
}

func (c *Config) Validate() (err error) {
	if c == nil				{return errors.New("nil log config")}

	if c.Level == "" 		{return errors.New("nil level config")}
	//if c.Formatter == nil 	{return errors.New("nil formater config")}
	return
}