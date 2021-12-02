package logz

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)


type logz struct {
	mLogger *zap.Logger
	mSugar  *zap.SugaredLogger
}

func InitCore(logPath string) *logz {

	// 日志级别
	logLevel := "DEBUG"

	atomicLevel := zap.NewAtomicLevel()
	switch logLevel {
	case "DEBUG":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "DPANIC":
		atomicLevel.SetLevel(zapcore.DPanicLevel)
	case "PANIC":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case "FATAL":
		atomicLevel.SetLevel(zapcore.FatalLevel)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	cmdPath, _ := os.Getwd()

	fileName := cmdPath + "/logs/" + logPath

	// 日志轮转
	writer := &lumberjack.Logger{
		// 日志名称
		Filename: fileName,
		// 日志大小限制，单位MB
		MaxSize: 100,
		// 历史日志文件保留天数
		MaxAge: 30,
		// 最大保留历史日志数量
		MaxBackups: 10,
		// 本地时区
		LocalTime: true,
		// 历史日志文件压缩标识
		Compress: false,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		atomicLevel,
	)
	fLogger := zap.New(zapCore, zap.AddCaller())

	fSugar := fLogger.Sugar()
	defer fLogger.Sync()
	fLogger.Info("logger construction succeeded")
	fSugar.Info("mSugar logger construction succeeded")

	logZ := &logz{
		mLogger: fLogger,
		mSugar:  fSugar,
	}
	return logZ

}

func Init(logPath string) *logz {
	cmdPath, _ := os.Getwd()
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "` + cmdPath + "/" + logPath + `"],
	  "errorOutputPaths": ["stderr"],
	  
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	// "initialFields": {"key1": "value1"},

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	var err error
	fLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	fSugar := fLogger.Sugar()
	defer fLogger.Sync()
	fLogger.Info("logger construction succeeded")
	fSugar.Info("mSugar logger construction succeeded")

	logZ := &logz{
		mLogger: fLogger,
		mSugar:  fSugar,
	}
	return logZ

}

//"Failed to fetch URL: %s", url
func (log logz) SugarInfoF(template string, args ...interface{}) {
	log.mSugar.Infof(template, args)
}

func (log logz) SugarInfo(str string) {
	log.mSugar.Info(str)
}
