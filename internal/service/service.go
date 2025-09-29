package service

import "github.com/n0l3r/fiberplate/internal/repository"

type Service struct{}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
