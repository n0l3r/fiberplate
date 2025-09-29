package service

import "boilerplate/internal/repository"

type Service struct{}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
