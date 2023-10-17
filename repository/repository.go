package repository

import (
	"gorm.io/gorm"
	"kobasemi_attendance/domain"
)

type AttendanceRepository interface {
	GetAll() ([]domain.Attendance, error)
	Register(attendance *domain.Attendance) error
	UpdateStatus(name string, status bool) error
}

type AttendanceGormRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &AttendanceGormRepository{db: db}
}

func (repo *AttendanceGormRepository) GetAll() ([]domain.Attendance, error) {
	var attendances []domain.Attendance
	if err := repo.db.Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (repo *AttendanceGormRepository) Register(attendance *domain.Attendance) error {
	return repo.db.Create(attendance).Error
}

func (repo *AttendanceGormRepository) UpdateStatus(name string, status bool) error {
	return repo.db.Model(&domain.Attendance{}).Where("name = ?", name).Update("status", status).Error
}
