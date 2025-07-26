package lesson02

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "_12123"
	return
}

type CreditCard struct {
	gorm.Model
	Number     string
	BankUserID uint
}

type BankUser struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&BankUser{})
	// db.AutoMigrate(&CreditCard{})

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// //
	// result := db.Create(&user) // 通过数据的指针来创建

	// // 检查错误
	// if result.Error != nil {
	// 	log.Fatal(result.Error) // 处理错误
	// }
	// // 获取受影响的行数
	// fmt.Println("插入记录数:", result.RowsAffected)

	// // 获取插入后的主键值（假设ID是自增主键）
	// fmt.Println("插入数据的ID:", user.ID)

	// users := []*User{
	// 	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	// 	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	// }

	// db.Create(users)
	// fmt.Println(users[0])

	// db.Create(&BankUser{
	// 	Name:       "jinzhuss",
	// 	CreditCard: CreditCard{Number: "5511111111111"},
	// })

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	// user := User{Name: "Jinzhu", Age: 29, Birthday: time.Now()}
	// user.ID = 1
	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "name"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"age"}),
	// }).Create(&user)

	// var user User
	// db.Debug().First(&user)
	// err := db.Debug().Take(&user).Error
	// if err != nil {
	// 	panic(err)
	// }

	// user.ID = 1
	// db.Debug().First(&user, []int{1, 2, 3})

	// var users []User
	// db.Debug().Find(&users)

	var user User
	user.ID = 1
	db.Debug().First(&user, "name = ?", "123123")
}
