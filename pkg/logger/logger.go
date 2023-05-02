package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func LogMyApp(typ, template, process string, err error) {
	logger := logrus.New()

	file, _ := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(file)

	StrBuild := logrus.Fields{}

	switch typ {
	case "i":
		StrBuild = logrus.Fields{
			"Process": process,
		}

		logger.WithFields(StrBuild).Info(template)
	case "e":
		StrBuild = logrus.Fields{
			"Process":   process,
			"error msg": err,
		}
		logger.WithFields(StrBuild).Error(template)
	}
}
