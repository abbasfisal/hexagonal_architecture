package repository

import (
	"errors"
	"hexagonal_architecture/internal/core/dto"
)

var (
	DuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	Insert(user dto.UserDTO) error
}
