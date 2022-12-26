package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	// e.Use(MiddlewareOne, MiddlewareTwo)
	e.Use(MiddlewareLogging)

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, useragent=${user_agent}\n",
	// }))

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	e.GET("/", func(ctx echo.Context) error {
		// fmt.Println("From handler")
		time.Sleep(10 * time.Millisecond)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})
	e.GET("/health", func(ctx echo.Context) error {
		time.Sleep(10 * time.Millisecond)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})
	e.GET("/health/:id", func(ctx echo.Context) error {
		time.Sleep(10 * time.Millisecond)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	e.Start(":4444")
}

func makeLogentry(c echo.Context) *logrus.Fields {
	if c == nil {
		return &logrus.Fields{
			"at": time.Now().Format(time.RFC3339),
		}
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.WarnLevel)
	}

	fields := logrus.Fields{
		"at":     time.Now().Format(time.RFC3339),
		"method": c.Request().Method,
		"ip":     c.Request().RemoteAddr,
	}
	return &fields
}

func addEntry(log *logrus.Fields, key string, value interface{}) {
	l := *log
	l[key] = value
}

func build(log *logrus.Fields) *logrus.Entry {
	return logrus.WithFields(*log)
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		now := time.Now()
		err := next(ctx)
		if err != nil {
			ctx.Error(err)
			return err
		}
		end := time.Since(now)
		fields := makeLogentry(ctx)
		addEntry(fields, "response_time", end.Milliseconds())
		addEntry(fields, "path", ctx.Request().URL.Path)

		build(fields).Warn()
		build(fields).Info()
		build(fields).Debug()
		// logrus.WithFields(*fields).Info("Hello")
		// makeLogentry(ctx).Warn("Warn Level")
		// makeLogentry(ctx).Info("Info Level")
		return nil
	}
}

func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("From middleware one")
		next(ctx)
		fmt.Println("After request from middleware one")
		return nil
	}
}
func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("From middleware two")
		next(ctx)
		fmt.Println("After request from middleware two")
		return nil
	}
}
