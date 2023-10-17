// usecase/attendance_usecase.go

package usecase

import (
	"kobasemi_attendance/domain"
	"kobasemi_attendance/repository"
)

type AttendanceUsecase struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceUsecase(repo repository.AttendanceRepository) *AttendanceUsecase {
	return &AttendanceUsecase{
		attendanceRepo: repo,
	}
}

func (uc *AttendanceUsecase) GetAllAttendances() ([]domain.Attendance, error) {
	return uc.attendanceRepo.GetAll()
}

func (uc *AttendanceUsecase) RegisterAttendance(attendance *domain.Attendance) error {
	return uc.attendanceRepo.Register(attendance)
}

func (uc *AttendanceUsecase) UpdateAttendanceStatus(name string, status bool) error {
	return uc.attendanceRepo.UpdateStatus(name, status)
}
