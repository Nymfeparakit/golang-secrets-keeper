package api

import (
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers"
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/interceptors"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/grpc"
)

// Server - grpc server.
type Server interface {
	Start(addr string) error
	Shutdown()
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}

// NewServer - creates new grpc server object.
func NewServer(enableHTTPS bool, authService handlers.AuthService, secretsService handlers.SecretsService) (Server, error) {
	authInterceptor := interceptors.NewAuthorizationServerInterceptor(authService)
	grpcOpts := []grpc.ServerOption{grpc.UnaryInterceptor(authInterceptor.Unary)}

	var server Server
	var err error
	if enableHTTPS {
		server, err = NewTLSServer(grpcOpts)
		if err != nil {
			return nil, err
		}
	} else {
		server = NewHTTPServer(grpcOpts)
	}

	itemsServer := handlers.NewSecretsServer(secretsService, authService)
	secrets.RegisterSecretsManagementServer(server, itemsServer)
	authServer := handlers.NewAuthServer(authService)
	auth.RegisterAuthManagementServer(server, authServer)

	return server, nil
}

// GRPCServer - grpc server.
type GRPCServer struct {
	grpcServer *grpc.Server
}

// RegisterService register provided service in grpc server.
func (s *GRPCServer) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.grpcServer.RegisterService(desc, impl)
}

// Shutdown - stop grpc server.
func (s *GRPCServer) Shutdown() {
	s.grpcServer.GracefulStop()
}
