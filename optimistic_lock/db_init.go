package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint64
	SaleCount uint64
	Version   uint64 // Version field for optimistic locking
}

func (Product) TableName() string {
	return "product"
}

func initDB() *gorm.DB {
	dsn := "root:12341234@tcp(127.0.0.1:3306)/golang-study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.Model(&Product{}).Where("1=1").Update("sale_count", 0).Error; err != nil {
		panic(err.Error())
	}

	if err := db.Model(&Product{}).Where("1=1").Update("version", 0).Error; err != nil {
		panic(err.Error())
	}

	return db
}
