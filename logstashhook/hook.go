package logstashhook

import (
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type Hook struct {
	writer    io.Writer
	formatter logrus.Formatter
}

func New(w io.Writer, f logrus.Formatter) logrus.Hook {
	return Hook{
		writer:    w,
		formatter: f,
	}
}

func (h Hook) Fire(e *logrus.Entry) error {
	dataBytes, err := h.formatter.Format(e)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(dataBytes)
	//if err != nil {
	//write tcp [::1]:60786->[::1]:51401: write: broken pipe
	//Failed to fire hook: dial tcp 127.0.0.1:51401: connect: connection refused
	//if strings.ContainsAny(err.Error(), "broken pipe"){
	//
	//}
	//}
	return err
}

func (h Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var entryPool = sync.Pool{
	New: func() interface{} {
		return &logrus.Entry{}
	},
}

func copyEntry(e *logrus.Entry, fields logrus.Fields) *logrus.Entry {
	ne := entryPool.Get().(*logrus.Entry)
	ne.Message = e.Message
	ne.Level = e.Level
	ne.Time = e.Time
	ne.Data = logrus.Fields{}
	for k, v := range fields {
		ne.Data[k] = v
	}
	for k, v := range e.Data {
		ne.Data[k] = v
	}
	return ne
}

func releaseEntry(e *logrus.Entry) {
	entryPool.Put(e)
}

type LogstashFormatter struct {
	logrus.Formatter
	logrus.Fields
}

var (
	logstashFields   = logrus.Fields{"@version": "1", "type": "log"}
	logstashFieldMap = logrus.FieldMap{
		logrus.FieldKeyTime: "@timestamp",
		logrus.FieldKeyMsg:  "message",
	}
)

func DefaultFormatter(fields logrus.Fields) logrus.Formatter {
	for k, v := range logstashFields {
		if _, ok := fields[k]; !ok {
			fields[k] = v
		}
	}

	return LogstashFormatter{
		Formatter: &logrus.JSONFormatter{FieldMap: logstashFieldMap},
		Fields:    fields,
	}
}

func (f LogstashFormatter) Format(e *logrus.Entry) ([]byte, error) {
	ne := copyEntry(e, f.Fields)
	dataBytes, err := f.Formatter.Format(ne)
	releaseEntry(ne)
	return dataBytes, err
}
