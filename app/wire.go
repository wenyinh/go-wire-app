//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/wenyinh/go-wire-app/pkg/config"
)

func InitializeApp(config *config.AppConfiguration) (*App, func(), error) {
	wire.Build(wireSet, New)
	return nil, nil, nil
}
