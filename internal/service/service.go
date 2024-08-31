package service

import (
	"context"
	"errors"

	"github.com/mirhijinam/wb-l0/internal/models"
	"github.com/nats-io/stan.go"
)

type repository interface {
	Save(ctx context.Context, data models.OrderData) error
	PutIntoCache(ctx context.Context) (map[string]models.OrderData, error)
}

type Service struct {
	sc    stan.Conn
	repo  repository
	Cache map[string]models.OrderData
}

func New(sc stan.Conn, repo repository) Service {
	return Service{
		sc:    sc,
		repo:  repo,
		Cache: make(map[string]models.OrderData),
	}
}

func (s *Service) GetFromCache(uid string) (models.OrderData, error) {
	order, ok := s.Cache[uid]
	if !ok {
		return models.OrderData{}, errors.New("no order with such uid")
	}

	return order, nil
}

func (s *Service) PutIntoCache(ctx context.Context) error {
	cache, err := s.repo.PutIntoCache(ctx)
	if err != nil {
		return err
	}

	s.Cache = cache

	return nil
}
