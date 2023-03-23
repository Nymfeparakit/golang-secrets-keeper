package interceptors

import (
	"context"
	mock_handlers "github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers/mocks"
	pb "github.com/Nymfeparakit/gophkeeper/server/internal/api/interceptors/proto"
	"github.com/Nymfeparakit/gophkeeper/server/internal/services"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func grpcClientSetup(
	authServiceMock *mock_handlers.MockAuthService,
) (pb.DummyServiceClient, func()) {
	buffer := 1024 * 1024
	listen := bufconn.Listen(buffer)

	authInterceptor := NewAuthorizationServerInterceptor(authServiceMock)
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary))
	pb.RegisterDummyServiceServer(s, &DummyServer{})
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatal().Err(err).Msg("")
		}
		log.Info().Msg("starting grpc server")
	}()

	bufDialer := func(ctx context.Context, address string) (net.Conn, error) {
		return listen.Dial()
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	client := pb.NewDummyServiceClient(conn)

	return client, func() {
		if err := conn.Close(); err != nil {
			log.Error().Err(err).Msg("")
		}
	}
}

func TestAuthorizationServerInterceptor(t *testing.T) {
	token := "token"
	email := "test@mail.com"

	tests := []struct {
		name                  string
		setupMocks            func(authServiceMock *mock_handlers.MockAuthService)
		prepareRequestContext func(ctx context.Context) context.Context
		expResponse           *pb.DummyResponse
		expError              error
	}{
		{
			name: "token is valid",
			setupMocks: func(authServiceMock *mock_handlers.MockAuthService) {
				authServiceMock.EXPECT().ParseUserToken(token).Return(email, nil)
				authServiceMock.EXPECT().AddUserToContext(gomock.Any(), email).Return(context.Background())
			},
			prepareRequestContext: func(ctx context.Context) context.Context {
				ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
				return ctx
			},
		},
		{
			name: "token is not valid",
			setupMocks: func(authServiceMock *mock_handlers.MockAuthService) {
				authServiceMock.EXPECT().ParseUserToken(token).Return("", services.ErrInvalidAccessToken)
			},
			prepareRequestContext: func(ctx context.Context) context.Context {
				ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
				return ctx
			},
			expError: status.Error(codes.Unauthenticated, "Provided token is not valid"),
		},
		{
			name:       "token is not provided",
			setupMocks: func(authServiceMock *mock_handlers.MockAuthService) {},
			expError:   status.Error(codes.Unauthenticated, "Authentication credentials not provided"),
			prepareRequestContext: func(ctx context.Context) context.Context {
				return ctx
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			tt.setupMocks(authServiceMock)

			client, teardownFunc := grpcClientSetup(authServiceMock)
			defer teardownFunc()
			ctx := tt.prepareRequestContext(context.Background())

			_, err := client.DoNothing(ctx, &pb.DummyRequest{})

			assert.Equal(t, tt.expError, err)
		})
	}
}
