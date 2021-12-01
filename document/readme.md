
docUrl:https://pkg.go.dev/go.uber.org/zap


go mod init github.com/marszhanglin/LogZ



//package main
//?
//import (
//	"encoding/json"
//	"os"
//	"time"
//)
//
//func main() {
//	//logger,_:=logZInit("logs");
//	//sugarLog(logger)
//	//loggerLog(logger)
//
//	mLogZ := logz.InitCore(time.Now().Format("2006_0102_1504_05"))
//	mLogZ.SugarInfoF("Failed to fetch URL: %s", "demo测试")
//
//}
//
//func logZInit(logPath string) (*zap.Logger, error) {
//	cmdPath, _ := os.Getwd()
//	rawJSON := []byte(`{
//	  "level": "debug",
//	  "encoding": "json",
//	  "outputPaths": ["stdout", "` + cmdPath + "/" + logPath + `"],
//	  "errorOutputPaths": ["stderr"],
//	  "initialFields": {"foo": "bar"},
//	  "encoderConfig": {
//	    "messageKey": "message",
//	    "levelKey": "level",
//	    "levelEncoder": "lowercase"
//	  }
//	}`)
//
//	var cfg zap.Config
//	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
//		panic(err)
//	}
//	logger, err := cfg.Build()
//	if err != nil {
//		panic(err)
//	}
//	defer logger.Sync()
//	logger.Info("logger construction succeeded")
//	return logger, err
//
//}
//
///**
//更快
//*/
//func loggerLog(logger *zap.Logger) {
//	defer logger.Sync()
//	//cmdPath,_:=logZ.GetCurrentPath();
//	dir, _ := os.Getwd()
//	logger.Info("failed to fetch URL",
//		// Structured context as strongly typed Field values.
//		zap.String("url", dir),
//		zap.Int("attempt", 3),
//		zap.Duration("backoff", time.Second),
//	)
//
//}
//
///**
//use the SugaredLogger. It's 4-10x faster than other structured logging packages
//*/
//func sugarLog(logger *zap.Logger) {
//	defer logger.Sync()
//	sugar := logger.Sugar()
//	sugar.Infow("failed to fetch URL",
//		"attempt", 3,
//		"backoff", time.Second,
//	)
//	sugar.Infof("Failed to fetch URL: %s", "什么鬼")
//}


