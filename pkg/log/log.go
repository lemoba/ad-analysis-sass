package log

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const ctxLoggerKey = "zapLogger"

type Logger struct {
	*zap.Logger
}

func NewLog(conf *viper.Viper) *Logger {
	lf := conf.GetString("log.log_file_name")
	lv := conf.GetString("log.log_level")

	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	hook := lumberjack.Logger{
		Filename:   lf,
		MaxSize:    conf.GetInt("log.log_max_size"),
		MaxBackups: conf.GetInt("log.log_max_backups"),
		MaxAge:     conf.GetInt("log.log_max_age"),
		Compress:   conf.GetBool("log.log_compress"),
	}

	var encoder zapcore.Encoder

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "Logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseColorLevelEncoder,
	}

	if conf.GetString("log.encoding") == "console" {
		encoderConfig.EncodeTime = timeEncoder
		encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig.FunctionKey = zapcore.OmitKey
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
		encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // Print to console and file
		level,
	)

	if conf.GetString("env") != "prod" {
		return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}

	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}
