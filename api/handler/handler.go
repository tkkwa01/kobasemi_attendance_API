// api/handler/attendance_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"kobasemi_attendance/domain"
	"kobasemi_attendance/usecase"
)

type AttendanceHandler struct {
	attendanceUsecase *usecase.AttendanceUsecase
}

func NewAttendanceHandler(uc *usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceUsecase: uc,
	}
}

func (h *AttendanceHandler) RegisterAttendance(c *gin.Context) {
	var attendance domain.Attendance
	if err := c.BindJSON(&attendance); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	if err := h.attendanceUsecase.RegisterAttendance(&attendance); err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Registered"})
}

func (h *AttendanceHandler) GetAllAttendances(c *gin.Context) {
	attendances, err := h.attendanceUsecase.GetAllAttendances()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, attendances)
}

// 実質出席メソッド。出席したらtrueになる
func (h *AttendanceHandler) UpdateAttendanceStatus(c *gin.Context) {
	name := c.Param("name")
	var update domain.Attendance
	if err := c.BindJSON(&update); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	if err := h.attendanceUsecase.UpdateAttendanceStatus(name, update.Status); err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Status updated"})
}
