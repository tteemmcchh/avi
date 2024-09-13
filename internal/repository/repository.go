package repository

import (
	"context"
	"mini_url/internal/models"
)

type TenderRepository interface {
	Create(ctx context.Context, t *models.Tender) (*models.Tender, error)
}
