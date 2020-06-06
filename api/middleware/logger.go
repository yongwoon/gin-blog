package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Log to file
func LoggerToFile() gin.HandlerFunc {
	logFilePath := os.Getenv("LOG_FILE_PATH")
	logFileName := os.Getenv("LOG_FILE_NAME")
	fileName := path.Join(logFilePath, logFileName)

	// write file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// instantiation
	logger := logrus.New()
	// Set output
	logger.Out = src
	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	// Set rotatelogs
	logWriter, err := rotatelogs.New(
		// Split file name
		fileName+".%Y%m%d.log",
		// Generate soft chain, point to the latest log file
		rotatelogs.WithLinkName(fileName),
		// Set maximum save time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// Set log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// Add Hook
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

// Log to Mysql
func LoggerToMysql() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Log to ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Logging to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
