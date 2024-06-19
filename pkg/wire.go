//go:build wireinject

package pkg

import "github.com/google/wire"

func InitializeUserService() *UserService {
	wire.Build(NewUserService, NewMySQLUserRepository, wire.Bind(new(UserRepository), new(*MySQLUserRepository)))
	return nil
}
