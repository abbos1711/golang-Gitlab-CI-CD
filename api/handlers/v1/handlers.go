package handlers

import (
	t "gitlab.com/tizim-back/api/tokens"
	"gitlab.com/tizim-back/config"
	"gitlab.com/tizim-back/pkg/logger"
	"gitlab.com/tizim-back/storage"
)

type handlerV1 struct {
	cfg        *config.Config
	Storage    storage.StorageI
	log        logger.Logger
	jwthandler t.JWTHandler
}

type HandlerV1Options struct {
	Cfg        *config.Config
	Storage    *storage.StorageI
	Log        logger.Logger
	JWTHandler t.JWTHandler
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:        options.Cfg,
		Storage:    *options.Storage,
		log:        options.Log,
		jwthandler: options.JWTHandler,
	}
}
