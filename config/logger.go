package config

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

var Log *logrus.Logger

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}


func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func InitLogger(){
	Log := logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(ioutil.Discard)

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.AddHook(&WriterHook{
			Writer: file,
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		})
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

	Log.AddHook(&WriterHook{
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
}

func Logger() *logrus.Logger  {

	return Log
}

