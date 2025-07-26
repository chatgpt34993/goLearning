package lesson03_02

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&CreditCard{})

	// user := User{}
	// db.Preload("CreditCards").First(&user, 1)
	// fmt.Println(user)

	// 创建用户和信用卡
	// user := User{
	// 	CreditCards: []CreditCard{
	// 		{Number: "1234-5678-9012-3456"},
	// 		{Number: "4321-8765-2109-6543"},
	// 	},
	// }
	// db.Create(&user) // 会同时创建用户和关联的信用卡

	// 查询用户及其信用卡
	var user User
	db.Preload("CreditCards").First(&user, 1) // 预加载关联的信用卡

}
