package models

import (
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)

type Hotel struct {
	gorm.Model
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Country      string `json:"country"`
	State        string `json:"state"`
	City         string `json:"city"`
	Localisation string `json:"localisation"`
	Restaurant   bool   `json:"restaurant"`
	Rooms        []Room `gorm:"foreignKey:HotelId"`
}

func init() {
	config.Connect()
	DBS = config.GetDB()
	// DBS.DropTableIfExists(&Hotel{})
	DBS.AutoMigrate(&Hotel{})
}

func (h *Hotel) CreateHotel() *Hotel {
	DBS.NewRecord(h)
	DBS.Create(&h)
	return h
}

func GetAllHotel() []Hotel {
	var Hotels []Hotel
	DBS.Preload("Rooms").Find(&Hotels)
	return Hotels
}

func GetHotelById(Id int64) (*Hotel, *gorm.DB) {
	var getHotel Hotel
	db := DBS.Preload("Rooms").Where("ID = ?", Id).Find(&getHotel)
	return &getHotel, db
}

func DeleteHotelId(Id int64) Hotel {
	var hotel Hotel
	DBS.Where("ID = ?", Id).Delete(&hotel)
	return hotel
}
