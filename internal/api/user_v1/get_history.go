package user_v1

import (
	"context"
	"strconv"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/grpc/metadata"
)

func (i *Implementation) GetUserHistoryCsv(req *desc.GetUserHistoryCsvRequest, responseStream desc.UserV1Service_GetUserHistoryCsvServer) error {
	ctx := context.Background()
	buffer, err := i.userService.GetUserHistoryCsv(ctx, req.GetId())
	if err != nil {
		return err
	}

	buff := buffer.Bytes()
	resp := &desc.GetUserHistoryCsvResponse{
		Chunk: buff,
	}

	//изменение заголовка ответа
	header1 := metadata.Pairs("Content-Type", "text/csv")
	header2 := metadata.Pairs("content-length", strconv.Itoa(len(buff)))

	err = responseStream.SetHeader(header1)
	if err != nil {
		return err
	}
	err = responseStream.SetHeader(header2)
	if err != nil {
		return err
	}
	err = responseStream.Send(resp)
	if err != nil {
		return err
	}

	return nil
}
