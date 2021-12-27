package handler

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	//instantiation
	logger := logrus.New()
	//Set log level
	logger.SetLevel(logrus.DebugLevel)
	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		// initialise log writer
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
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
		// Response status
		respStatus := c.Writer.Status()
		// Response length in bytes
		respContLength := c.Writer.Size()
		// Response headers
		respHeaders := c.Writer.Header()
		// Log format
		logger.Infof("\n| %3d | %13v | %15s | %s | %s |"+
			"\nResponse status: %d"+
			"\nResponse content length: %d"+
			"\nResponse headers: %v",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			respStatus,
			respContLength,
			respHeaders,
		)
	}
}
