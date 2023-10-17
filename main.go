// main.go

package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"kobasemi_attendance/api/handler"
	"kobasemi_attendance/database"
	"kobasemi_attendance/domain"
	"kobasemi_attendance/repository"
	"kobasemi_attendance/usecase"
	"os"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewAttendanceRepository(db)
	uc := usecase.NewAttendanceUsecase(repo)
	handler := handler.NewAttendanceHandler(uc)

	// テキストファイルを開く
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening names.txt:", err)
		return
	}
	defer file.Close()

	// ファイルから行を一つずつ読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()

		// データベースに書き込む
		attendance := &domain.Attendance{Name: name, Status: false}
		result := db.Create(attendance)
		if result.Error != nil {
			fmt.Println("Error writing to database:", result.Error)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading names.txt:", scanner.Err())
		return
	}

	r := gin.Default()
	r.POST("/attendance/register", handler.RegisterAttendance)
	r.GET("/attendance/watch", handler.GetAllAttendances)
	r.PATCH("/attendance/:name", handler.UpdateAttendanceStatus)
	r.Run(":8085")
}
