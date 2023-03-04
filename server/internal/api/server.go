package api

import (
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers"
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/interceptors"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server *grpc.Server
}

func NewServer(authService handlers.AuthService, itemsService handlers.ItemsService) *Server {
	authInterceptor := interceptors.NewAuthorizationServerInterceptor(authService)
	server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary))

	itemsServer := handlers.NewItemsServer(itemsService, authService)
	items.RegisterItemsManagementServer(server, itemsServer)
	authServer := handlers.NewAuthServer(authService)
	auth.RegisterAuthManagementServer(server, authServer)

	return &Server{server: server}
}

func (s *Server) Start(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		if err := s.server.Serve(listen); err != nil {
			log.Fatal().Err(err).Msg("grpc server can not serve")
		}
	}()
	log.Info().Msg("starting server")

	return nil
}
