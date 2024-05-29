package models

import (
	// "fmt"

	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var DBS *gorm.DB

type User struct {
	gorm.Model
	// Id string `gorm:"primary_key;type:varchar(255)" json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func init(){
	config.Connect()
	DBS = config.GetDB()
	DBS.AutoMigrate(&User{})
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
	DBS.Find(&Users)
	return Users
}

func GetBookById(Id string) (*User, *gorm.DB){
	var GetUser User
	db := DBS.Where("ID=?", Id).Find(&GetUser)
	return &GetUser, db
}

func DeleteBookId(Id string) User{
	var user User
	DBS.Where("ID=?", Id).Delete(user)
	return user
}
