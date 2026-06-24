package starter

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/haolie/goUtil/config"
	"github.com/haolie/goUtil/logutil"
	"github.com/haolie/goUtil/park"
)

const configKeyLogPath = "log_path"

// Start 是应用程序的统一启动入口，负责以下流程：
//  1. 从可执行文件所在目录加载 Config.toml
//  2. 读取 log_path 配置并初始化日志系统
//  3. 调用 fn 执行业务初始化，fn 内的 panic 会被捕获并记录，返回的 error 会被记录并取消 ctx
//  4. 监听系统退出信号（SIGHUP / SIGINT / SIGTERM / SIGQUIT），收到后取消 ctx
//  5. 阻塞直到 ctx 被取消
func Start(fn func(context.Context) error) {
	if err := config.Load(path.Dir(os.Args[0])); err != nil {
		panic("load config file error: " + err.Error())
	}

	logPath, exists := config.GetValue[string](configKeyLogPath)
	if !exists {
		panic("log_path not found in config")
	}
	logutil.InitLog(logPath)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for sig := range sigCh {
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logutil.Info(fmt.Sprintf("exit signal: %v", sig))
				cancel()
			}
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				logutil.Error(fmt.Sprintf("%v", r))
			}
		}()

		errList := start(ctx, fn)
		if len(errList) > 0 {
			for _, err := range errList {
				logutil.Error(err.Error())
			}

			cancel()
		}
	}()

	<-ctx.Done()
}

func start(ctx context.Context, fn func(context.Context) error) []error {
	errList := park.RunBeforeStart(ctx)
	if len(errList) > 0 {
		return errList
	}

	errList = park.RunStart(ctx)
	if len(errList) > 0 {
		return errList
	}

	err := fn(ctx)
	if err != nil {
		errList = append(errList, err)
	}

	return errList
}
