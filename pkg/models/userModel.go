package models

import (
	// "fmt"

	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var DBS *gorm.DB

type User struct {
	// gorm.Model
	ID        string `gorm:"primary_key"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time 
	UpdatedAt time.Time  
	DeletedAt *time.Time
	Hotel        []Hotel `gorm:"foreignKey:UserID"`
}

func init(){
	config.Connect()
	DBS = config.GetDB()
	DBS.DropTableIfExists(&User{})
	DBS.AutoMigrate(&User{})
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("ID", id)
}

func (u *User) CreateUser() *User{
	// u.Id = fmt.Sprintf("%x", u.Id)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword) 
	DBS.NewRecord(u)
	DBS.Create(&u)
	return u
}

func GetAllUser() []User{
	var Users []User
	DBS.Preload("Hotel").Find(&Users)
	return Users
}

func GetUserById(Id string) (*User, *gorm.DB){
	var GetUser User
	db := DBS.Preload("Hotel").Where("ID=?", Id).Find(&GetUser)
	return &GetUser, db
}

func DeleteUserId(Id string) User{
	var user User
	DBS.Where("ID=?", Id).Delete(user)
	return user
}
