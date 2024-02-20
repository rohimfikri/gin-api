package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ZeroJSONLogger(ginMode *string, logLevel *string, logPretty *uint8) gin.HandlerFunc {
	// wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
	// 	fmt.Printf("Logger Dropped %d messages", missed)
	// })
	// log := zerolog.New(wr)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if strings.ToUpper(*ginMode) == "DEBUG" || strings.ToUpper(*logLevel) == "DEBUG" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if strings.ToUpper(*logLevel) == "INFO" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if strings.ToUpper(*logLevel) == "WARNING" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
	if *logPretty == 1 {
		log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
	logger := &log.Logger

	return func(c *gin.Context) {

		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Log using the params
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)
	}
}
