package models

import (
	"github.com/jinzhu/gorm"
)


func (PacModel) TableName() string {
	return "pac"
}

func (model *PacModel) Exist(id int) (bool, error) {
	err := db.Select("id").Where("id=?", id).First(model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if model.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (model *PacModel) Total(maps interface{}) (int, error) {
	var count int
	err := db.Model(model).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (model *PacModel) Count() (int, error) {
	var count int
	err := db.Model(model).Count(&count).Error
	if err != nil {
		return 0, nil
	}

	return count, nil
}

func (model *PacModel) Query(where interface{}, pageNum int, pageSize int) ([]PacModel,
	error) {
	var pacModels []PacModel
	err := db.Model(model).Where(where).Offset(pageNum).Limit(pageSize).Find(&pacModels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return pacModels, nil
}

func (model *PacModel) Get(id int) (*PacModel, error) {
	var pac PacModel
	err := db.Where("id = ?", id).First(&pac).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pac, nil
}

func (model *PacModel) Select(where interface{}) (*PacModel, error) {
	var pac PacModel
	err := db.Model(model).Where(where).Offset(0).Limit(1).First(&pac).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pac, nil
}

func (model *PacModel) Edit() error {
	return db.Save(model).Error
}

func (model *PacModel) Add() error {
	return db.Create(model).Error
}

func (model *PacModel) Delete() error {
	return db.Delete(model).Error
}

func (model *PacModel) DeleteByCnd(where interface{}) error {
	return db.Where(where).Delete(model).Error
}

func (model *PacModel) CleanAll() error {
	return db.Delete(model).Error
}

func (model *PacModel) FindOrCreate(domain string) error {
	return db.FirstOrCreate(model, PacModel{Domain: domain}).Error
}
