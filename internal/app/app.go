package app

import (
	"context"
	"os"
	"sooty-tern/internal/app/config"
	"sooty-tern/pkg/logger"

	"go.uber.org/dig"
)

type options struct {
	ConfigFile string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobalConfig(o.ConfigFile)
	handleError(err)

	cfg := config.GetGlobalConfig()

	logger.Printf(ctx, "service started , run mode：%s , 版本号：%s , 进程号：%d\n", cfg.RunMode, o.Version, os.Getpid())

	loggerCall, err := InitLogger()
	handleError(err)

	err = InitMonitor()
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}

	// 创建依赖注入容器
	container := dig.New()

	// init http server
	httpCall := InitHTTPServer(ctx, container)
	return func() {
		if httpCall != nil {
			httpCall()
		}
		if loggerCall != nil {
			loggerCall()
		}
	}
}
