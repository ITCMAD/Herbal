package pkg

import "gorm.io/gorm"

type User struct {
	ID       string `gorm:"primarykey"`
	Username string `gorm:""`
	Password string
}

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &UserManager{
		db: db,
	}
}
