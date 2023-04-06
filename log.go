package hass_mqtt

import (
	"fmt"
	"github.com/go-logr/logr"
)

type logError struct {
	logger logr.Logger
	level  int
}

func newLogError(logger logr.Logger, level int) logError {
	return logError{
		logger: logger,
		level:  level,
	}
}

func (l logError) Println(v ...interface{}) {
	l.logger.V(l.level).Error(fmt.Errorf("%v", v), fmt.Sprintf("%v", v))
}

func (l logError) Printf(format string, v ...interface{}) {
	l.logger.V(l.level).Error(fmt.Errorf(format, v...), fmt.Sprintf(format, v...))
}

type logInfo struct {
	logger logr.Logger
	level  int
}

func newLogInfo(logger logr.Logger, level int) logInfo {
	return logInfo{
		logger: logger,
		level:  level,
	}
}

func (l logInfo) Println(v ...interface{}) {
	l.logger.V(l.level).Info(fmt.Sprintf("%v", v))
}

func (l logInfo) Printf(format string, v ...interface{}) {
	l.logger.V(l.level).Info(fmt.Sprintf(format, v...))
}
