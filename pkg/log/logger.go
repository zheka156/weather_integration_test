package log

import (
	"fmt"

	"go.uber.org/zap"
)

type TestLogger struct {
	Zap *zap.Logger
}

func NewLogger() *TestLogger {
	log, err := zap.NewDevelopment()
	if err != nil {
		log.Sugar().Fatalf("can't start zap logger: %e", err)
	}
	defer log.Sync()
	return &TestLogger{Zap: log}
}

func (log *TestLogger) InfofJSON(logtext string, a ...interface{}) {
	log.Zap.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(logtext+": %+v", a...))
}

func (log *TestLogger) Debugf(format string, a ...interface{}) {
	log.Zap.Sugar().Debugf(format, a...)
}
