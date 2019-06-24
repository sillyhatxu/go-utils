package logconfig

import (
	"fmt"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	logger "log"
	"net"
	"os"
	"time"
)

type DefaultFieldHook struct {
	project string
	module  string
}

func (h *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *DefaultFieldHook) Fire(e *log.Entry) error {
	e.Data["project"] = h.project
	e.Data["module"] = h.module
	return nil
}

type LogConfig struct {
	LogLevel        log.Level
	ReportCaller    bool
	Project         string
	Module          string
	OpenLogstash    bool
	LogstashAddress string
	OpenLogfile     bool
	FilePath        string
}

func (lc LogConfig) String() string {
	return fmt.Sprintf(`LogConfig{Project='%s', Module='%s', OpenLogstash=%t, LogstashAddress='%s', OpenLogfile=%t, FilePath='%s'}`, lc.Project, lc.Module, lc.OpenLogstash, lc.LogstashAddress, lc.OpenLogfile, lc.FilePath)
}

func (lc LogConfig) InitialLogConfig() {
	logger.Println("InitialLogConfig :", lc)
	logFormatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyTime: "@timestamp",
		},
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(lc.LogLevel)
	log.SetReportCaller(lc.ReportCaller)
	log.SetFormatter(logFormatter)
	log.AddHook(&DefaultFieldHook{project: lc.Project, module: lc.Module})
	if lc.OpenLogstash {
		conn, err := net.Dial("tcp", lc.LogstashAddress)
		if err != nil {
			logger.Panicf("net.Dial('tcp', %v); Error : %v", lc.LogstashAddress, err)
		}
		hook := logrustash.New(conn, logrustash.DefaultFormatter(log.Fields{"project": lc.Project, "module": lc.Module}))
		log.AddHook(hook)
	}
	if lc.OpenLogfile {
		path := lc.FilePath + lc.Module + ".log"
		WithMaxAge := time.Duration(876000) * time.Hour
		WithRotationTime := time.Duration(24) * time.Hour
		infoWriter, err := rotatelogs.New(
			lc.FilePath+"info.log.%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(WithMaxAge),
			rotatelogs.WithRotationTime(WithRotationTime),
		)
		if err != nil {
			logger.Panicf("rotatelogs.New [info writer] error; Error : %v", err)
		}
		errorWriter, err := rotatelogs.New(
			lc.FilePath+"error.log.%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(WithMaxAge),
			rotatelogs.WithRotationTime(WithRotationTime),
		)
		if err != nil {
			logger.Panicf("rotatelogs.New [error writer] error; Error : %v", err)
		}
		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				log.InfoLevel:  infoWriter,
				log.WarnLevel:  infoWriter,
				log.ErrorLevel: infoWriter,
			},
			logFormatter,
		))
		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				log.WarnLevel:  errorWriter,
				log.ErrorLevel: errorWriter,
			},
			logFormatter,
		))
	}
}
