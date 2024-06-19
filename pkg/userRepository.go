package pkg

import (
	"fmt"
)

// UserRepository defines the interface for user persistence operations.
type UserRepository interface {
	FindUserByID(id string) error
}

// MySQLUserRepository is an implementation of UserRepository that uses MySQL.
type MySQLUserRepository struct{}

func NewMySQLUserRepository() *MySQLUserRepository {
	return &MySQLUserRepository{}
}

func (repo *MySQLUserRepository) FindUserByID(id string) error {
	fmt.Println("funding user by id")
	return nil
}
