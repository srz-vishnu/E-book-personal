package log

// import (
// 	"fmt"
// )

// // Global logger
// var logger Logger

// // Initialize default logger
// func init() {
// 	if logger == nil {
// 		NewZerologLogWrapper(map[string]interface{}{}, true)
// 	}
// }

// // GetLogger: Fetches the initialized logger that satisfies the Logger interface
// func GetLogger() Logger {
// 	return logger
// }

// // Info: Log an info message
// func Info(message string) {
// 	logger.Info(message)
// }

// // Debug: Log a debug message
// func Debug(message string) {
// 	logger.Debug(message)
// }

// // Warn: Log a warning message with an error
// func Warn(message string, err error) {
// 	logger.Warn(message, err)
// }

// // Error: Log an error message
// func Error(message string, err error) {
// 	logger.Error(message, err)
// }

// // Fatal: Log a fatal error message
// func Fatal(message string, err error) {
// 	logger.Fatal(message, err)
// }

// // Panic: Log a panic message
// func Panic(message string, err error) {
// 	logger.Panic(message, err)
// }

// // Infof: Format and log an info message
// func Infof(message string, args ...interface{}) {
// 	msg := fmt.Sprintf(message, args...)
// 	logger.Info(msg)
// }
