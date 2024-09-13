package service

import (
	"context"
	"mini_url/internal/models"
)

type TenderService interface {
	Create(ctx context.Context, info *models.Tender) (*models.Tender, error)
}
