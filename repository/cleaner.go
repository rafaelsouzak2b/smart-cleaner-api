package repository

import (
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/guicazaroto/learning-go/util"
	"gorm.io/gorm"
)

type CleanerRepository struct {
	Db *gorm.DB
}

type ICleanerRepositoryport interface {
	GetCleaners(city string) []schemas.Cleaner
	GetCleanerById(cleanerID string) *schemas.Cleaner
	GetCleanerByEmailAndCpf(email, cpf string) int64
	CreateCleaner(cleaner *schemas.Cleaner) error
	SaveCleaner(cleaner *schemas.Cleaner) error
	DeleteCleaner(cleanerID string) error
	GetCleanerByEmailAndPassword(email, password string) *schemas.Cleaner
	UpdateImgUrlCleaner(cleaner *schemas.Cleaner, location string) error
}

func NewCleanerRepository(db *gorm.DB) *CleanerRepository {
	return &CleanerRepository{
		Db: db,
	}
}

func (r *CleanerRepository) GetCleaners(city string) []schemas.Cleaner {
	var cleaners []schemas.Cleaner
	if city != "" {
		r.Db.Joins("UserInfos").Where("active = ?", true).Where("cidade = ?", city).Find(&cleaners)
	} else {
		r.Db.Joins("UserInfos").Where("active = ?", true).Find(&cleaners)
	}
	return cleaners
}

func (r *CleanerRepository) GetCleanerById(cleanerID string) *schemas.Cleaner {
	var cleaner *schemas.Cleaner
	result := r.Db.Preload("UserInfos").First(&cleaner, cleanerID)
	if result.RowsAffected == 0 {
		return nil
	}

	return cleaner
}

func (r *CleanerRepository) GetCleanerByEmailAndCpf(email, cpf string) int64 {
	var cleaner schemas.Cleaner
	var count int64
	r.Db.Joins("UserInfos").Where("email = ?", email).Or("cpf = ?", cpf).First(&cleaner).Count(&count)
	return count
}

func (r *CleanerRepository) CreateCleaner(cleaner *schemas.Cleaner) error {
	if err := r.Db.Create(&cleaner).Error; err != nil {
		return err
	}
	return nil
}

func (r *CleanerRepository) SaveCleaner(cleaner *schemas.Cleaner) error {
	if err := r.Db.Save(&cleaner).Error; err != nil {
		return err
	}
	return nil
}

func (r *CleanerRepository) DeleteCleaner(cleanerID string) error {
	var cleaner schemas.Cleaner
	if err := r.Db.Unscoped().Delete(&cleaner, cleanerID).Error; err != nil {
		return err
	}
	return nil
}

func (r *CleanerRepository) GetCleanerByEmailAndPassword(email, password string) *schemas.Cleaner {
	var cleaner schemas.Cleaner
	result := r.Db.Joins("UserInfos").Where("active = ?", true).Where("email = ?", email).Where("password = ?", util.HashString(password)).First(&cleaner)
	if result.RowsAffected == 0 {
		return nil
	}
	return &cleaner
}

func (r *CleanerRepository) UpdateImgUrlCleaner(cleaner *schemas.Cleaner, location string) error {
	if err := r.Db.Model(&cleaner.UserInfos).Update("ImagemUrl", location).Error; err != nil {
		return err
	}
	return nil
}
