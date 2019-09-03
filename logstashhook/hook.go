package logstashhook

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type Hook struct {
	writer    io.Writer
	formatter logrus.Formatter
}

type LogEntry struct {
	Timestamp string `json:"@timestamp"`
	Version   string `json:"@version"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	Module    string `json:"module"`
	Project   string `json:"project"`
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
	var logEntry LogEntry
	err = json.Unmarshal(dataBytes, &logEntry)
	if err != nil {
		return err
	}
	if logEntry.Level == "error" {
		logEntry.Level = "ERROR"
		formatJSON, err := json.Marshal(logEntry)
		if err != nil {
			return err
		}
		_, err = h.writer.Write(formatJSON)
		return err
	} else {
		_, err = h.writer.Write(dataBytes)
		return err
	}
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
