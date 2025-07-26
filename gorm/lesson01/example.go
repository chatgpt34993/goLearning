package lesson01

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

type Blog2 struct {
	ID     int64
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	// Author  Author
	Upvotes int32
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Member{})
	// db.AutoMigrate(&Blog{})
	// db.AutoMigrate(&Blog2{})

	// user := &User{}
	// user.MemberNumber.Valid = true
	// db.Create(user)

	// create传指针
	// mem := Member{}
	// db.Create(&mem)
	// fmt.Println(mem.ID)

	db.Unscoped().First(&Member{}, 1)

	result := db.Delete(&Member{}, 2)
	if result.Error != nil {
		log.Fatalf("删除失败: %v", result.Error)
	}
	fmt.Printf("删除的记录数: %d\n", result.RowsAffected) // 若为0表示未找到记录
}
