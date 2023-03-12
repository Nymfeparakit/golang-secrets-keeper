package api

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"math/big"
	"net"
	"time"
)

type TLSServer struct {
	GRPCServer
	tlsConfig *tls.Config
}

func NewTLSServer(grpcOpts []grpc.ServerOption) (*TLSServer, error) {
	tlsServer := &TLSServer{}

	err := tlsServer.createTLSConfig()
	if err != nil {
		return nil, err
	}

	tlsCredentials := credentials.NewTLS(tlsServer.tlsConfig)
	grpcOpts = append(grpcOpts, grpc.Creds(tlsCredentials))
	server := grpc.NewServer(grpcOpts...)
	tlsServer.grpcServer = server
	return tlsServer, nil
}

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

func (s *TLSServer) createTLSConfig() error {
	if s.tlsConfig == nil {
		cert, err := s.createCertificate()
		if err != nil {
			return err
		}
		s.tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	}

	return nil
}

func (s *TLSServer) createCertificate() (tls.Certificate, error) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1338),
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return tls.Certificate{}, err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey)
	if err != nil {
		return tls.Certificate{}, err
	}

	var certPem bytes.Buffer
	err = pem.Encode(&certPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		return tls.Certificate{}, err
	}

	var privateKeyPEM bytes.Buffer
	err = pem.Encode(&privateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return tls.Certificate{}, err
	}

	tlsCert, err := tls.X509KeyPair(certPem.Bytes(), privateKeyPEM.Bytes())
	if err != nil {
		return tls.Certificate{}, err
	}

	return tlsCert, nil
}
