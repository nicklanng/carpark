package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nicklanng/carpark/api/handlers"
	"github.com/nicklanng/carpark/logging"
)

// StartHTTPSServer creates an HTTPS server on the specified address using the supplied certificate and key
func StartHTTPSServer(addr string, TLSCertPath, TLSKeyPath string, routes http.Handler) (server *http.Server) {
	logger, dispose := logging.WarnLogger()

	handler := handlers.EnsureRequestID(handlers.LoggingAndMetrics(routes))

	server = &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		ErrorLog:     logger,
	}

	go func() {
		logging.Info("Started HTTPS Server on " + addr)
		err := server.ListenAndServeTLS(TLSCertPath, TLSKeyPath)
		dispose()
		if err != nil {
			if err == http.ErrServerClosed {
				logging.Warn(err.Error())
				return
			}
			logging.Fatal(err.Error())
		}
	}()

	return
}

// GracefulShutdownOnSignal will run the supplied shutdown procedure when one of the specified signals is received
func GracefulShutdownOnSignal(signals []syscall.Signal, shutdownProcedure func()) {
	stopChan := make(chan os.Signal)
	for _, s := range signals {
		signal.Notify(stopChan, s)
	}

	<-stopChan
	logging.Info("Shutting down...")

	shutdownProcedure()

	logging.Info("Shut down complete")
}

// ShutdownHTTPServer will allow all connections to drain before closing the server
func ShutdownHTTPServer(server *http.Server, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	server.Shutdown(ctx)
}
