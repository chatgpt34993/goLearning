package lesson03_04

import (
	"gorm.io/gorm"
)

type Address struct {
	ID       uint `gorm:"primarykey"`
	Address1 string
	UserID   uint
}

type Email struct {
	ID     uint `gorm:"primarykey"`
	Email  string
	UserID uint
}

type Language struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Company struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type User struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	BillingAddress Address
	Emails         []Email
	Languages      []Language `gorm:"many2many:user_languages;"`
	CompanyID      uint
	Company        Company
}

func CreateUser(db *gorm.DB) {
	comp := Company{Name: "Comp"}
	db.Create(&comp)

	user := User{
		Name:           "jinzhu1111",
		BillingAddress: Address{Address1: "Billing Address - Address 1111"},
		Emails: []Email{
			{Email: "11@example.com"},
			{Email: "22@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
		CompanyID: comp.ID,
	}
	db.Create(&user)
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	//db.AutoMigrate(&Language{})

	// db.AutoMigrate(&Email{})
	// db.AutoMigrate(&Address{})
	// db.AutoMigrate(&Company{})

	// CreateUser(db)

	// var user User
	// db.First(&user)
	// db.Preload("BillingAddress").Preload("Emails").First(&user)
	// db.Preload(clause.Associations).First(&user)
	// fmt.Println(user)

	// var langs []Language
	// db.Model(&User{ID: 1}).Association("Languages").Find(&langs)
	// fmt.Println(langs)

	var user User
	//db.Preload(clause.Associations).First(&user)
	user.BillingAddress = Address{Address1: "111"}
	// user.BillingAddress.Address1 = "444"
	// // db.Debug().Updates(&user)
	// db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

}
