package conf

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriterSyncer(), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	stSeparator := string(os.PathSeparator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"

	luberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   true,
	}

	return zapcore.AddSync(luberjackSyncer)
}
