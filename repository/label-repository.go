package repository

import (
	"github.com/zerodev/golang_api/entity"
	"gorm.io/gorm"
)

type LabelRepository interface {
	GetLabel(userID uint64) ([]entity.Label_get, error)
	InsertLabel(label entity.Label) (entity.Label, error)
	UpdateLabel(label entity.Label) (entity.Label, error)
	DeleteLabel(label entity.Label) (entity.Label, error)
}

type labelConnection struct {
	connection *gorm.DB
}

func NewLabelRepository(db *gorm.DB) LabelRepository {
	return &labelConnection{
		connection: db,
	}
}

func (db *labelConnection) GetLabel(userID uint64) ([]entity.Label_get, error) {
	var label []entity.Label_get
	err := db.connection.Model(&entity.Label{}).Select("labels.id_label, labels.name_label, labels.color_label, (SELECT COUNT(*) FROM tasks WHERE tasks.label_id = labels.id_label) as num_task").Where("user_id = ?", userID).Find(&label)
	// err := db.connection.Where("user_id = ?", userID).Find(&label)
	return label, err.Error
}

func (db *labelConnection) InsertLabel(label entity.Label) (entity.Label, error) {
	err := db.connection.Save(&label)
	db.connection.Preload("User").Find(&label)
	return label, err.Error
}

func (db *labelConnection) UpdateLabel(label entity.Label) (entity.Label, error) {
	err := db.connection.Save(&label)
	return label, err.Error
}

func (db *labelConnection) DeleteLabel(label entity.Label) (entity.Label, error) {
	err := db.connection.Delete(&label)
	return label, err.Error
}
