package repository

import "go.uber.org/fx"

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(NewUserRepository),
)
