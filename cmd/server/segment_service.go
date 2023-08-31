package main

import (
	"context"
	"flag"
	"log"

	"github.com/nikitads9/segment-service-api/internal/app"
)

var pathConfig string

// "C:\\Users\\swnik\\Desktop\\projects\\segment-service-api\\
func init() {
	flag.StringVar(&pathConfig, "config", "config.yml", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	app, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("failed to start app err:%s\n", err.Error())
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run application err:%s\n", err.Error())
	}
}
