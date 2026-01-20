package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := `Server=(localdb)\MSSQLLocalDB;Database=TodoList;Trusted_Connection=True;`

	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	DB = database
}
