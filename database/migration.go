package database

import (
	"fmt"
	"kenangan/model"
	"kenangan/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&model.Kenangan{},
		&model.User{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
