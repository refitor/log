# rslog
Minimalist log unified interface library. Supports quick access to third-party log components. The default support system and zap log components

极简日志统一接口库. 支持快速接入第三方日志组件. 默认支持系统及zap日志组件

# Usage
```go
package main

import (
    "github.com/refitor/rslog"
)

func main() {
    // rslog.UseLog(rslog.C_Log_Zap)
    
    //// reset zap logger
    // fileWriter := rslog.GetWriter(logFile, 30)
    // core := zapcore.NewCore(
    //     zapcore.NewJSONEncoder(rslog.ZapNewEncoderConfig()),
    //     // zapcore.NewConsoleEncoder(rslog.ZapNewEncoderConfig()),
    //     zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(fileWriter)),
    //     zap.NewAtomicLevelAt(zapcore.InfoLevel),
    // )
    // rslog.ResetZapLog(zap.New(core, zap.AddCaller()))

    rslog.Info("test Info")
    rslog.Infof("test Infof: %s", "hello log")

    rslog.Debug("test Debug")
    rslog.Debugf("test Debugf: %s", "hello log")

    rslog.Warning("test Warning")
    rslog.Warningf("test Warningf: %s", "hello log")

    rslog.Error("test Error")
    rslog.Errorf("test Errorf: %s", "hello log")
}
```
