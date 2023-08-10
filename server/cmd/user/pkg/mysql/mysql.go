package mysql

import "gorm.io/gorm"

type User struct {
	ID       string `gorm:"primarykey"`
	Username string `gorm:"type:varchar(40)"`
	Password string `gorm:"type:varchar(40)"`
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
