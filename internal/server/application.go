package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Run() {
	// fmt.Print("Hello World")
	serverStopChannel := make(chan struct{})
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	app, _ := newUserServerApplication()
	s := &http.Server{Addr: ":8080", Handler: app.userserverHttpHandler}
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Error(err)
		}
		close(serverStopChannel)
	}()
	select {
	case <-signalCh:
		_ = s.Shutdown(context.Background())
	case <-serverStopChannel:
	}
}
