package service

import (
	"hexagonal_architecture/internal/core/model/request"
	"hexagonal_architecture/internal/core/model/response"
)

// primary port

type UserService interface {
	SignUp(req *request.SignUpRequest) *response.Response
}
