package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

func New(out io.Writer, verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.Out = out
	logger.Formatter = &logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	}
	logger.Level = logrus.InfoLevel

	if verbose {
		logger.Level = logrus.TraceLevel
	}

	return logger
}
