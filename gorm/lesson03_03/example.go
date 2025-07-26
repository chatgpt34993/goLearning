package lesson03_03

import (
	"fmt"

	"gorm.io/gorm"
)

//	type User struct {
//		gorm.Model
//		Languages []Language
//	}
type Language struct {
	gorm.Model
	Name  string
	Refer uint `gorm:"unique"`
}

type User struct {
	gorm.Model
	Refer     uint       `gorm:"unique"`
	Languages []Language `gorm:"many2many:user_lang;foreignkey:Refer;joinForeignKey:UserReferID;References:Refer;joinReferences:LangReferID;"`
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	//db.AutoMigrate(&Language{})

	// user := User{}
	// db.Create(&user)

	// language := Language{Name: "english"}
	// db.Create(&language)
	//
	// language2 := Language{Name: "chinese", Refer: 12345}
	// db.Create(&language2)

	// user := User{
	// 	Refer: 12345, // 确保该值在数据库中唯一
	// 	Languages: []Language{
	// 		{Name: "en1"},
	// 		{Name: "en2"},
	// 	},
	// }
	// db.Create(&user)
	user := User{}
	err := db.Preload("Languages").Find(&user, 3).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
