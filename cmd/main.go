package main

import (
	"context"
	"flag"
	"fmt"
	"ltgc-api-create-user/internal/client"
	endpoint "ltgc-api-create-user/internal/endpoints"
	h "ltgc-api-create-user/internal/handler"
	"ltgc-api-create-user/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

var (
	ctx        context.Context
	logger     log.Logger
	listenAddr string
)

func main() {

	var httpAddr = flag.String("http", ":8080", "HTTP listen address")

	var logger log.Logger

	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "create-user",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("main", "msg", "service started")
	defer level.Info(logger).Log("main", "msg", "service ended")

	logger.Log("Server is starting...")

	flag.Parse()

	ctx = context.Background()

	var svc service.Service
	{
		firestoreclient := client.NewFirestorerepository("firebase.json", "firebase", "users")
		svc = service.MakeService(logger, firestoreclient)

	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// endpoint := endpoint.MakeEndpoints(svc, logger)
	ep := endpoint.MakeServiceEndpoint(ctx, svc)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := h.NewHTTPServer(ctx, ep)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)

}
