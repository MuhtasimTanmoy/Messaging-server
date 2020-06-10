package logger

import (
	"fmt"
	"os"
	"time"
    "github.com/spf13/viper"
	"github.com/google/logger"
)

const LogsPath = "var/logs"

// Info log function
func Info(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()

		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)

		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

// 		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
// 		defer out.Close()

		out.Info(v...)
	}
}

// Infoln log function
func Infoln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Infoln(v...)
	}
}

// Infof log function
func Infof(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Infof(format, v...)
	}
}

// Warning log function
func Warning(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Warning(v...)
	}
}

// Warningln log function
func Warningln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Warningln(v...)
	}
}

// Warningf log function
func Warningf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Warningf(format, v...)
	}
}

// Error log function
func Error(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Error(v...)
	}
}

// Errorln log function
func Errorln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Errorln(v...)
	}
}

// Errorf log function
func Errorf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Errorf(format, v...)
	}
}

// Fatal log function
func Fatal(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Fatal(v...)
	}
}

// Fatalln log function
func Fatalln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Fatalln(v...)
	}
}

// Fatalf log function
func Fatalf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Hamster", false, false, lf)
		defer out.Close()

		out.Fatalf(format, v...)
	}
}
