package signal

import (
	"os"
	"os/signal"
)

/*
We need mechanism to gracefully and forcefully stop the service
We could utilize os signal for interrupt, once we recieve it we will close stopCh
stopCh will notify the service and proceed with graceful closure
In posix env, user can interrupt it twice to forcefully close service
In posix case we will termiate the service directly
*/
func SetupSignalHandler() chan struct{} {
	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, shutdownSignals...)

	go func() {
		<-sigCh
		close(stopCh)
		<-sigCh
		os.Exit(1)
	}()

	return stopCh
}
