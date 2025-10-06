package main

import (
	"flag"
	"github.com/aabbuukkaarr8/internal/handler"
	"github.com/aabbuukkaarr8/internal/repository"
	"github.com/aabbuukkaarr8/internal/service"

	"github.com/BurntSushi/toml"
	"github.com/aabbuukkaarr8/internal/apiserver"
	"github.com/aabbuukkaarr8/internal/store"
	"github.com/wb-go/wbf/zlog"
)

var (
	configPath string
)

func main() {

	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()
	zlog.Init()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("config load error")
	}
	db := store.New()
	err = db.Open(config.Store.DatabaseURL)
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("db open error")
		return
	}

	//repo
	repo := repository.NewRepository(db)
	//service
	srv := service.NewService(repo)
	//handler
	handler := handler.NewHandler(srv)
	s := apiserver.New(config)
	s.ConfigureRouter(handler)
	if err = s.Run(); err != nil {
		panic(err)
	}

}
