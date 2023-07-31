package user

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type userService struct {
	conn *gorm.DB
}

func (u *userService) Save() error {
	return nil
}

func (u *userService) Get() InterfaceUserService {
	return nil
}

func New() InterfaceUserService {
	dsn := "host=localhost user=postgres password=adelaida2011 dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.AutoMigrate(user{})
	if err != nil {
		fmt.Println(err)
	}

	return &userService{
		conn: db,
	}
}
