package lesson04

import (
	"fmt"

	"gorm.io/gorm"
)

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Cat struct {
	ID   int
	Name string
	Age  int
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Toy struct {
	ID        int
	Name      string
	Kind      string
	TID       int
	OwnerID   uint
	OwnerType string
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&Dog{}, &Cat{}, &Toy{})

	//多态
	// db.Create(&Dog{Name: "wangcai", Toy: Toy{Name: "gutou"}})
	// db.Create(&Cat{Name: "mimi", Toy: Toy{Name: "doumaobang"}, Age: 1})
	// db.Create(&Cat{Name: "mini2", Toy: Toy{Name: "doumaobang"}, Age: 2})
	// db.Create(&Cat{Name: "mini3", Toy: Toy{Name: "doumaobang"}, Age: 3})

	var dog Dog
	var cat Cat
	db.Preload("Toy").First(&dog)
	db.Preload("Toy").First(&cat)
	fmt.Println(dog, cat)

	// RunHook(db)
	// RunTransaction(db)
	// RunDefinition(db)

	db.Debug().Where("age > ?", 2).Delete(&Cat{})

}
