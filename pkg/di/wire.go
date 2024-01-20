//go:build wireinject
// +build wireinject

package dd

import (
	"github.com/google/wire"
	http "github.com/kannan112/token-application/pkg/api"
	handler "github.com/kannan112/token-application/pkg/api/handler"
	config "github.com/kannan112/token-application/pkg/config"
	db "github.com/kannan112/token-application/pkg/db"
	repository "github.com/kannan112/token-application/pkg/repository"
	usecase "github.com/kannan112/token-application/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
