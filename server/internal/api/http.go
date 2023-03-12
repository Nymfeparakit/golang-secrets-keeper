package api

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type HTTPServer struct {
	GRPCServer
}

func NewHTTPServer(grpcOpts []grpc.ServerOption) *HTTPServer {
	httpServer := &HTTPServer{}
	server := grpc.NewServer(grpcOpts...)
	httpServer.grpcServer = server
	return httpServer
}

func (s *HTTPServer) Start(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		if err := s.grpcServer.Serve(listen); err != nil {
			log.Fatal().Err(err).Msg("grpc grpcServer can not serve")
		}
	}()
	log.Info().Msg("starting grpcServer")

	return nil
}
