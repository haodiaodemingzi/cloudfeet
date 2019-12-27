package models

import (
	"github.com/jinzhu/gorm"
)


func (ProxyModel) TableName() string {
	return "proxy"
}

func (model *ProxyModel) Exist(id int) (bool, error) {
	err := db.Select("id").Where("id=?", id).First(model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if model.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (model *ProxyModel) Total(maps interface{}) (int, error) {
	var count int
	err := db.Model(model).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (model *ProxyModel) Count() (int, error) {
	var count int
	err := db.Model(model).Count(&count).Error
	if err != nil {
		return 0, nil
	}

	return count, nil
}

func (model *ProxyModel) Query(where interface{}, pageNum int, pageSize int) ([]ProxyModel,
	error) {
	var proxyModels []ProxyModel
	err := db.Model(model).Where(where).Offset(pageNum).Limit(pageSize).Find(&proxyModels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return proxyModels, nil
}

func (model *ProxyModel) Get(id int) (*ProxyModel, error) {
	var proxy ProxyModel
	err := db.Where("id = ?", id).First(&proxy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &proxy, nil
}

func (model *ProxyModel) Select(where interface{}) (*ProxyModel, error) {
	var proxy ProxyModel
	err := db.Model(model).Where(where).Offset(0).Limit(1).First(&proxy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &proxy, nil
}

func (model *ProxyModel) Edit() error {
	return db.Save(model).Error
}

func (model *ProxyModel) Add() error {
	return db.Create(model).Error
}

func (model *ProxyModel) Delete() error {
	return db.Delete(model).Error
}

func (model *ProxyModel) DeleteByCnd(where interface{}) error {
	return db.Where(where).Delete(model).Error
}

func (model *ProxyModel) CleanAll() error {
	return db.Delete(model).Error
}

func (model *ProxyModel) FindOrCreate(domain string) error {
	return db.FirstOrCreate(model, ProxyModel{Domain: domain}).Error
}

func (model *ProxyModel) RandomProxy() (ProxyModel, error) {
	var proxyModel ProxyModel

	where := make(map[string]interface{})
	where["status"] = 1

	err := db.Model(model).Where(where).Order(gorm.Expr("rand()")).First(&proxyModel).Error
	return proxyModel, err

}
