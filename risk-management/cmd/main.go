package main

import (
	"fmt"
	"os"
	"time"

	commonlogger "github.com/Mohammed-Aadil/common-core/pkg/logger"
	"github.com/Mohammed-Aadil/common-core/pkg/signal"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/Mohammed-Aadil/risk-management/pkg/api"
	"github.com/Mohammed-Aadil/risk-management/pkg/persistence"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)

	hostname, _ := os.Hostname()
	fs.Int("http-port", 8080, "http listening port")
	fs.String("host-name", hostname, "service host name")
	fs.String("level", "info", "log level")
	fs.String("service-name", "risk-management-service", "service name")
	fs.Duration("http-read-timeout", 2*time.Second, "http server read timeout")
	fs.Duration("http-write-timeout", 5*time.Second, "http server write timeout")
	fs.Duration("server-shutdown-timeout", 5*time.Second, "server graceful shutdown timeout")
	fs.String("storage-backend", "inmemory", "storage backend for persistence layer")
	fs.Int("default-pagination-limit", 100, "default pagination limit")

	// parse cli config
	err := fs.Parse(os.Args[1:])
	switch err {
	case pflag.ErrHelp:
		os.Exit(0)
	case nil:
	default:
		fmt.Fprintf(os.Stderr, "%s", err)
		fs.PrintDefaults()
		os.Exit(2)
	}

	// map cli config to viper
	err = viper.BindPFlags(fs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}

	logger, stdlog := commonlogger.InitLogger(viper.GetString("level"))
	defer logger.Sync()
	defer stdlog()

	var config model.Config

	if err = viper.Unmarshal(&config); err != nil {
		logger.Fatal("unable to parse configs", zap.Error(err))
	}

	stopCh := signal.SetupSignalHandler()
	store := persistence.Init(&config, logger)
	api := api.NewAPI(&config, store, logger)
	api.StartServer(stopCh)
}
