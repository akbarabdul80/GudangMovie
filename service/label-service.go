package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/repository"
)

type LabelService interface {
	GetLabel(userID uint64) ([]entity.Label_get, error)
	GetLabelByID(userID uint64, labelID uint64) (entity.Label_get, error)
	CreateLabel(label dto.LabelCreateDTO) (entity.Label, error)
	UpdateLabel(label dto.LabelUpdateDTO) (entity.Label, error)
	DeleteLabel(label dto.LabelDeleteDTO) (entity.Label, error)
}

type labelService struct {
	labelRepository repository.LabelRepository
}

func NewLabelService(labelRepo repository.LabelRepository) LabelService {
	return &labelService{
		labelRepository: labelRepo,
	}
}

func (service *labelService) GetLabel(userID uint64) ([]entity.Label_get, error) {
	res, err := service.labelRepository.GetLabel(userID)
	return res, err
}

func (service *labelService) GetLabelByID(userID uint64, labelID uint64) (entity.Label_get, error) {
	res, err := service.labelRepository.GetLabelByID(userID, labelID)
	return res, err
}

func (service *labelService) CreateLabel(label dto.LabelCreateDTO) (entity.Label, error) {
	labelToInsert := entity.Label{}
	err := smapping.FillStruct(&labelToInsert, smapping.MapFields((&label)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.labelRepository.InsertLabel(labelToInsert)

	return res, err
}

func (service *labelService) UpdateLabel(label dto.LabelUpdateDTO) (entity.Label, error) {
	labelToUpdate := entity.Label{}

	err := smapping.FillStruct(&labelToUpdate, smapping.MapFields((&label)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.labelRepository.UpdateLabel(labelToUpdate)

	return res, err
}

func (service *labelService) DeleteLabel(label dto.LabelDeleteDTO) (entity.Label, error) {
	labelToDelete := entity.Label{}

	err := smapping.FillStruct(&labelToDelete, smapping.MapFields((&label)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.labelRepository.DeleteLabel(labelToDelete)

	return res, err
}
