package services

import (
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/daos"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/models"
)

type DadadService struct {
	dadadDao *daos.DadadDao
}

func NewDadadService() (*DadadService, error) {
	dadadDao, err := daos.NewDadadDao()
	if err != nil {
		return nil, err
	}
	return &DadadService{
		dadadDao: dadadDao,
	}, nil
}

func (dadadService *DadadService) CreateDadad(dadad *models.Dadad) (*models.Dadad, error) {
	return dadadService.dadadDao.CreateDadad(dadad)
}

func (dadadService *DadadService) GetDadad(id int64) (*models.Dadad, error) {
	return dadadService.dadadDao.GetDadad(id)
}
