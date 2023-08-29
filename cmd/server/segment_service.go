package main

import (
	"context"
	"flag"
)

var pathConfig string

func init() {
	pathConfig = *flag.String("config path", "config.yml", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

}
