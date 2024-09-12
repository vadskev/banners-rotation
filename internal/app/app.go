package app

import (
	"context"
	"log"
	"net"

	"github.com/vadskev/banners-rotation/internal/config"
	"github.com/vadskev/banners-rotation/internal/logger"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.loadDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) loadDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.loadConfig,
		a.loadServiceProvider,
		a.initGRPCServer,
		a.loadLogger,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) loadConfig(_ context.Context) error {
	err := config.Load()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) loadLogger(_ context.Context) error {
	err := logger.Init(a.serviceProvider.LogConfig().Level())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) loadServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)
	//desc.RegisterRotationServer(a.grpcServer, a.serviceProvider.RotationImpl(ctx))
	desc.RegisterRotationServer(a.grpcServer, a.serviceProvider.RotationImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}
