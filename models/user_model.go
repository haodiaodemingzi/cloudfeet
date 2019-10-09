package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	Model
	UserName string `gorm:"column:username;type:varchar(100);unique_index"`
	Password string
	Role     string
	Region   string
	Status   int
	Source   string
	quota    int
}

func (UserModel) TableName() string {
	return "user"
}

func (model *UserModel) Exist(id int) (bool, error) {
	err := db.Select("id").Where("id=?", id).First(model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if model.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (model *UserModel) Total(maps interface{}) (int, error) {
	var count int
	err := db.Model(model).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (model *UserModel) Count() (int, error) {
	var count int
	err := db.Model(model).Count(&count).Error
	if err != nil {
		return 0, nil
	}

	return count, nil
}

func (model *UserModel) Query(where interface{}, pageNum int, pageSize int) ([]UserModel,
	error) {
	var userModels []UserModel
	err := db.Model(model).Where(where).Offset(pageNum).Limit(pageSize).Find(&userModels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return userModels, nil
}

func (model *UserModel) Get(id int) (*UserModel, error) {
	var user UserModel
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func (model *UserModel) Select(where interface{}) (*UserModel, error) {
	var user UserModel
	err := db.Model(model).Where(where).Offset(0).Limit(1).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func (model *UserModel) Edit() error {
	return db.Save(model).Error
}

func (model *UserModel) Add() error {
	return db.Create(model).Error
}

func (model *UserModel) Delete() error {
	return db.Delete(model).Error
}

func (model *UserModel) DeleteByCnd(where interface{}) error {
	return db.Where(where).Delete(model).Error
}

func (model *UserModel) CleanAll() error {
	return db.Delete(model).Error
}

func (model *UserModel) FindOrCreate(domain string) error {
	return db.FirstOrCreate(model, UserModel{UserName: domain}).Error
}
