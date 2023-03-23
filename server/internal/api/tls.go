package api

import (
	"crypto/tls"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

// TLSServer - grpc server with tls.
type TLSServer struct {
	GRPCServer
	tlsConfig *tls.Config
}

// NewTLSServer - creates new TLSServer object.
func NewTLSServer(grpcOpts []grpc.ServerOption, certFile string, keyFile string) (*TLSServer, error) {
	tlsServer := &TLSServer{}

	err := tlsServer.loadTLSConfig(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	tlsCredentials := credentials.NewTLS(tlsServer.tlsConfig)
	grpcOpts = append(grpcOpts, grpc.Creds(tlsCredentials))
	server := grpc.NewServer(grpcOpts...)
	tlsServer.grpcServer = server
	return tlsServer, nil
}

// Start starts serving grpc connection on provided address.
func (s *TLSServer) Start(addr string) error {
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

func (s *TLSServer) loadTLSConfig(certFile string, keyFile string) error {
	serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("can't load certificate for server: %v", err)
	}

	s.tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return nil
}
