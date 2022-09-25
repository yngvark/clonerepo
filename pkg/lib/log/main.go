package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

func New(out io.Writer) *logrus.Logger {
	logger := logrus.New()
	logger.Out = out
	logger.Formatter = &logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	}
	logger.Level = logrus.InfoLevel

	return logger
}
