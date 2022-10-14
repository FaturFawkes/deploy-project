package main

import (
	"net/http"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
}

func AllUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var resQuery []User
		if err := db.Find(&resQuery).Error; err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Error on Database",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Data",
			"data":    resQuery,
		})
	}
}

func connectDB() *gorm.DB {
	dsn := "root:@tcp(mysql:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return db
}

func main() {
	e := echo.New()
	dbConn := connectDB()
	dbConn.AutoMigrate(&User{})
	e.Use(middleware.Logger())

	e.GET("/users", AllUser(dbConn))
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.Start(":8000")
}