package account

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	uuid "github.com/satori/go.uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid := uuid.NewV4()
	id := uuid.String()
	user := &User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	return "Success", nil
}

func (s *service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("GetUser", email)

	return email, nil
}