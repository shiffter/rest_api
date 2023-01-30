package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/test_task/rest_api/internal/app/apiserver"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf-path", "configs/apiserver.toml", "path to CONFIG file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(confPath, config)
	fmt.Println(config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
