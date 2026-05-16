package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	logLevelEnv    = "LOG_LEVEL"
	serviceNameEnv = "SERVICE_NAME"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	log.SetLevel(getLogLevel())
}

func serviceName() string {
	if name := os.Getenv(serviceNameEnv); name != "" {
		return name
	}
	return "unknown"
}

func NewLogger() *logrus.Entry {
	return log.WithField("service_name", serviceName())
}

func getLogLevel() logrus.Level {
	levelStr := os.Getenv(logLevelEnv)
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		level = logrus.InfoLevel // Default to info level
	}
	return level
}

func WithRequestID(requestID string) *logrus.Entry {
	return NewLogger().WithField("request_id", requestID)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return NewLogger().WithFields(fields)
}
