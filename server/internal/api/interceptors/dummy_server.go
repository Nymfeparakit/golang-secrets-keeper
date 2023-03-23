package interceptors

import (
	"context"

	pb "github.com/Nymfeparakit/gophkeeper/server/internal/api/interceptors/proto"
)

// DummyServer заглушка для grpc сервера.
type DummyServer struct {
	pb.UnimplementedDummyServiceServer
}

// DoNothing заглушка для метода grpc сервера.
func (s *DummyServer) DoNothing(ctx context.Context, in *pb.DummyRequest) (*pb.DummyResponse, error) {
	return &pb.DummyResponse{}, nil
}
