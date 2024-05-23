package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "log"
)

var DB *gorm.DB

func Init() {
    
    dsn := "root:root@tcp(127.0.0.1:3306)/venta_de_cursos?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
}