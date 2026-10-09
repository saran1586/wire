package database

import (
"fmt"
"mypackage/model"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {

dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
err = db.AutoMigrate(&model.Student{})


if err != nil {
return nil, err
}


fmt.Println("database connected successfully")



return db, nil
}

