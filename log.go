package hass_mqtt

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
)

type mqttLog struct {
	logger *slog.Logger
	level  slog.Level
}

func newMQTTLog(logger *slog.Logger, level slog.Level) mqtt.Logger {
	return mqttLog{
		logger: logger,
		level:  level,
	}
}

func (l mqttLog) Println(v ...interface{}) {
	l.logger.Log(context.Background(), l.level, fmt.Sprintf("(hass-mqtt) %v", v))
}

func (l mqttLog) Printf(format string, v ...interface{}) {
	l.logger.Log(context.Background(), l.level, fmt.Sprintf("(hass-mqtt) "+format, v...))
}
