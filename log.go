package rslog

import (
	"github.com/refitor/rslog/sys"
)

var v_logger i_log = sys.New()

type i_log interface {
	Info(datas ...interface{})
	Infof(format string, datas ...interface{})

	Debug(datas ...interface{})
	Debugf(format string, datas ...interface{})

	Warn(datas ...interface{})
	Warnf(format string, datas ...interface{})

	Error(datas ...interface{})
	Errorf(format string, datas ...interface{})

	SetLevel(l string)
	SetDepth(depth int)
	ResetLog(l interface{})
}

func Info(datas ...interface{}) {
	v_logger.Info(datas...)
}

func Infof(format string, datas ...interface{}) {
	v_logger.Infof(format, datas...)
}

func Debug(datas ...interface{}) {
	v_logger.Debug(datas...)
}

func Debugf(format string, datas ...interface{}) {
	v_logger.Debugf(format, datas...)
}

func Warn(datas ...interface{}) {
	v_logger.Warn(datas...)
}

func Warnf(format string, datas ...interface{}) {
	v_logger.Warnf(format, datas...)
}

func Error(datas ...interface{}) {
	v_logger.Error(datas...)
}

func Errorf(format string, datas ...interface{}) {
	v_logger.Errorf(format, datas...)
}

func SetDepth(depth int) {
	v_logger.SetDepth(depth)
}

func SetLevel(l string) {
	v_logger.SetLevel(l)
}

func ResetLog(l interface{}) {
	v_logger.ResetLog(l)
}

func UseLog(l i_log) {
	if l == nil {
		l = sys.New()
	}
	v_logger = l
}
