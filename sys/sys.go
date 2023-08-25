package sys

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type sysLog struct {
	level string

	depth int

	logger *log.Logger
}

func New() *sysLog {
	return &sysLog{
		level:  "info",
		depth:  4,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (p *sysLog) SetDepth(depth int) {
	p.depth = depth
}

func (p *sysLog) print(level string, fprintf func(string, ...interface{}), datas ...interface{}) {
	if getLogLevel(level) > getLogLevel(p.level) {
		return
	}
	lrunes := []rune(level)
	slevel := strings.ToUpper(string(lrunes[:1]) + string(lrunes[1:]))
	format := fmt.Sprintf("%s %s ", slevel, stackTrace(4)) + "%v"
	fprintf(format, datas...)
}

func (p *sysLog) printf(level string, fprintf func(string, ...interface{}), format string, datas ...interface{}) {
	if getLogLevel(level) > getLogLevel(p.level) {
		return
	}
	lrunes := []rune(level)
	slevel := strings.ToUpper(string(lrunes[:1]) + string(lrunes[1:]))
	format = fmt.Sprintf("%s %s ", slevel, stackTrace(4)) + format
	fprintf(format, datas...)
}

func (p *sysLog) Info(datas ...interface{}) {
	p.print("info", p.logger.Printf, datas...)
}

func (p *sysLog) Infof(format string, datas ...interface{}) {
	p.printf("info", p.logger.Printf, format, datas...)
}

func (p *sysLog) Debug(datas ...interface{}) {
	p.print("debug", p.logger.Printf, datas...)
}

func (p *sysLog) Debugf(format string, datas ...interface{}) {
	p.printf("debug", p.logger.Printf, format, datas...)
}

func (p *sysLog) Warn(datas ...interface{}) {
	p.print("warn", p.logger.Printf, datas...)
}

func (p *sysLog) Warnf(format string, datas ...interface{}) {
	p.printf("warn", p.logger.Printf, format, datas...)
}

func (p *sysLog) Error(datas ...interface{}) {
	p.print("error", p.logger.Printf, datas...)
}

func (p *sysLog) Errorf(format string, datas ...interface{}) {
	p.printf("error", p.logger.Printf, format, datas...)
}

func (p *sysLog) SetLevel(l string) {
	p.level = l
}

func (p *sysLog) ResetLog(l interface{}) {
	if logger, ok := l.(*log.Logger); ok {
		p.logger = logger
	}
}
