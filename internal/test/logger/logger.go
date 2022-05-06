package logger

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func NoopLogger() logrus.FieldLogger {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	return logger
}
