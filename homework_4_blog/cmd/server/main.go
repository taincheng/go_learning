package main

import (
	"homework_4_blog/internal/router"
	"homework_4_blog/pkg/util"
)

func main() {
	// 初始化日志
	err := util.Init(util.Config{
		Filename:      "./logs/app.log",
		MaxSize:       10,
		MaxBackups:    5,
		MaxAge:        30,
		Compress:      true,
		EnableConsole: true,
		EnableFile:    false,
		Level:         "debug", // 开发时用 debug，生产用 info
	})
	if err != nil {
		panic(err)
	}

	r := router.SetupRouter()
	r.Run()
}
