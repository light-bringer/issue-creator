package middleware

import (
	"fmt"
	"issue-creator/config"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log to file
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.Log_FILE_PATH
	logFileName := config.LOG_FILE_NAME
	//log file
	fileName := path.Join(logFilePath, logFileName)
	//write file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//instantiation
	logger := logrus.New()
	//Set output
	logger.Out = src
	//Set log level
	logger.SetLevel(logrus.DebugLevel)
	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
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
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// Log to MongoDB
func LoggerToMongo() gin.HandlerFunc {
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
