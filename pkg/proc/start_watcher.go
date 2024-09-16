package proc

import (
	"time"

	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/clients/service"
	"github.com/deveshmishra34/groot/pkg/utils"
)

func StartWatcher() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	interval := utils.IntFromStr(config.WatcherSleepInterval)

	go func() {
		// This is a sample watcher
		// Command execution goes here ...

		logger.Info("Watcher started")
		for {
			// Watcher logic goes here ...
			logger.Info("Watcher running...")
			// Break the loop after 10 iterations
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}

	}()
}
