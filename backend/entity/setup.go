package entity


import (

"gorm.io/driver/sqlite"

"gorm.io/gorm"

)


var db *gorm.DB


func DB() *gorm.DB {

return db

}


func SetupDatabase() {

database, err := gorm.Open(sqlite.Open("ProjectSA.db"), &gorm.Config{})

if err != nil {

panic("failed to connect database")

}

// Migrate the schema

database.AutoMigrate(
	&Member{},
	&Creator{},
	&Cartoon{},
	&Episodes{},
	&Package{},
	&PaymentCoin{},
	&PaymentEpisode{},
	&History{},
	&Rating{},
	&Comment{},

)

db = database

member := Member{
	Username: "Chinnpatz",
	Password: "1234",
	Email: "test@gmail.com",
}
db.Model(&Member{}).Create(&member)
}