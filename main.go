package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/hans-m-song/archive-ingest/pkg/util"
	"github.com/hans-m-song/archive-provide/pkg/api"
	"github.com/hans-m-song/archive-provide/pkg/config"
	"github.com/hans-m-song/archive-provide/pkg/provide"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func startServer(wg *sync.WaitGroup, provider *provide.Provider) *http.Server {
	url := ":" + viper.GetString(config.ServerPort)

	server := &http.Server{Addr: url, Handler: api.NewApiHandler(provider)}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithField("err", err).Fatal("server exited with error")
		}
	}()

	logrus.WithField("url", url).Info("server started")

	return server
}

func main() {
	config.AugmentEnvSetup()

	provider, err := provide.NewProvider()
	if err != nil {
		logrus.WithField("err", err).Fatal("error creating database connection")
	}

	serverWg := &sync.WaitGroup{}
	server := startServer(serverWg, provider)

	cleanup := func() {
		if err := provider.Disconnect(); err != nil {
			logrus.WithField("err", err).Fatal("error disconnecting from database")
		}

		if err := server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			logrus.WithField("err", err).Fatal("error shutting down server")
		}
	}

	util.CatchSignal(cleanup)
	defer cleanup()

	serverWg.Wait()
}
