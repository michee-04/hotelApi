package models

import (
	// "fmt"

	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

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
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	// u.Id = fmt.Sprintf("%x", u.Id)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword) 
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUser() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

func GetBookById(Id string) (*User, *gorm.DB){
	var GetUser User
	db := db.Where("ID=?", Id).Find(&GetUser)
	return &GetUser, db
}

func DeleteBookId(Id string) User{
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}