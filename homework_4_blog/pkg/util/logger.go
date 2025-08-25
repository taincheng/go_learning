package util

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局变量，方便调用
var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

// Config 日志配置
type Config struct {
	// 日志文件路径
	Filename string

	// 最大文件大小（MB）
	MaxSize int

	// 最大备份文件数
	MaxBackups int

	// 最大保存天数
	MaxAge int

	// 是否压缩旧日志
	Compress bool

	// 是否启用控制台输出
	EnableConsole bool

	// 是否输出到文件
	EnableFile bool

	// 日志级别（debug, info, warn, error, dpanic, panic, fatal）
	Level string
}

// 默认配置
var defaultConfig = Config{
	Filename:      "./logs/app.log",
	MaxSize:       10,
	MaxBackups:    7,
	MaxAge:        30,
	Compress:      true,
	EnableConsole: false,
	EnableFile:    true,
	Level:         "info",
}

// Init 初始化日志器
func Init(config ...Config) error {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = defaultConfig
	}

	// 解析日志级别
	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return fmt.Errorf("无效的日志级别 '%s': %v", cfg.Level, err)
	}

	// 日志编码配置（JSON 格式）
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 可读时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var needWriteSyncer []zapcore.WriteSyncer
	if cfg.EnableFile {
		// 创建日志目录
		if err := os.MkdirAll("./logs", 0755); err != nil {
			return fmt.Errorf("无法创建日志目录: %v", err)
		}

		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		})
		needWriteSyncer = append(needWriteSyncer, fileWriteSyncer)
	}
	if cfg.EnableConsole {
		consoleWriteSyncer := zapcore.AddSync(os.Stdout)
		needWriteSyncer = append(needWriteSyncer, consoleWriteSyncer)
	}

	// 输出到文件、控制台
	// 根据输出目标数量创建合适的 WriteSyncer
	var writeSyncer zapcore.WriteSyncer
	if len(needWriteSyncer) == 0 {
		// 如果没有启用任何输出，则默认输出到控制台
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		// 多个输出目标，使用 MultiWriteSyncer
		writeSyncer = zapcore.NewMultiWriteSyncer(needWriteSyncer...)
	}

	// 构建 core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 构建 logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 添加调用者信息
		zap.AddCallerSkip(1),              // 跳过一层（用于封装）
		zap.AddStacktrace(zap.ErrorLevel), // 错误级别自动记录堆栈
	)

	// 设置全局 logger
	zap.ReplaceGlobals(Logger)

	// 创建 SugaredLogger（方便使用）
	Sugar = Logger.Sugar()

	return nil
}

// Sync 刷写日志缓冲区（程序退出前调用）
func Sync() {
	_ = Logger.Sync()
}

// SetLevel 动态设置日志级别
func SetLevel(level string) error {
	atomicLevel := Logger.Level()
	if err := atomicLevel.UnmarshalText([]byte(level)); err != nil {
		return err
	}
	return nil
}

// Close 关闭日志（Sync + 清理）
func Close() {
	Sync()
}
