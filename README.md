# log
Minimalist log unified interface library. Supports quick access to third-party log components. The default support system and zap log components

极简日志统一接口库. 支持快速接入第三方日志组件. 默认支持系统及zap日志组件

# Usage
```go
package main

import (
	"github.com/refitor/log"
)

func main() {
	// // zap log
    // log.UseLog(zap.New())
	
    // // zap log with file
	// fileWriter := zap.GetZapWriter("demo.log", 30)
	// core := zapcore.NewCore(
	// 	zapcore.NewJSONEncoder(zap.ZapNewEncoderConfig()),
	// 	// zapcore.NewConsoleEncoder(log.ZapNewEncoderConfig()),
	// 	zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(fileWriter)),
	// 	zaplog.NewAtomicLevelAt(zapcore.InfoLevel),
	// )
	// log.ResetLog(zaplog.New(core, zaplog.AddCaller()))

    log.Info("test Info")
    log.Infof("test Infof: %s", "hello log")

    log.Debug("test Debug")
    log.Debugf("test Debugf: %s", "hello log")

    log.Warn("test Warning")
    log.Warnf("test Warningf: %s", "hello log")

    log.Error("test Error")
    log.Errorf("test Errorf: %s", "hello log")
}
```
