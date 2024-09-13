package tender

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"mini_url/internal/models"
	"mini_url/internal/repository"
)

const (
	tableName = "tenders"

	organizationIdColumn  = "organization_id"
	nameColumn            = "name"
	descriptionColumn     = "description"
	serviceTypeColumn     = "service_type"
	statusColumn          = "status"
	creatorUsernameColumn = "creator_username"
	versionColumn         = "version"
)

type repo struct {
	db *pgxpool.Pool
}

//func NewPool(connURL string) (*pgxpool.Pool, error) {
//	config, err := pgxpool.ParseConfig(connURL)
//	if err != nil {
//		return nil, fmt.Errorf("unable to parse connection URL. %s", err)
//	}
//
//	pool, err := pgxpool.NewWithConfig(context.Background(), config)
//	if err != nil {
//		return nil, fmt.Errorf("unable to create connection pool: %w", err)
//	}
//
//	return pool, nil
//}

func NewRepository(db *pgxpool.Pool) repository.TenderRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, t *models.Tender) (*models.Tender, error) {
	//супер плохо, но ладно
	t.Status = "Created"
	t.Version = 1

	builder := squirrel.Insert(tableName).
		Columns(nameColumn, organizationIdColumn, descriptionColumn, serviceTypeColumn, creatorUsernameColumn, statusColumn, versionColumn).
		Values(t.Info.Name, t.Info.OrganizationId, t.Info.Description, t.Info.ServiceType, t.Info.CreatorUsername, t.Status, t.Version).
		Suffix("RETURNING \"id, created_at, updated_at\"")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRow(ctx, query, args...).Scan(&t.Id, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return t, err
}
