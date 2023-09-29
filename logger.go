package gowin

import "log"

var logger Logger = &defaultLogger{}

type Logger interface {
	Infoln(args ...any)
	Errorln(args ...any)
}

// SetLogger if original log output needed
func SetLogger(l Logger) {
	logger = l
}

type defaultLogger struct {
}

func (m *defaultLogger) Errorln(args ...any) {
	log.Println(args...)
}

func (m *defaultLogger) Infoln(args ...any) {
	log.Println(args...)
}
