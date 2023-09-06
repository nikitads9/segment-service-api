package app

import (
	"context"
	"fmt"
	"log"

	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	descSegmentV1 "github.com/nikitads9/segment-service-api/pkg/segment_api"
	descUserV1 "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	pathConfig      string
	grpcServer      *grpc.Server
	mux             *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)
	return a, err
}

func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	var err error

	defer a.serviceProvider.db.Close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = a.startGRPC()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = a.startHTTP()
	}()

	wg.Wait()

	return err

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))

	descSegmentV1.RegisterSegmentV1ServiceServer(a.grpcServer, a.serviceProvider.GetSegmentImpl(ctx))
	descUserV1.RegisterUserV1ServiceServer(a.grpcServer, a.serviceProvider.GetUserImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := descSegmentV1.RegisterSegmentV1ServiceHandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().Grpc.Port, opts)
	if err != nil {
		return err
	}

	err = descUserV1.RegisterUserV1ServiceHandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().Grpc.Port, opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) startGRPC() error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().Grpc.Port)
	if err != nil {
		return fmt.Errorf("failed to create listener %v", err.Error())
	}

	defer list.Close()

	if err = a.grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to process gRPC server: %s", err.Error())
	}

	return nil
}

func (a *App) startHTTP() error {
	return http.ListenAndServe(a.serviceProvider.GetConfig().Http.Port, a.mux)
}
