package logger

import (
	"github.com/geekymedic/neonx"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Entry
}

func (m *Logger) LogInfo(format string, v ...interface{}) {
	m.logger.Infof(format, v...)
}

func (m *Logger) LogError(format string, v ...interface{}) {
	m.logger.Errorf(format, v...)
}

func NewLogger(session *neonx.Session) *Logger {

	if session != nil {
		x := logrus.WithFields(
			logrus.Fields{
				"_uid":      session.Uid,
				"_token":    session.Token,
				"_platform": session.Platform,
				"_version":  session.Version,
				"_net":      session.Net,
				"_mobile":   session.Mobile,
				"_os":       session.OS,
				"_device":   session.Device,
				"_describe": session.Describe,
				"_trace":    session.Trace,
				"_sequence": session.Sequence,
				"_time":     session.Time,
			})

		return &Logger{x}
	}

	return &Logger{
		logger: logrus.WithFields(
			logrus.Fields{},
		),
	}
}
