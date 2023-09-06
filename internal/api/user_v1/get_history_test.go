package user_v1

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	bf "google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context, api *Implementation) (client desc.UserV1ServiceClient, closer func()) {
	buffer := 101024 * 1024
	lis := bf.Listen(buffer)

	baseServer := grpc.NewServer()
	desc.RegisterUserV1ServiceServer(baseServer, api)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer = func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client = desc.NewUserV1ServiceClient(conn)

	return client, closer
}

func Test_GetHistory(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		userRepoMock = userRepoMocks.NewMockRepository(mock)
		api          = NewMockImplementation(Implementation{
			userService: user.NewMockUserService(userRepoMock)})
		userId = gofakeit.Int64()

		userErr = status.Error(codes.Internal, gofakeit.Phrase())

		validRequest = &desc.GetUserHistoryCsvRequest{
			Id: userId,
		}
	)

	res, _ := gofakeit.CSV(nil)
	bytes := bytes.NewBuffer(res)
	client, closer := server(ctx, api)
	defer closer()

	t.Run("success case", func(t *testing.T) {
		userRepoMock.EXPECT().GetUserHistoryCsv(ctx, userId).Return(bytes, nil).Times(1)

		out, err := client.GetUserHistoryCsv(ctx, validRequest)

		var outs []*desc.GetUserHistoryCsvResponse

		for {
			o, err := out.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			outs = append(outs, o)
		}

		require.Nil(t, err)
		require.Equal(t, res, outs[0].Chunk)
	})

	t.Run("error case", func(t *testing.T) {
		userRepoMock.EXPECT().GetUserHistoryCsv(ctx, userId).Return(nil, userErr).Times(1)

		out, _ := client.GetUserHistoryCsv(ctx, validRequest)
		_, err := out.Recv()
		require.Error(t, err)
		require.Equal(t, userErr.Error(), err.Error())
	})

}
