package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	serviceName = "users"
	logLevelEnv = "LOG_LEVEL"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	level := getLogLevel()
	log.SetLevel(level)
}

func NewLogger() *logrus.Logger {
	return log
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
	return log.WithField("request_id", requestID)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return log.WithFields(fields)
}
