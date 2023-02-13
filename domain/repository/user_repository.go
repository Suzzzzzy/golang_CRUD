package repository

import (
	"context"
	"github.com/golang_CRUD/domain/entity"
)

type NewsRepositoryInterface interface {
	// SignUp 회원가입
	SignUp(ctx context.Context, user entity.Users) ([]*entity.Users, error)
	// SignIn 로그인
	SignIn(ctx context.Context, email string, password string) ([]*entity.Users, error)
}
