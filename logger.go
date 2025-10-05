/*
TeaBrewSim ログ管理
構造化ログとログレベル管理
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLevel ログレベル
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// LogEntry ログエントリ
type LogEntry struct {
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Timestamp time.Time              `json:"timestamp"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// Logger ログ管理構造体
type Logger struct {
	level  LogLevel
	writer io.Writer
}

// NewLogger 新しいロガーを作成
func NewLogger(level LogLevel, writer io.Writer) *Logger {
	return &Logger{
		level:  level,
		writer: writer,
	}
}

// SetLevel ログレベルを設定
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// SetWriter ログ出力先を設定
func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

// log 内部ログ関数
func (l *Logger) log(level LogLevel, message string, fields map[string]interface{}) {
	if level < l.level {
		return
	}

	levelStr := l.levelToString(level)
	entry := LogEntry{
		Level:     levelStr,
		Message:   message,
		Timestamp: time.Now(),
		Fields:    fields,
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ログのシリアライズに失敗: %v\n", err)
		return
	}

	fmt.Fprintln(l.writer, string(jsonData))
}

// levelToString ログレベルを文字列に変換
func (l *Logger) levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Debug デバッグログ
func (l *Logger) Debug(message string, fields ...map[string]interface{}) {
	var mergedFields map[string]interface{}
	if len(fields) > 0 {
		mergedFields = fields[0]
	}
	l.log(DEBUG, message, mergedFields)
}

// Info 情報ログ
func (l *Logger) Info(message string, fields ...map[string]interface{}) {
	var mergedFields map[string]interface{}
	if len(fields) > 0 {
		mergedFields = fields[0]
	}
	l.log(INFO, message, mergedFields)
}

// Warn 警告ログ
func (l *Logger) Warn(message string, fields ...map[string]interface{}) {
	var mergedFields map[string]interface{}
	if len(fields) > 0 {
		mergedFields = fields[0]
	}
	l.log(WARN, message, mergedFields)
}

// Error エラーログ
func (l *Logger) Error(message string, fields ...map[string]interface{}) {
	var mergedFields map[string]interface{}
	if len(fields) > 0 {
		mergedFields = fields[0]
	}
	l.log(ERROR, message, mergedFields)
}

// Fatal 致命的エラーログ
func (l *Logger) Fatal(message string, fields ...map[string]interface{}) {
	var mergedFields map[string]interface{}
	if len(fields) > 0 {
		mergedFields = fields[0]
	}
	l.log(FATAL, message, mergedFields)
	os.Exit(1)
}

// ParseLogLevel 文字列からログレベルを解析
func ParseLogLevel(levelStr string) LogLevel {
	switch levelStr {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INFO
	}
}

// グローバルロガー
var globalLogger *Logger

// InitLogger ロガーを初期化
func InitLogger(level LogLevel, writer io.Writer) {
	globalLogger = NewLogger(level, writer)
}

// GetLogger グローバルロガーを取得
func GetLogger() *Logger {
	if globalLogger == nil {
		globalLogger = NewLogger(INFO, os.Stdout)
	}
	return globalLogger
}

// LogMiddleware Gin用のログミドルウェア
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// リクエストを処理
		c.Next()

		// ログエントリを作成
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		bodySize := c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		fields := map[string]interface{}{
			"method":     method,
			"path":       path,
			"status":     statusCode,
			"latency":    latency.String(),
			"client_ip":  clientIP,
			"body_size":  bodySize,
		}

		logger := GetLogger()
		if statusCode >= 400 {
			logger.Error("HTTP Request", fields)
		} else {
			logger.Info("HTTP Request", fields)
		}
	}
}
