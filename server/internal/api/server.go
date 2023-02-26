package api

import (
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server *grpc.Server
}

func NewServer() *Server {
	server := grpc.NewServer()
	return &Server{server: server}
}

func (s *Server) RegisterHandlers(itemsService handlers.ItemsService, authService handlers.AuthService) {
	itemsServer := handlers.NewItemsServer(itemsService, authService)
	items.RegisterItemsManagementServer(s.server, itemsServer)
	authServer := handlers.NewAuthServer(authService)
	auth.RegisterAuthManagementServer(s.server, authServer)
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
