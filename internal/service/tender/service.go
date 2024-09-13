package tender

import (
	"mini_url/internal/repository"
	def "mini_url/internal/service"
)

var _ def.TenderService = (*serv)(nil)

type serv struct {
	tenderRepository repository.TenderRepository
}

//func (s serv) Create(ctx context.Context, info *models.Tender) (*models.Tender, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewService(tenderRepository repository.TenderRepository) *serv {
	return &serv{tenderRepository: tenderRepository}
}
