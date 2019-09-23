package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sooty-tern/internal/app"
	"sooty-tern/pkg/logger"
	"sooty-tern/pkg/util"
	"sync/atomic"
	"syscall"
)

// VERSION 版本号，
var VERSION = "v0.0.1"

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "", "the path to config file")
}

func main() {
	flag.Parse()

	if configFile == "" {
		panic("please specify the config file path")
	}

	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := logger.NewTraceIDContext(context.Background(), util.MustUUID())
	span := logger.StartSpanWithCall(ctx)

	call := app.Init(
		ctx,
		app.SetConfigFile(configFile),
		app.SetVersion(VERSION),
	)

	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		span().Printf("grab the signal [%s]", sig.String())
	}

	if call != nil {
		call()
	}
	span().Printf("service exit...")
	os.Exit(int(atomic.LoadInt32(&state)))
}
