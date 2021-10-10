package log

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func Get() *logrus.Logger {
	return log
}

func InitializeLog() {
	log = logrus.New()
}
