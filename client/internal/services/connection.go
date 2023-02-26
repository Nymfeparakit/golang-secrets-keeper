package services

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToServer(serverAddr string) (*grpc.ClientConn, error) {
	return grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
