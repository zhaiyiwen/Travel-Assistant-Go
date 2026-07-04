// Package main 是旅行服务 (Kitex RPC) 的入口点。
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fuyx-goai/Travel-Assistant/pkg/config"
	"github.com/fuyx-goai/Travel-Assistant/pkg/logger"
	"go.uber.org/zap"
)

func init() {
	configPath := flag.String("config", "configs", "配置文件目啝路径")
	env := flag.String("env", "", "运行环境：dev, test, prod")
	servicePort := flag.Int("port", 9001, "服务 gRPC 端口")
	flag.Parse()

	cfg, err := config.Load(*configPath, *env)
	if err != nil {
		fmt.Printf("[致命错误] 旅行服务：配置加载失败: %v\n", err)
		os.Exit(1)
	}

	loggerCfg := logger.LoggerConfig{
		Level:       cfg.Logger.Level,
		Format:      cfg.Logger.Format,
		MaxSizeMB:   cfg.Logger.MaxSizeMB,
		MaxBackups:  cfg.Logger.MaxBackups,
		MaxAgeDays:  cfg.Logger.MaxAgeDays,
		Compress:    cfg.Logger.Compress,
		OutputPaths: []string{"stdout", "logs/travel-service.log"},
		ErrorOutput: []string{"logs/travel-error.log"},
	}

	logger.Init(loggerCfg)
	log := logger.Get()
	log.Info("[旅行服务已启动]", zap.Int("port", *servicePort))

	_ = cfg
	_ = log
	_ = servicePort
}

func main() { _ = logger.Get() }