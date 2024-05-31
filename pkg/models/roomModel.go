package models

import (
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)

type Room struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	BedCount    string `json:"bedCount"`
	Image       string `json:"image"`
	RoomPrice   string `json:"roomPrice"`
	RoomService bool   `json:"roomService"`
	OceanView   bool   `json:"oceanView"`
	CityView    bool   `json:"cityView"`
	HotelId     uint   `json:"hotelId"`
	Hotel       Hotel  `gorm:"foreignKey:HotelId" json:"-"`
}

func init() {
	config.Connect()
	DBS = config.GetDB()
	// DBS.DropTableIfExists(&Room{})
	DBS.AutoMigrate(&Room{})
}

func (r *Room) CreateRoom() *Room {
	DBS.NewRecord(r)
	DBS.Create(&r)
	return r
}

func GetAllRoom() []Room {
	var rooms []Room
	DBS.Find(&rooms)
	return rooms
}

func GetRoomById(Id int64) (*Room, *gorm.DB) {
	var getRoom Room
	db := DBS.Where("ID = ?", Id).Find(&getRoom)
	return &getRoom, db
}

func DeleteRoomById(Id int64) Room {
	var room Room
	DBS.Where("ID = ?", Id).Delete(&room)
	return room
}
