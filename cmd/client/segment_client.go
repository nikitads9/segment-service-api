package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAdress = "localhost:50051"

type History struct {
	Slug       string `csv:"slug"`
	AddedAt    string `csv:"added_at"`
	ExpireTime string `csv:"time_of_expire"`
}

// Это клиент, который можно запустить при наличии данных в таблице users_segments_junction
// он получает байты от сервера, демаршалирует их в массив структур, а массив структур записывает в csv файл
func main() {
	ctx := context.Background()
	//nolint
	con, err := grpc.Dial(grpcAdress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v\n", err.Error())
		return
	}
	defer con.Close()

	client := desc.NewUserV1ServiceClient(con)

	cl, err := client.GetUserHistoryCsv(ctx, &desc.GetUserHistoryCsvRequest{
		Id: int64(1),
	})
	if err != nil {
		log.Printf("failed to get history: %v\n", err.Error())
		return
	}

	res, err := cl.Recv()
	if err != nil {
		log.Printf("failed to receive history: %v\n", err.Error())
		return
	}

	result := []*History{}

	err = gocsv.Unmarshal(bytes.NewReader(res.Chunk), &result)
	if err != nil {
		log.Printf("failed to unmarshal bytes array file: %v\n", err.Error())
		return
	}

	historyFile, err := os.Create("clients.csv")
	if err != nil {
		log.Printf("failed to create file: %v\n", err.Error())
		return
	}
	defer historyFile.Close()

	err = gocsv.MarshalFile(&result, historyFile)
	if err != nil {
		log.Printf("failed to marshal the interface to file: %v\n", err.Error())
		return
	}
}
