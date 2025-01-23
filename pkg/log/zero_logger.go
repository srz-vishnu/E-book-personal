package log

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/rs/zerolog"
// )

// // ZerologStdLogger is the actual logger that implements the Logger interface
// type ZerologStdLogger struct {
// 	Fields map[string]interface{}
// 	logger *zerolog.Logger
// }

// // Info: Log an info level entry
// func (l *ZerologStdLogger) Info(message string) {
// 	payloadFields := l.Fields
// 	l.logger.Info().Fields(payloadFields).Msg(message)
// }

// // InfoF: Log a formatted info message
// func (l *ZerologStdLogger) InfoF(message string, args ...interface{}) {
// 	msg := fmt.Sprintf(message, args...)
// 	payloadFields := l.Fields
// 	l.logger.Info().Fields(payloadFields).Msg(msg)
// }

// // Debug: Log a debug message
// func (l *ZerologStdLogger) Debug(message string) {
// 	payloadFields := l.Fields
// 	l.logger.Debug().Fields(payloadFields).Msg(message)
// }

// // Warn: Log a warning message
// func (l *ZerologStdLogger) Warn(message string, err error) {
// 	payloadFields := l.Fields
// 	l.logger.Warn().Err(err).Fields(payloadFields).Msg(message)
// }

// // Error: Log an error message
// func (l *ZerologStdLogger) Error(message string, err error) {
// 	payloadFields := l.Fields
// 	l.logger.Error().Err(err).Fields(payloadFields).Msg(message)
// }

// // Fatal: Log a fatal message
// func (l *ZerologStdLogger) Fatal(message string, err error) {
// 	payloadFields := l.Fields
// 	l.logger.Fatal().Err(err).Fields(payloadFields).Msg(message)
// }

// // Panic: Log a panic message
// func (l *ZerologStdLogger) Panic(message string, err error) {
// 	payloadFields := l.Fields
// 	l.logger.Panic().Err(err).Fields(payloadFields).Msg(message)
// }

// // LogRequest: Logs the details of an incoming HTTP request
// func (l *ZerologStdLogger) LogRequest(r *http.Request) {
// 	requestMetaData := map[string]interface{}{
// 		"method": r.Method,
// 		"url":    r.URL.String(),
// 		"header": r.Header,
// 	}

// 	l.logger.Info().
// 		Fields(l.Fields).
// 		Fields(requestMetaData).
// 		Msg("received request")
// }

// // LogResponse: Logs the response details after a request is processed
// func (l *ZerologStdLogger) LogResponse(status, sz, ms int, body string, r *http.Request) {
// 	responseFields := map[string]interface{}{
// 		"status_code":   status,
// 		"size":          sz,
// 		"response_time": ms,
// 		"body":          body,
// 	}
// 	l.logger.Info().
// 		Fields(l.Fields).
// 		Fields(responseFields).
// 		Msg("sent response")
// }
